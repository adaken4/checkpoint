package checkpoint

func ReverseBits(octet byte) byte {
	var reversed byte
	for i := 0; i < 8; i++ {
		reversed |= (octet & (1 << i)) >> i << (7 - i)
	}
	return reversed
}
