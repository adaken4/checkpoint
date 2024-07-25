package checkpoint

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return "" // Only bases from 2 to 16 are supported
	}

	// Define characters for digits 0-9 and letters A-F (up to base 16)
	baseChars := "0123456789ABCDEF"

	// Handle negative numbers
	negative := false
	if value < 0 {
		negative = true
		value = -value
	}

	// Handle the case when value is 0
	if value == 0 {
		return "0"
	}

	result := ""
	for value > 0 {
		digit := value % base
		result = string(baseChars[digit]) + result
		value = value / base
	}

	if negative {
		result = "-" + result
	}

	return result
}
