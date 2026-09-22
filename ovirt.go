package libdetectcloud

const VendorOVirt = "oVirt"

func detectOVirt() string {
	if dmiVendorHasPrefix("oVirt") {
		return VendorOVirt
	}
	return ""
}
