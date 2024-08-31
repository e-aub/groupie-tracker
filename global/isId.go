package global

func IsId(id string) bool {
	if len(id) > 3 {
		return false
	}
	for _, char := range id {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
