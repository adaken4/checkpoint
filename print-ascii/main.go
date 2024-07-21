package main

import (
	"os"

	"checkpoint"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}
	if len(args[0]) != 1 {
		return
	}

	Rune := []rune(args[0])[0]

	if !letter(Rune) {
		return
	}

	os.Stdout.WriteString(checkpoint.Itoa(int(Rune)) + "\n")
}

func letter(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}
