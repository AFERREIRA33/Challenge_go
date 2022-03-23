package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for a := 1; a < len(os.Args); a++ {
		for _, j := range os.Args[a] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
