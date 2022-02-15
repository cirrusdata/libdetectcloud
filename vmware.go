package libdetectcloud

import (
	"io/ioutil"
	"runtime"
	"strings"
	"exec"
)

func detectVMware() string {
	if runtime.GOOS != "windows" {
		data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			return ""
		}
		if strings.Contains(string(data), "VMware, Inc.") {
			return "VMware"
		}
		return ""
	} else {
		cmd := exec.Command("powershell.exe","get-wmiobject win32_computersystem | fl model")
		out,err := cmd.CombinedOutput()
		if err != nil {
			return ""
		}
		if strings.Contains(string(out), "VMware Virtual Platform") {
			return "VMware"
		}
	}
	return ""
}
