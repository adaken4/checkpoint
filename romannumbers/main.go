package main

import (
	"os"

	"checkpoint"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := checkpoint.Atoi(os.Args[1])
	if num < 1 || num > 3999 {
		os.Stdout.WriteString("ERROR: cannot convert to roman digit\n")
		return
	}
	result, breakdown := toRoman(num)
	os.Stdout.WriteString(breakdown + "\n")
	os.Stdout.WriteString(result + "\n")
}

func toRoman(num int) (string, string) {
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
