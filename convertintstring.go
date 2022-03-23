package piscine

func Convertintstr(n int) string {
	var tab []int
	var result string
	o := n
	c := 0
	e := 1
	r := 0
	moins := false
	if n < 0 {
		n = n * -1
		moins = true
	}
	for o > 9 {
		o = o / 10
		c++
	}
	o = n
	for j := 0; j <= c; j++ {
		if o > 9 {
			for o > 9 {
				o = o / 10
				e = e * 10
			}
			tab = append(tab, o)
			r = r + o*e
			o = n - r
			if o < e/10 {
				tab = append(tab, 0)
			}
			e = 1
		} else {
			tab = append(tab, o)
		}
		if tab[j] == 0 {
			j++
		}
	}
	if moins {
		result = result + "-"
	}
	for i := 0; i < len(tab); i++ {
		result = result + string(tab[i]+48)
	}
	return result
}
