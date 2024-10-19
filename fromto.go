package checkpoint

func FromTo(from, to int) string {
	result := ""
	if (from >= 0 && from <= 99) && (to >= 0 && to <= 99) {
		for from < to {
			str := Itoa(from)
			if len(str) < 2 {
				str = "0" + str
			}
			result += str + ", "
			from++
		}
		for from > to {
			str := Itoa(from)
			if len(str) < 2 {
				str = "0" + str
			}
			result += str + ", "
			from--
		}
		if from == to {
			str := Itoa(from)
			if len(str) < 2 {
				str = "0" + str
			}
			result += str + "\n"
		}
		return result
	}
	return "Invalid\n"
}
