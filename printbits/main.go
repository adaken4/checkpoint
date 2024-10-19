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
	numStr := args[0]
	num, _ := checkpoint.Atoi(numStr)
	binaryVal := itoabase2(num)
	padding := 8 - len(binaryVal)
	for i := 0; i < padding; i++ {
		binaryVal = "0" + binaryVal
	}
	os.Stdout.WriteString(binaryVal + "\n")
}

func itoabase2(num int) string {
	result := ""
	for num > 0 {
		result = string('0'+num%2) + result
		num /= 2
	}
	return result
}
