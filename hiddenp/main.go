package main

import "os"

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if hiddenP(s1, s2) {
		os.Stdout.WriteString("1\n")
	} else {
		os.Stdout.WriteString("0\n")
	}
}

func hiddenP(s1, s2 string) bool {
	if s1 == "" {
		return true
	}
	j := 0
	s1Runes := []rune(s1)
	for _, v := range s2 {
		if v == s1Runes[j] {
			j++
		}
		if j == len(s1) {
			return true
		}
	}
	return j == len(s1)
}
