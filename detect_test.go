package libdetectcloud

import (
	"reflect"
	"sync"
	"testing"
)

// setDMI installs fake SMBIOS values for DMI-based detectors. Tests run in the
// same package, so the cache can be reset between cases.
func setDMI(vendors []string, product string) {
	dmiOnce = sync.Once{}
	dmiData = dmi{vendorFields: vendors, product: product}
}

func TestDMIDetectors(t *testing.T) {
	cases := []struct {
		name    string
		vendors []string
		product string
		detect  detector
		want    string
	}{
		{"vmware", []string{"VMware, Inc."}, "VMware Virtual Platform", detectVMware, "VMware"},
		{"vmware-vmw", []string{"VMW"}, "", detectVMware, "VMware"},
		{"nutanix", []string{"Nutanix"}, "", detectNutanix, "Nutanix"},
		{"ovirt", []string{"oVirt", "Red Hat"}, "RHEL", detectOVirt, "oVirt"},
		{"cloudstack", []string{"Apache Software Foundation"}, "CloudStack KVM Hypervisor", detectCloudStack, "CloudStack"},
		{"kubevirt", []string{"KubeVirt"}, "None", detectKubeVirt, "KubeVirt"},
		{"openstack", []string{"OpenStack Foundation"}, "OpenStack Nova", detectOpenStack, "OpenStack"},
		{"hyperv", []string{"Microsoft Corporation"}, "Virtual Machine", detectHyperV, "Hyper-V"},
		{"hyperv-needs-model", []string{"Microsoft Corporation"}, "Surface Pro", detectHyperV, ""},
		{"xen", []string{"Xen"}, "HVM domU", detectXen, "Xen"},
		{"xen-product", []string{"OEM"}, "HVM domU", detectXen, "Xen"},
		{"virtualbox", []string{"innotek GmbH"}, "VirtualBox", detectVirtualBox, "VirtualBox"},
		{"virtualbox-product", []string{"Oracle Corporation"}, "VirtualBox", detectVirtualBox, "VirtualBox"},
		{"opennebula", []string{"OpenNebula"}, "", detectOpenNebula, "OpenNebula"},
		{"linode", []string{"Linode"}, "Linode", detectLinode, "Linode"},
		{"upcloud", []string{"UpCloud"}, "", detectUpCloud, "UpCloud"},
		{"huawei", []string{"Huawei Cloud"}, "", detectHuawei, "Huawei Cloud"},
		{"huawei-baremetal", []string{"Huawei"}, "RH2288H V3", detectHuawei, ""},
		{"aws-dmi", []string{"Amazon EC2"}, "t3.large", detectAWS, "Amazon Web Services"},
		{"gce-dmi", []string{"Google"}, "Google Compute Engine", detectGCE, "Google Compute Engine"},
		{"do-dmi", []string{"DigitalOcean"}, "Droplet", detectDigitalOcean, "Digital Ocean"},
		{"oci-dmi", []string{"OracleCloud"}, "VM.Standard.E4", detectOCI, "Oracle Cloud"},
		{"alibaba-dmi", []string{"Alibaba Cloud"}, "Alibaba Cloud ECS", detectAlibaba, "Alibaba Cloud"},
		{"alibaba-product", []string{"Other"}, "Alibaba Cloud ECS", detectAlibaba, "Alibaba Cloud"},
		{"tencent-dmi", []string{"Tencent Cloud"}, "CVM", detectTencent, "Tencent Cloud"},
		{"tencent-legacy", []string{"Smdbmds"}, "", detectTencent, "Tencent Cloud"},
		{"hetzner-dmi", []string{"Hetzner"}, "", detectHetzner, "Hetzner"},
		{"kvm-qemu", []string{"QEMU"}, "Standard PC (i440FX + PIIX, 1996)", detectKVM, "KVM"},
		{"kvm-product", []string{"Red Hat"}, "KVM", detectKVM, "KVM"},
		{"kvm-empty", []string{"Dell Inc."}, "PowerEdge R740", detectKVM, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setDMI(c.vendors, c.product)
			if got := c.detect(); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

// When the machine runs on a hypervisor that matches no signature, Detect
// falls back to KVM rather than returning empty. A bare-metal machine (no
// hypervisor-present flag) still returns empty.
func TestKVMHypervisorFallback(t *testing.T) {
	setDMI([]string{"Some Vendor"}, "Some Product")
	dmiData.hypervisor = true
	if got := detectKVM(); got != "KVM" {
		t.Errorf("hypervisor fallback: got %q, want KVM", got)
	}

	setDMI([]string{"Dell Inc."}, "PowerEdge R740")
	if got := detectKVM(); got != "" {
		t.Errorf("bare metal: got %q, want empty", got)
	}
}

func detectorIndex(d detector) int {
	p := reflect.ValueOf(d).Pointer()
	for i, x := range detectors {
		if reflect.ValueOf(x).Pointer() == p {
			return i
		}
	}
	return -1
}

// Ordering-sensitive cases: CloudStack, oVirt and KubeVirt guests are all
// QEMU/KVM underneath, so the specific vendor check must win over the generic
// KVM one. Only the DMI-based span of detectors is exercised so the test does
// not depend on network access.
func TestDMIDetectorOrdering(t *testing.T) {
	start := detectorIndex(detectVMware)
	end := detectorIndex(detectKVM)
	if start < 0 || end < 0 || end <= start {
		t.Fatalf("unexpected detector order: vmware=%d kvm=%d", start, end)
	}
	cases := []struct {
		name       string
		vendors    []string
		product    string
		hypervisor bool
		want       string
	}{
		{"cloudstack-not-kvm", []string{"Apache Software Foundation"}, "CloudStack KVM Hypervisor", true, "CloudStack"},
		{"ovirt-not-kvm", []string{"oVirt", "Red Hat"}, "RHEL", true, "oVirt"},
		{"kubevirt-not-kvm", []string{"KubeVirt"}, "None", true, "KubeVirt"},
		{"proxmox-is-kvm", []string{"QEMU"}, "Standard PC (i440FX + PIIX, 1996)", true, "KVM"},
		{"unidentified-hypervisor-is-kvm", []string{"Some Vendor"}, "Some Product", true, "KVM"},
		{"baremetal", []string{"Dell Inc."}, "PowerEdge R740", false, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setDMI(c.vendors, c.product)
			dmiData.hypervisor = c.hypervisor
			got := ""
			for _, d := range detectors[start : end+1] {
				if got = d(); got != "" {
					break
				}
			}
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
