package main

import "github.com/01-edu/z01"

func main() {
	let := 0
	var alaphabet string = ("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(alaphabet[let]))
		let = let + 1
	}
	z01.PrintRune('\n')
}
