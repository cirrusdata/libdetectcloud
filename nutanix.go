package libdetectcloud

const VendorNutanix = "Nutanix"

func detectNutanix() string {
	if dmiVendorHasPrefix("Nutanix") {
		return VendorNutanix
	}
	return ""
}
