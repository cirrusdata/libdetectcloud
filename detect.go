package libdetectcloud

import (
	"net/http"
	"runtime"
	"sync"
	"time"
)

var hc = &http.Client{Timeout: 300 * time.Millisecond}

func init() {
	trans := http.DefaultTransport.(*http.Transport).Clone()
	trans.Proxy = nil
	hc.Transport = trans
}

// detector probes one environment and returns its display name or "".
type detector func() string

// detectors run concurrently, but their results are considered in list order
// and the first non-empty result wins. Order matters:
//   - cloud metadata endpoint probes come first because they are authoritative
//   - the container check follows, as before
//   - DMI-based hypervisor checks come next, most specific first
//   - detectKVM is last: QEMU/KVM SMBIOS values are the default that more
//     specific platforms (CloudStack, oVirt, Proxmox, ...) share
var detectors = []detector{
	detectAWS,
	detectAzure,
	detectDigitalOcean,
	detectGCE,
	detectSoftlayer,
	detectVultr,
	detectOCI,
	detectAlibaba,
	detectTencent,
	detectHetzner,
	detectContainer,
	detectVMware,
	detectNutanix,
	detectOVirt,
	detectCloudStack,
	detectKubeVirt,
	detectOpenStack,
	detectHyperV,
	detectXen,
	detectVirtualBox,
	detectOpenNebula,
	detectLinode,
	detectUpCloud,
	detectHuawei,
	detectKVM,
}

// Detect returns the detected cloud or virtualization environment, or an
// empty string when the environment is unknown.
func Detect() string {
	if runtime.GOOS == "darwin" {
		return ""
	}
	results := make([]string, len(detectors))
	var wg sync.WaitGroup
	for i, d := range detectors {
		wg.Add(1)
		go func(i int, fn detector) {
			defer wg.Done()
			results[i] = fn()
		}(i, d)
	}
	wg.Wait()
	for _, r := range results {
		if r != "" {
			return r
		}
	}
	return ""
}
