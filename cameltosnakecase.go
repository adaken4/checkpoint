package checkpoint

func CamelToSnakeCase(s string) string {
	result := ""
	if s == "" || containsNumber(s) || zikifuatana(s) || endsWithCapital(s) {
		return s
	}
	for i, v := range s {
		if i != 0 && v >= 'A' && v <= 'Z' {
			result+=string('_')
			result+=string(v)
		} else {
			result+=string(v)
		}
	}
	return result
}

func containsNumber(s string) bool {
	for _, v := range s {
		if v >= '0' && v <= '9' {
			return true
		}
	}
	return false
}

func zikifuatana(s string) bool {
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			if i != len(s)-1 && s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return true
			}
		}
	}
	return false
}

func endsWithCapital(s string) bool {
	return  s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z'
}
