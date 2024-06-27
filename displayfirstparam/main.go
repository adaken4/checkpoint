package main

import "os"

func main() {
	if len(os.Args) == 1 {
		return
	}
	os.Stdout.WriteString(os.Args[1] + "\n")
}
