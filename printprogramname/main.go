package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, j := range os.Args[0] {
		if i == 0 || i == 1 && (os.Args[0][0] == 46 && os.Args[0][1] == 47) {
		} else {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
