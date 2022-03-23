package piscine

func ToUpper(s string) string {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] >= 'a' && z[i] <= 'z' {
			z[i] = z[i] - 32
		}
	}
	l := string(z)
	return l
}
