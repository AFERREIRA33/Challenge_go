package piscine

func ToLower(s string) string {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] >= 'A' && z[i] <= 'Z' {
			z[i] = z[i] + 32
		}
	}
	l := string(z)
	return l
}
