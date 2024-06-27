package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stdout.WriteString("0\n")
	} else {
		os.Stdout.WriteString(string(len(args)+'0') + "\n")
	}
}
