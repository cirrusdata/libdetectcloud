package libdetectcloud

import (
	"net/http"
)

const VendorTencent = "Tencent Cloud"

func detectTencent() string {
	resp, err := hc.Get("http://metadata.tencentyun.com/latest/meta-data/instance-id")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return VendorTencent
		}
	}
	// "Smdbmds" is the sys_vendor on older Tencent CVM guests.
	if dmiVendorHasPrefix("Tencent Cloud") || dmiVendorHasPrefix("Smdbmds") {
		return VendorTencent
	}
	return ""
}
