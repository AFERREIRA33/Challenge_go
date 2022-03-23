package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	var stock []int
	var ba []rune
	r := 0
	n := nbr
	lounge := len(base)
	var verif bool
	for a := 0; a < len(base); a++ {
		ba = append(ba, rune(base[a]))
	}
	for index, i := range base {
		for dedex, j := range base {
			if i == j && index != dedex {
				verif = true
				break
			} else {
				verif = false
			}
		}
		if i == '-' || i == '+' {
			verif = true
		}
		if verif {
			break
		}
	}
	if (len(base) < 2 && base != "01") || verif == true {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if n < 0 {
			for n > 0 || n < 0 {
				r = n % lounge
				n = n / lounge
				if r == 0 {
					stock = append(stock, 0)
				} else {
					stock = append(stock, -r)
				}
			}
		} else {
			for n > 0 || n < 0 {
				r = n % lounge
				n = n / lounge
				if r == 0 {
					stock = append(stock, 0)
				} else {
					stock = append(stock, r)
				}
			}
		}
		if nbr < 0 {
			z01.PrintRune('-')
			for t := len(stock) - 1; t >= 0; t-- {
				z01.PrintRune(ba[stock[t]])
			}
		} else {
			for t := len(stock) - 1; t >= 0; t-- {
				z01.PrintRune(ba[stock[t]])
			}
		}
	}
}
