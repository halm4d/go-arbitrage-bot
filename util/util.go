package util

func Contains(slice []string, t string) bool {
	for _, s := range slice {
		if s == t {
			return true
		}
	}
	return false
}
