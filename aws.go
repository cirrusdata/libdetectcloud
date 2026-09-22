package libdetectcloud

import (
	"net/http"
)

func detectAWS() string {
	// instance-identity is AWS-only; /latest/ also answers on other
	// AWS-compatible metadata services.
	resp, err := hc.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return "Amazon Web Services"
		}
	}
	// DMI fallback for hosts where the metadata endpoint is unreachable
	// (for example when IMDSv1 is disabled).
	if dmiVendorHasPrefix("Amazon EC2") {
		return "Amazon Web Services"
	}
	return ""
}
