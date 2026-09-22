package libdetectcloud

const VendorTencent = "Tencent Cloud"

func detectTencent() string {
	// Deliberately no metadata probe: we do not want deployed agents to
	// contact tencentyun.com. Detection is DMI-only.
	// "Smdbmds" is the sys_vendor on older Tencent CVM guests.
	if dmiVendorHasPrefix("Tencent Cloud") || dmiVendorHasPrefix("Smdbmds") {
		return VendorTencent
	}
	return ""
}
