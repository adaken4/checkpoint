package checkpoint

func FindPrevPrime(nb int) int {
	for nb > 1 {
		if IsPrime(nb) {
			return nb
		}
		nb--
	}
	return 0
}
