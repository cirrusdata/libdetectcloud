package libdetectcloud

const VendorAlibaba = "Alibaba Cloud"

func detectAlibaba() string {
	// Deliberately no metadata probe: we do not want deployed agents to
	// contact the ECS metadata service (100.100.100.200). Detection is
	// DMI-only.
	if dmiVendorHasPrefix("Alibaba Cloud") || dmiProductContains("Alibaba Cloud ECS") {
		return VendorAlibaba
	}
	return ""
}
