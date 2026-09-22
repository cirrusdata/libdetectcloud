package libdetectcloud

const VendorVirtualBox = "VirtualBox"

func detectVirtualBox() string {
	if dmiProductContains("VirtualBox") || dmiVendorHasPrefix("innotek") {
		return VendorVirtualBox
	}
	return ""
}
