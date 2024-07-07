package main

import (
	"os"

	"checkpoint"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := checkpoint.Atoi(os.Args[1])
	if num > 0 && num&(num-1) == 0 {
		os.Stdout.WriteString("true\n")
	} else {
		os.Stdout.WriteString("false\n")
	}
}
