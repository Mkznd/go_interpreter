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

func IsIdentifierStartingSymbol(s rune) bool {
	allowedSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	return slices.Contains([]rune(allowedSymbols), s)
}

func IsIdentifierSymbol(s rune) bool {
	allowedSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_1234567890"

	return slices.Contains([]rune(allowedSymbols), s)
}

func IsIdentifier(s string) bool {
	if len(s) == 0 || !IsIdentifierStartingSymbol(rune(s[0])) {
		return false
	}
	for _, l := range s[1:] {
		if !IsIdentifierSymbol(l) {
			return false
		}
	}
	return true
}
