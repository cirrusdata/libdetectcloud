package libdetectcloud

const VendorKubeVirt = "KubeVirt"

func detectKubeVirt() string {
	// Covers plain KubeVirt plus platforms built on it (Harvester, OpenShift
	// Virtualization / OKD).
	if dmiVendorHasPrefix("KubeVirt") {
		return VendorKubeVirt
	}
	return ""
}
