package libdetectcloud

import (
	"net/http"
)

const VendorOCI = "Oracle Cloud"

func detectOCI() string {
	resp, err := hc.Get("http://169.254.169.254/opc/v1/instance/")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return VendorOCI
		}
	}
	if dmiVendorHasPrefix("OracleCloud") {
		return VendorOCI
	}
	return ""
}
