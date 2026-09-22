package libdetectcloud

const VendorOpenNebula = "OpenNebula"

func detectOpenNebula() string {
	if dmiVendorHasPrefix("OpenNebula") {
		return VendorOpenNebula
	}
	return ""
}
