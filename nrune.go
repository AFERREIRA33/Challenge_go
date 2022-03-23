package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n <= 0 {
		return 0
	}
	l := []rune(s)
	return l[n-1]
}
