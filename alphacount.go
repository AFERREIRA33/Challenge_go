package piscine

func AlphaCount(s string) int {
	z := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if z[i] >= 'A' && z[i] <= 'Z' || z[i] >= 'a' && z[i] <= 'z' {
			count++
		}
	}
	return count
}
