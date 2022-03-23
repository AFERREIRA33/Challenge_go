package piscine

func IsNumeric(s string) bool {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] >= '0' && z[i] <= '9' {
		} else {
			return false
		}
	}
	return true
}
