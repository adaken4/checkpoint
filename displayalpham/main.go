package main

import "github.com/01-edu/z01"

func main() {
	small := 'a'
	cap := 'B'
	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			z01.PrintRune(small)
			small += 2
		} else {
			z01.PrintRune(cap)
			cap += 2
		}
	}
	z01.PrintRune('\n')
}
