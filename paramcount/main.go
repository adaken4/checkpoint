package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stdout.WriteString("0\n")
	} else {
		n := len(args)
		params := ""
		for n > 0 {
			params = string(n%10+'0') + params
			n /= 10
		}
		os.Stdout.WriteString(params + "\n")
	}
}
