package libdetectcloud

func detectOpenStack() string {
	// "OpenStack Foundation" is the usual sys_vendor; prefix matching also
	// covers variants like "OpenStackNova".
	if dmiVendorHasPrefix("OpenStack") {
		return "OpenStack"
	}
	return ""
}
