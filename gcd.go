package checkpoint

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	// var i uint
	// if a > b {
	// 	i = a
	// } else {
	// 	i = b
	// }
	// for i >= 1 {
	// 	if a%i == 0 && b%i == 0 {
	// 		return i
	// 	}
	// 	i--
	// }
	// return 0

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
