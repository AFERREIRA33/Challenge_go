package piscine

func Split(s, sep string) []string {
	var tab []string
	r := []rune(s)
	rp := []rune(sep)
	var m string
	c := 0
	for i := 0; i < len(s); i++ {
		if r[i] == rp[0] {
			for a := 0; a < len(sep); a++ {
				if r[a+i] == rp[a] {
					c++
				} else {
					m = m + string(r[i])
					break
				}
				if c == len(sep) {
					i = i + len(sep) - 1
					tab = append(tab, m)
					m = ""
					c = 0
				}
			}
		} else {
			m = m + string(r[i])
		}
		c = 0
	}
	tab = append(tab, m)
	return tab
}
