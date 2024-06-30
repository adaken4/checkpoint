package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	rslt := ""
	for _, v := range args[0] {
		if v >= 'a' && v <= 'z' {
			v += 13
			if v > 'z' {
				v = (v - 'z') + ('a' - 1)
			}
			rslt += string(v)
		} else if v >= 'A' && v <= 'Z' {
			v += 13
			if v > 'Z' {
				v = (v - 'Z') + ('A' - 1)
			}
			rslt += string(v)
		} else {
			rslt += string(v)
		}
	}
	os.Stdout.WriteString(rslt + "\n")
}
