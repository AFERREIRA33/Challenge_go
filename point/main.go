package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := 0; i < len(points.x); i++ {
		z01.PrintRune(rune(points.x[i]))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := 0; i < len(points.y); i++ {
		z01.PrintRune(rune(points.y[i]))
	}
	z01.PrintRune('\n')
}
