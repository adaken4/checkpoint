package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	var mirrored string
	for _, v := range args[0] {
		if v >= 'a' && v <= 'z' {
			v = 'z' - (v - 'a')
			mirrored += string(v)
		} else if v >= 'A' && v <= 'Z' {
			v = 'Z' - (v - 'A')
			mirrored += string(v)
		} else {
			mirrored += string(v)
		}
	}
	os.Stdout.WriteString(mirrored + "\n")
}
