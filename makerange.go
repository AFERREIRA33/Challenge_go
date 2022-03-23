package piscine

func MakeRange(min, max int) []int {
	size := 0
	if min >= max {
		tab := make([]int, size)
		tab = nil
		return tab
	} else {
		size = max - min
		tab := make([]int, size)
		for i := 0; i < size; i++ {
			tab[i] = i + min
		}
		return tab
	}
}
