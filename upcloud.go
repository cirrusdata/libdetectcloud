package libdetectcloud

const VendorUpCloud = "UpCloud"

func detectUpCloud() string {
	if dmiVendorHasPrefix("UpCloud") {
		return VendorUpCloud
	}
	return ""
}
