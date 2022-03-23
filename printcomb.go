package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i <= 55; i++ {
		for j := i + 1; j <= 56; j++ {
			for k := j + 1; k <= 57; k++ {
				if i == 55 && j == 56 && k == 57 {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
				} else {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
