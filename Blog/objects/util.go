package soul

func Contains(element string, array []string) bool {
	for _, v := range array {
		if element == v {
			return true
		}
	}
	return false
}
