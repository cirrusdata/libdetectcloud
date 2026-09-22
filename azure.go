package libdetectcloud

import (
	"net/http"
)

func detectAzure() string {
	// Azure IMDS requires the Metadata header.
	r, err := http.NewRequest("GET", "http://169.254.169.254/metadata/instance?api-version=2020-09-01", nil)
	if err == nil {
		r.Header.Add("Metadata", "true")
		resp, err := hc.Do(r)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return "Microsoft Azure"
			}
		}
	}
	// Fallback for Azure Stack-style metadata services.
	resp, err := hc.Get("http://169.254.169.254/metadata/v1/InstanceInfo")
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return "Microsoft Azure"
		}
	}
	return ""
}
