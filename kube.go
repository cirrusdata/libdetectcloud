package libdetectcloud

import (
	"os"
	"strings"
)

func detectContainer() string {
	// cgroup v2 no longer carries runtime names; the marker files do.
	if _, err := os.Stat("/run/.containerenv"); err == nil {
		return "Container"
	}
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "Container"
	}

	b, err := os.ReadFile("/proc/self/cgroup")
	if err != nil {
		return ""
	}

	fc := string(b)
	kube := strings.Contains(fc, "kube")
	container := strings.Contains(fc, "containerd")

	if kube {
		return "K8S Container"
	}

	if container {
		return "Container"
	}

	return ""
}
