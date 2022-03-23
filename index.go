package piscine

func Index(s string, toFind string) int {
	count := 0
	z := []rune(toFind)
	l := []rune(s)
	for q, lll := range s {
		if count == 0 {
			if lll == z[0] {
				count = q
				break
			}
		}
	}
	for i := 0; i < len(z); i++ {
		if l[count+i] == z[i] {
		} else {
			return -1
		}
	}
	return count
}
