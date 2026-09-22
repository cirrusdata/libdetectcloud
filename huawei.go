package libdetectcloud

const VendorHuawei = "Huawei Cloud"

func detectHuawei() string {
	// Bare-metal Huawei servers report plain "Huawei"; the ECS guest vendor
	// string is "Huawei Cloud". Its metadata service is AWS-compatible, so
	// DMI is the distinguishing signal.
	if dmiVendorHasPrefix("Huawei Cloud") {
		return VendorHuawei
	}
	return ""
}
