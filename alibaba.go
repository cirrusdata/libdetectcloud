package libdetectcloud

import (
	"net/http"
)

const VendorAlibaba = "Alibaba Cloud"

func detectAlibaba() string {
	// Alibaba ECS metadata lives on a carrier-grade NAT address, not link-local.
	resp, err := hc.Get("http://100.100.100.200/latest/meta-data/instance/instance-id")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return VendorAlibaba
		}
	}
	if dmiVendorHasPrefix("Alibaba Cloud") || dmiProductContains("Alibaba Cloud ECS") {
		return VendorAlibaba
	}
	return ""
}
