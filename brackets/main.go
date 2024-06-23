package main

import "os"

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	for _, v := range args {
		if len(v) == 0 || isValidExpression(v) {
			os.Stdout.WriteString("OK\n")
		} else {
			os.Stdout.WriteString("Error\n")
		}
	}
}

func isValidExpression(expression string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range expression {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
