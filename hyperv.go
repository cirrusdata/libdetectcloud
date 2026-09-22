package libdetectcloud

const VendorHyperV = "Hyper-V"

func detectHyperV() string {
	// Hyper-V guests report manufacturer "Microsoft Corporation" and product
	// "Virtual Machine". Azure guests report the same values, so this runs
	// only after the Azure metadata probe has had its say.
	if dmiVendorHasPrefix("Microsoft Corporation") && dmiProductContains("Virtual Machine") {
		return VendorHyperV
	}
	return ""
}
