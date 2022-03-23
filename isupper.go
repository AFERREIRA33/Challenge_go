package piscine

func IsUpper(s string) bool {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] < 'A' || z[i] > 'Z' {
			return false
		}
	}
	return true
}
