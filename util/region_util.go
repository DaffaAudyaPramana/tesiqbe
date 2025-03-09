package util

// FormatRegion formats region details into a readable string.
func FormatRegion(villageName, districtName, regencyName, provinceName string) string {
	return villageName + ", " + districtName + ", " + regencyName + ", " + provinceName
}
