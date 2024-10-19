package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	for _, arg := range args {
		result := ""
		for _, v := range arg {
			if v >= 'a' && v <= 'z' {
				result += string(v)
			}
		}
	}
}

func lower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func upper(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}
