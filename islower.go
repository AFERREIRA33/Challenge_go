package piscine

func IsLower(s string) bool {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] < 'a' || z[i] > 'z' {
			return false
		}
	}
	return true
}
