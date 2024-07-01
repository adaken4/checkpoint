package checkpoint

func Atoi(s string) (int, bool) {
	if Invalidinput(s) {
		return 0, true
	}
	result := 0
	sign := 1
	unsigned := s
	if s[0] == '-' {
		sign *= -1
		unsigned = s[1:]
	} else if s[0] == '+' {
		unsigned = s[1:]
	}
	for _, v := range unsigned {
		result *= 10
		result += int(v - '0')
	}
	result *= sign
	if s[0] == '-' && result > 0 {
		return 0, true
	}
	if (s[0] == '+' || s[0] != '-') && result < 0 {
		return 0, true
	}
	return result, false
}

func Invalidinput(s string) bool {
	for i, v := range s {
		if i == 0 && (v == '-' || v == '+') {
			continue
		} else if v < '0' || v > '9' {
			return true
		}
	}
	return false
}
