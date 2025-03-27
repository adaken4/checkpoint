package main

import "os"

const optionsString = "abcdefghijklmnopqrstuvwxyz"

func main() {
	args := os.Args[1:]
	// No arguments: print valid options.
	if len(args) == 0 {
		os.Stdout.Write([]byte("options: " + optionsString + "\n"))
		return
	}

	var options int = 0

	for _, arg := range args {
		// Check: must start with '-' and have at least one option character.
		if len(arg) < 2 || arg[0] != '-' {
			os.Stdout.Write([]byte("Invalid Option\n"))
			return
		}

		// Process each character after the '-' sign.
		for i := 1; i < len(arg); i++ {
			c := arg[i]
			// 'h' anywhere should print all options.
			if c == 'h' {
				os.Stdout.Write([]byte("options: " + optionsString + "\n"))
				return
			}

			// Check if c is a valid option.
			valid := false
			for j := 0; j < len(optionsString); j++ {
				if optionsString[j] == c {
					valid = true
					// Set the bit corresponding to c.
					options |= (1 << (c - 'a'))
					break
				}
			}
			if !valid {
				os.Stdout.Write([]byte("Invalid Option\n"))
				return
			}
		}
	}

	// Build a string with four groups of 8 bits.
	out := ""
	for shift := 24; shift >= 0; shift -= 8 {
		b := byte((options >> uint(shift)) & 0xFF)
		out += byteToBinary(b)
		if shift != 0 {
			out += " "
		}
	}
	out += "\n"
	os.Stdout.Write([]byte(out))
}

// byteToBinary converts a byte to an 8-character string representing its binary value.
func byteToBinary(b byte) string {
	s := ""
	for i := 7; i >= 0; i-- {
		if b&(1<<uint(i)) != 0 {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}
