package interpreter

import (
	"errors"
	"strings"
)

type LexemsMap = map[string]string

type Lexemes struct {
	Lexemes LexemsMap
	errors  Errors
}

func NewLexemes(errors Errors) Lexemes {
	m := LexemsMap{
		"(":  "LEFT_PAREN",
		")":  "RIGHT_PAREN",
		"{":  "LEFT_BRACE",
		"}":  "RIGHT_BRACE",
		"*":  "STAR",
		".":  "DOT",
		",":  "COMMA",
		"+":  "PLUS",
		"-":  "MINUS",
		";":  "SEMICOLON",
		":":  "COLON",
		"=":  "EQUAL",
		"==": "EQUAL_EQUAL",
	}

	return Lexemes{
		Lexemes: m,
		errors:  errors,
	}
}

func (l Lexemes) ResolveLexems(line string, pos int) (string, int, error) {
	currentLexeme := string(line[pos])
	var count int
	for count = CountPrefixInMapKeys(l.Lexemes, currentLexeme); count > 1; count = CountPrefixInMapKeys(l.Lexemes, currentLexeme) {
		if pos += 1; pos > len(line)-1 {
			if _, found := l.Lexemes[currentLexeme]; found {
				return currentLexeme, pos, nil
			}
			return "", pos, errors.New(l.errors.unexpectedChar)
		}
		currentLexeme = currentLexeme + string(line[pos])
	}
	if count == 0 {
		currentLexeme = currentLexeme[:len(currentLexeme)-1]
		if _, found := l.Lexemes[currentLexeme]; found {
			return currentLexeme, pos - 1, nil
		}
		return "", pos, errors.New(l.errors.unexpectedChar)
	}
	return currentLexeme, pos, nil
}

func CountPrefixInMapKeys(m LexemsMap, prefix string) int {
	count := 0
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			count++
		}
	}
	return count
}
