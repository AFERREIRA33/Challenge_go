package piscine

func LastRune(s string) rune {
	l := []rune(s)
	return l[len(s)-1]
}
