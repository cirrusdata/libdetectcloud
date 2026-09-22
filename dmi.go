package libdetectcloud

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// dmi holds the SMBIOS fields used for vendor detection. The fields are read
// once per process: from sysfs on Linux and from a single WMI query on
// Windows, so all DMI-based detectors share one read.
type dmi struct {
	vendorFields []string
	product      string
}

var (
	dmiOnce sync.Once
	dmiLoad = loadDMI
	dmiData dmi
)

func loadDMI() {
	if runtime.GOOS == "windows" {
		loadWindowsDMI()
		return
	}
	for _, f := range []string{"sys_vendor", "board_vendor", "chassis_vendor", "bios_vendor"} {
		if v, err := os.ReadFile("/sys/class/dmi/id/" + f); err == nil {
			dmiData.vendorFields = append(dmiData.vendorFields, strings.TrimSpace(string(v)))
		}
	}
	if v, err := os.ReadFile("/sys/class/dmi/id/product_name"); err == nil {
		dmiData.product = strings.TrimSpace(string(v))
	}
}

func loadWindowsDMI() {
	out, err := exec.Command("powershell.exe", "-NoProfile", "-Command",
		"Get-WmiObject win32_computersystem | fl Manufacturer, Model").CombinedOutput()
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(line, ":", 2)
		if len(kv) != 2 {
			continue
		}
		switch strings.TrimSpace(kv[0]) {
		case "Manufacturer":
			dmiData.vendorFields = append(dmiData.vendorFields, strings.TrimSpace(kv[1]))
		case "Model":
			dmiData.product = strings.TrimSpace(kv[1])
		}
	}
}

func dmiVendors() []string {
	dmiOnce.Do(dmiLoad)
	return dmiData.vendorFields
}

func dmiProduct() string {
	dmiOnce.Do(dmiLoad)
	return dmiData.product
}

// dmiVendorHasPrefix reports whether any SMBIOS vendor field (system,
// baseboard, chassis, or BIOS vendor) starts with s.
func dmiVendorHasPrefix(s string) bool {
	for _, v := range dmiVendors() {
		if strings.HasPrefix(v, s) {
			return true
		}
	}
	return false
}

// dmiProductContains reports whether the SMBIOS product name contains s.
func dmiProductContains(s string) bool {
	return strings.Contains(dmiProduct(), s)
}
