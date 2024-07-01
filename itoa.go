package checkpoint

func Itoa(n int) string {
	result := ""
	nb := n
	if nb == 0 {
		return "0"
	}
	if nb < 0 {
		nb *= -1
	}
	for nb > 0 {
		result = string('0'+nb%10) + result
		nb /= 10
	}
	if n < 0 {
		result = "-" + result
	}
	return result
}
