package libdetectcloud

import (
	"os"
	"runtime"
)

const VendorKVM = "KVM"

func detectKVM() string {
	// Generic QEMU/KVM SMBIOS values. More specific platforms built on QEMU
	// (CloudStack, oVirt, Proxmox, KubeVirt) are checked before this one.
	if dmiVendorHasPrefix("QEMU") || dmiVendorHasPrefix("KVM") ||
		dmiProductContains("KVM") || dmiProductContains("Standard PC") {
		return VendorKVM
	}
	if runtime.GOOS != "windows" {
		// OpenShift Virtualization reports vendor "Red Hat"; the qemu guest
		// agent channel still identifies the guest as KVM.
		if _, err := os.Stat("/dev/virtio-ports/org.qemu.guest_agent.0"); err == nil {
			return VendorKVM
		}
	}
	// Last resort: a hypervisor we cannot name beats an empty result. KVM is
	// the sensible default since unidentified hypervisors in our fleet are
	// QEMU/KVM-family. On Windows, HypervisorPresent is also true on a
	// Hyper-V host's root partition, so that edge reports KVM too.
	if dmiHypervisor() {
		return VendorKVM
	}
	return ""
}
