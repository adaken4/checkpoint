package checkpoint

// AtoiBase converts a string in a given base to an integer.
func AtoiBase(s string, base string) int {
	// Check if base is valid
	if !isValidBase(base) {
		return 0
	}

	// Create a map to store the value of each character in the base
	charValue := make(map[rune]int)
	for i, char := range base {
		charValue[char] = i
	}

	// Convert the string to an integer
	result := 0
	baseLen := len(base)
	for _, char := range s {
		value, exists := charValue[char]
		if !exists {
			// If the character is not in the base, return 0 (shouldn't happen due to problem constraints)
			return 0
		}
		result = result*baseLen + value
	}

	return result
}

// isValidBase checks if the given base is valid
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
