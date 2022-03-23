package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for a := len(os.Args) - 1; a >= 1; a-- {
		for _, j := range os.Args[a] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
