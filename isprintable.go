package piscine

func IsPrintable(s string) bool {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] < 32 {
			return false
		}
	}
	return true
}
