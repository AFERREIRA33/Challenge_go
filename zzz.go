package piscine

import "os"

func Zzz(f func(string) int, a []string) {
	var result int
	var tab [2]int
	vf := false
	var r string
	v := "No division by 0"
	var q []byte
	var z []byte
	for i := 0; i < len(v); i++ {
		q = append(q, byte(v[i]))
	}
	tab[0] = Atoi(a[0])
	tab[1] = Atoi(a[2])
	if tab[0] > 9223372036854775806 || tab[1] > 9223372036854775806 || tab[0] < -9223372036854775808 || tab[1] < -9223372036854775808 {
		vf = true
	} else if a[1] == "-" {
		result = tab[0] - tab[1]
	} else if a[1] == "+" {
		result = tab[0] + tab[1]
	} else if a[1] == "*" {
		result = tab[0] * tab[1]
	} else if a[1] == "/" {
		if tab[1] == 0 {
			os.Stdout.Write(q)
			vf = true
		} else {
			result = tab[0] / tab[1]
		}
	} else if a[1] == "%" {
		if tab[1] == 0 {
			os.Stdout.Write(q)
			vf = true
		} else {
			result = tab[0] % tab[1]
		}
	}
	if !vf {
		r = Convertintstr(result)
		for i := 0; i < len(r); i++ {
			z = append(z, byte(r[i]))
		}
		os.Stdout.Write(z)
	}
}
