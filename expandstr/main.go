package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	words := []string{}
	word := ""
	for _, v := range arg {
		if v != ' ' {
			word += string(v)
		} else {
			if word != "" {
				words = append(words, word)
			}
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	for i, v := range words {
		if i != len(words)-1 {
			os.Stdout.Write([]byte(v))
			os.Stdout.Write([]byte("   "))
		} else {
			os.Stdout.Write([]byte(v))
			os.Stdout.Write([]byte("\n"))
		}
	}
}
