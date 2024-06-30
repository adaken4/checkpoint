package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	first := args[0]
	second := args[1]

	ind := 0
	firstRunes := []rune(first)
	rebuiltFirst := ""

	for _, v := range second {
		if ind < len(first) && firstRunes[ind] == v {
			rebuiltFirst += string(v)
			ind++
		}
	}
	if rebuiltFirst == first {
		os.Stdout.WriteString(rebuiltFirst + "\n")
	}
}
