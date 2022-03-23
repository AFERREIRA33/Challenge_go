package piscine

func SortWordArr(r []string) {
	v := ""
	for i := 1; i < len(r); i++ {
		for a := 0; a < len(r); a++ {
			if r[a] > r[i] {
				v = r[i]
				r[i] = r[a]
				r[a] = v
			}
		}
	}
}
