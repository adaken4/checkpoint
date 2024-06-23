package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := collectWords(os.Args[1])
	var reversed string
	for i := len(words) - 1; i >= 0; i-- {
		if len(reversed) != 0 {
			reversed += " "
		}
		reversed += words[i]
	}
	os.Stdout.WriteString(reversed + "\n")
}

func collectWords(s string) []string {
	var word string
	var result []string
	for _, v := range s {
		if v == '_' || (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			word += string(v)
		} else if v == ' ' {
			result = append(result, word)
			word = ""
		}
	}
	if len(word) != 0 {
		result = append(result, word)
	}
	return result
}
