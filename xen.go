package libdetectcloud

const VendorXen = "Xen"

func detectXen() string {
	if dmiVendorHasPrefix("Xen") || dmiProductContains("HVM domU") {
		return VendorXen
	}
	return ""
}
