package libdetectcloud

import (
	"net/http"
)

func detectDigitalOcean() string {
	resp, err := hc.Get("http://169.254.169.254/metadata/v1.json")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return "Digital Ocean"
		}
	}
	if dmiVendorHasPrefix("DigitalOcean") {
		return "Digital Ocean"
	}
	return ""
}
