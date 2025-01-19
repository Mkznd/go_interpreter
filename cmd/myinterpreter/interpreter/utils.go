package interpreter

import "slices"

func IsDigit(l rune) bool {
	digits := []rune{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	}

	return slices.Contains(digits, l)
}

func isNumber(s string) bool {
	for _, l := range s {
		if !IsDigit(l) {
			return false
		}
	}
	return true
}
