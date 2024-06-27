package libdetectcloud

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const VendorNutanix = "Nutanix"

func detectNutanix() string {
	if runtime.GOOS != "windows" {
		data, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			return ""
		}
		if strings.Contains(string(data), "Nutanix") {
			return VendorNutanix
		}
		return ""
	} else {
		cmd := exec.Command("powershell.exe", "get-wmiobject win32_computersystem | fl model")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return ""
		}
		if strings.Contains(string(out), "Nutanix") {
			return VendorNutanix
		}
	}
	return ""
}
