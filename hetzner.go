package libdetectcloud

import (
	"net/http"
)

const VendorHetzner = "Hetzner"

func detectHetzner() string {
	resp, err := hc.Get("http://169.254.169.254/hetzner/v1/metadata")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return VendorHetzner
		}
	}
	if dmiVendorHasPrefix("Hetzner") {
		return VendorHetzner
	}
	return ""
}
