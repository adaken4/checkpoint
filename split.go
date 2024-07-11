package checkpoint

func Split(s, sep string) []string {
	var result []string
	sepRunes := []rune(sep)
	word := ""
	j := 0
	for _, v := range s {
		if j < len(sepRunes) && v == sepRunes[j] {
			j++
			if j == len(sepRunes) {
				result = append(result, word)
				word = ""
				j = 0
			}
		} else {
			if j > 0 {
				word += string(sepRunes[:j])
				j = 0
			}
			word += string(v)
		}
	}
	if len(word) != 0 {
		result = append(result, word)
	}
	return result
}
