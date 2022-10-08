package slices

func CzyZawiera[T comparable](slice []T, val T) bool {
	for _, k := range slice {
		if k == val {
			return true
		}
	}
	return false
}
