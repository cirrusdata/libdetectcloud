package libdetectcloud

const VendorLinode = "Linode"

func detectLinode() string {
	if dmiVendorHasPrefix("Linode") {
		return VendorLinode
	}
	return ""
}
