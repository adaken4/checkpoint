package main

import (
	"os"

	"checkpoint"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if checkpoint.Invalidinput(args[0]) || checkpoint.Invalidinput(args[2]) {
		return
	}
	op := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"%": true,
	}
	if !op[args[1]] {
		return
	}
	n1, overflow := checkpoint.Atoi(args[0])
	if overflow {
		return
	}
	n2, overflow := checkpoint.Atoi(args[2])
	if overflow {
		return
	}
	switch args[1] {
	case "+":
		result := add(n1, n2)
		if result < 0 && (n1 > 0 && n2 > 0) {
			return
		}
		os.Stdout.WriteString(checkpoint.Itoa(result) + "\n")
	case "-":
		result := sub(n1, n2)
		if result > 0 && n1 < 0 && n2 < 0 {
			return
		}
		os.Stdout.WriteString(checkpoint.Itoa(result) + "\n")
	case "*":
		result := mul(n1, n2)
		if n2 != 0 && div(result, n2) != n1 {
			return
		}
		os.Stdout.WriteString(checkpoint.Itoa(result) + "\n")
	case "/":
		if n2 == 0 {
			os.Stdout.WriteString("no division by zero\n")
			return
		}
		os.Stdout.WriteString(checkpoint.Itoa(div(n1, n2)) + "\n")
	case "%":
		if n2 == 0 {
			os.Stdout.WriteString("no modulo by zero\n")
			return
		}
		os.Stdout.WriteString(checkpoint.Itoa(mod(n1, n2)) + "\n")
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}
