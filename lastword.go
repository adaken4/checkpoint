package checkpoint

// func LastWord(s string) string {
// 	var lstwrd string
// 	sRunes := []rune(s)
// 	for i := len(s) - 1; i > 0; i-- {
// 		if sRunes[i] == ' ' {
// 			s = s[:i]
// 		} else {
// 			break
// 		}
// 	}
// 	for _, v := range s {
// 		if v != ' ' {
// 			lstwrd += string(v)
// 		} else {
// 			lstwrd = ""
// 		}
// 	}
// 	return lstwrd + "\n"
// }

func LastWord(s string) string {
	sRunes := []rune(s)
	start, end := -1, -1

	for i := len(sRunes) - 1; i >= 0; i-- {
		if sRunes[i] != ' ' {
			if end == -1 {
				end = i + 1
			} else if i == 0 && start == -1 {
				start = start + 1
			}
		} else if end != -1 {
			start = i + 1
			break
		}
	}

	if start == -1 || end == -1 {
		return "\n"
	}

	return s[start:end] + "\n"
}
