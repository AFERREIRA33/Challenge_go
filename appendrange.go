package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if min >= max {
	} else {
		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
	}
	return tab
}
