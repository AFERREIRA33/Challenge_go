package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	z01.PrintRune(rune(div))
	z01.PrintRune('\n')
	z01.PrintRune(rune(mod))
}
