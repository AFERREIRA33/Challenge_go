package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	z01.PrintRune(rune(a))
	z01.PrintRune('\n')
	z01.PrintRune(rune(b))
}
