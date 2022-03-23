package piscine

func Join(strs []string, sep string) string {
	var z string
	for i := 0; i < len(strs); i++ {
		z = z + strs[i]
		if i == len(strs)-1 {
		} else {
			z = z + sep
		}
	}
	return z
}
