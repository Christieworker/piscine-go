package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for i, w := range arg {
		if i > 0 {
			for _, word := range w {
				z01.PrintRune(word)
			}
			z01.PrintRune('\n')
		}
	}
}
