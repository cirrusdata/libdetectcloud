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
	return ""
}
