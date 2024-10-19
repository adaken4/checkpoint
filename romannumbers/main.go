package main

import (
	"os"

	"checkpoint"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := checkpoint.Atoi(os.Args[1])
	if err {
		os.Stdout.WriteString(checkpoint.Itoa(romanToInt(os.Args[1]))+"\n")
		return
	}
	if num < 1 || num > 3999 {
		os.Stdout.WriteString("ERROR: cannot convert to roman digit\n")
		return
	}
	result, breakdown := intToRoman(num)
	os.Stdout.WriteString(breakdown + "\n")
	os.Stdout.WriteString(result + "\n")
}

func intToRoman(num int) (string, string) {
	var result, breakdown string

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calculations := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	for i, v := range values {
		for num >= v {
			num -= v
			result += symbols[i]
			if len(breakdown) > 0 {
				breakdown += "+"
			}
			breakdown += calculations[i]
		}
	}

	return result, breakdown
}

func romanToInt(s string) int {
	num := 0
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if i+1<len(s) && romanValues[s[i+1]] > romanValues[s[i]] {
			num+=romanValues[s[i+1]]-romanValues[s[i]]
			i++
		} else {
			num+=romanValues[s[i]]
		}
		// if s[i] == 'I' {
		// 	if i+1 < len(s) && s[i+1] == 'V' {
		// 		num += 4
		// 		i++
		// 	} else if i+1 < len(s) && s[i+1] == 'X' {
		// 		num += 9
		// 		i++
		// 	} else {
		// 		num += 1
		// 	}
		// } else if s[i] == 'X' {
		// 	if i+1 < len(s) && s[i+1] == 'L' {
		// 		num += 40
		// 		i++
		// 	} else if i+1 < len(s) && s[i+1] == 'C' {
		// 		num += 90
		// 		i++
		// 	} else {
		// 		num += 10
		// 	}
		// } else if s[i] == 'C' {
		// 	if i+1 < len(s) && s[i+1] == 'D' {
		// 		num += 400
		// 		i++
		// 	} else if i+1 < len(s) && s[i+1] == 'M' {
		// 		num += 900
		// 		i++
		// 	} else {
		// 		num += 100
		// 	}
		// } else if s[i] == 'V' {
		// 	num += 5
		// } else if s[i] == 'L' {
		// 	num += 50
		// } else if s[i] == 'D' {
		// 	num += 500
		// } else if s[i] == 'M' {
		// 	num += 1000
		// }
	}
	return num
}
