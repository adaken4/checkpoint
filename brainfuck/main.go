package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	source := os.Args[1]
	tape := make([]byte, 2048)
	pointer := 0
	bracketMap := make(map[int]int)
	stack := []int{}

	for i := 0; i < len(source); i++ {
		switch source[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bracketMap[start] = i
			bracketMap[i] = start
		}
	}

	// Second pass: interpret the code
	for i := 0; i < len(source); i++ {
		switch source[i] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			tape[pointer]++
		case '-':
			tape[pointer]--
		case '.':
			z01.PrintRune(rune(tape[pointer]))
		case '[':
			if tape[pointer] == 0 {
				i = bracketMap[i]
			}
		case ']':
			if tape[pointer] != 0 {
				i = bracketMap[i]
			}
		}
	}
}
