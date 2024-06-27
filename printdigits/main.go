package main

import "os"

func main() {
	for i := '0'; i <= '9'; i++ {
		os.Stdout.WriteString(string(i))
	}
	os.Stdout.WriteString("\n")
}
