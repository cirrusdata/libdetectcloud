package libdetectcloud

const VendorCloudStack = "CloudStack"

func detectCloudStack() string {
	// CloudStack injects SMBIOS type 1 with manufacturer
	// "Apache Software Foundation" and product "CloudStack KVM Hypervisor".
	if dmiVendorHasPrefix("Apache Software Foundation") {
		return VendorCloudStack
	}
	return ""
}
