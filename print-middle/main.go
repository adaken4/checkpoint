package main

import "os"

func main() {
	args := os.Args[1:]
	middle := len(args) / 2
	if len(args) == 0 {
		os.Stdout.WriteString("\n")
	} else if len(args)%2 == 0 {
		os.Stdout.WriteString(args[middle-1] + " " + args[middle] + "\n")
	} else if len(args)%2 == 1 {
		os.Stdout.WriteString(args[middle] + "\n")
	}
}
