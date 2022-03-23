package piscine

func BasicAtoi(s string) int {
	o := 0
	c := 0
	a_s := []rune(s)
	for _, word := range a_s {
		for i := '0'; i < word; i++ {
			c++
		}
		o = o*10 + c
		c = 0
	}
	return o
}
