package main

import (
	"github.com/01-edu/z01"
)

func main() {
	small := 'z'
	cap := 'Y'
	for small >= 'a' {
		z01.PrintRune(small)
		z01.PrintRune(cap)
		small-=2
		cap-=2
	}
	z01.PrintRune('\n')
}