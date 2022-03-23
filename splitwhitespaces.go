package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	r := []rune(s)
	var m string
	for i := 0; i < len(s); i++ {
		if r[i] == 32 || r[i] == '\n' || r[i] == 9 {
			if r[i+1] == 32 || r[i+1] == '\n' || r[i+1] == 9 {
			} else {
				tab = append(tab, m)
				m = ""
			}
		} else {
			m = m + string(r[i])
		}
	}
	tab = append(tab, m)
	return tab
}
