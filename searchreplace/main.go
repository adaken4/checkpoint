package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	sentence := args[0]
	search := args[1]
	replace := args[2]
	result := ""

	for _, v := range sentence {
		if v == []rune(search)[0] {
			result += replace
		} else {
			result += string(v)
		}
	}
	os.Stdout.WriteString(result + "\n")
}
