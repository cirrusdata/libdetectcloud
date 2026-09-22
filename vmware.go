package libdetectcloud

func detectVMware() string {
	if dmiVendorHasPrefix("VMware") || dmiVendorHasPrefix("VMW") {
		return "VMware"
	}
	return ""
}
