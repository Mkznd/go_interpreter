package interpreter

import (
	"slices"
)

func IsDigit(s rune) bool {
	digits := []rune{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	}

	return slices.Contains(digits, s)
}

func IsNumber(s string) bool {
	for _, l := range s {
		if !IsDigit(l) {
			return false
		}
	}
	return true
}

func IsIdentifierSymbol(s rune) bool {
	allowedSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	return slices.Contains([]rune(allowedSymbols), s)
}

func IsIdentifier(s string) bool {
	for _, l := range s {
		if !IsIdentifierSymbol(l) {
			return false
		}
	}
	return true
}
