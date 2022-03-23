package piscine

func ConcatParams(args []string) string {
	var txt string
	z := '\n'
	for i := 0; i < len(args); i++ {
		txt = txt + args[i]
		if i != len(args)-1 {
			txt = txt + string(z)
		}
	}
	return txt
}
