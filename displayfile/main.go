package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) == 2 {
		file, _ := os.Open("quest8.txt")
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	} else {
		fmt.Println("Too many arguments")
	}
}
