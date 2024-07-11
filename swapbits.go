package checkpoint

func SwapBits(octet byte) byte {
	highNibble := octet & 0xf0
	lowNibble := octet & 0x0f
	highNibble = highNibble >> 4
	lowNibble = lowNibble << 4
	return lowNibble | highNibble
}
