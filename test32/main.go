package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	z01.PrintRune(rune(a))
}
