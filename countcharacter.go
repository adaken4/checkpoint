package checkpoint

func CountChar(str string, c rune) int {
	count := 0
	for _, v := range str {
		if v == c {
			count++
		}
	}
	return count
}
