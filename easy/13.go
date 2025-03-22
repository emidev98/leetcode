package easy

import (
	"strings"
)

// https://leetcode.com/problems/roman-to-integer/solutions/6566138/13-roman-to-integer-by-emidev98-7op3/
func RomanToInt(input string) int {
	dictionary := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	splitInput := strings.Split(input, "")
	acc := 0

	for i := 0; i < len(splitInput); i++ {
		currentRomanNumberValue := dictionary[splitInput[i]]
		nextRomanNumberIndex := i + 1
		if nextRomanNumberIndex == len(splitInput) {
			acc = acc + currentRomanNumberValue
			return acc
		}
		nextRomanNumberValue := dictionary[splitInput[nextRomanNumberIndex]]

		if nextRomanNumberValue <= currentRomanNumberValue {
			acc = acc + currentRomanNumberValue
		} else {
			acc = acc - currentRomanNumberValue
		}
	}

	return acc
}
