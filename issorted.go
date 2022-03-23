package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	c := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) >= 0 {
			c++
		}
	}
	if c == len(a)-1 {
		return true
	} else {
		c = 0
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) <= 0 {
				c++
			}
		}
	}
	if c == len(a)-1 {
		return true
	}
	return false
}
