package libdetectcloud

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const VendorKVM = "KVM"

func detectKVM() string {
	if runtime.GOOS != "windows" {
		data, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			return ""
		}
		// if vendor is qemu or kvm (case insensitive) then return KVM
		if strings.Contains(strings.ToLower(string(data)), "qemu") || strings.Contains(strings.ToLower(string(data)), "kvm") {
			return VendorKVM
		}
		// also check if /dev/virtio-ports/org.qemu.guest_agent.0 exists (in OpenShift Virtualization, sys_vendor says Red Hat)
		if _, err := os.Stat("/dev/virtio-ports/org.qemu.guest_agent.0"); err == nil {
			return VendorKVM
		}
		return ""
	} else {
		cmd := exec.Command("powershell.exe", "get-wmiobject win32_computersystem | fl model")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return ""
		}
		// if vendor is qemu or kvm (case insensitive) then return KVM
		if strings.Contains(strings.ToLower(string(out)), "qemu") || strings.Contains(strings.ToLower(string(out)), "kvm") {
			return VendorKVM
		}
	}
	return ""
}
