package checkpoint

func IsCapitalized(s string) bool {
	var check bool
	words := []string{}
	word := ""
	for _, v := range s {
		if v != ' ' {
			word += string(v)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	for _, v := range words {
		vRunes := []rune(v)
		if vRunes[0] >= 'a' && vRunes[0] <= 'z' {
			check = false
		} else {
			check = true
		}
	}
	return check
}
