package piscine

func Map(f func(int) bool, a []int) []bool {
	taboolé := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			taboolé[i] = true
		} else {
			taboolé[i] = false
		}
	}
	return taboolé
}
