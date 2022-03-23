package piscine

func Capitalize(s string) string {
	z := []rune(s)
	if z[0] >= 'a' && z[0] <= 'z' {
		z[0] = z[0] - 32
	}
	for i := 1; i < len(s); i++ {
		if z[i-1] >= 'A' && z[i-1] <= 'Z' || z[i-1] >= 'a' && z[i-1] <= 'z' || z[i-1] >= '0' && z[i-1] <= '9' {
			if z[i] >= 'A' && z[i] <= 'Z' {
				z[i] = z[i] + 32
			}
		} else if z[i] >= 'a' && z[i] <= 'z' {
			z[i] = z[i] - 32
		}
	}
	l := string(z)
	return l
}
