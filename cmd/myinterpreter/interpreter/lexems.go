package interpreter

import (
	"errors"
	"slices"
	"strings"
)

type LexemsMap = map[string]string

type Lexemes struct {
	Lexemes LexemsMap
	errors  Errors
	ignore  []string
}

func NewLexemes(errors Errors, ignore []string) Lexemes {
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
		"!":  "BANG",
		"!=": "BANG_EQUAL",
		"<":  "LESS",
		">":  "GREATER",
		"<=": "LESS_EQUAL",
		">=": "GREATER_EQUAL",
		"/":  "SLASH",
		"//": "COMMENT",
	}

	return Lexemes{
		Lexemes: m,
		errors:  errors,
		ignore:  ignore,
	}
}

type Token struct {
	TokenType string
	Lexeme    string
	Literal   string
}

func NewToken(tokenType string, lexeme string, literal string) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
	}
}

func (l Lexemes) ResolveLexems(line string, pos int) (Token, int, error) {
	currentLexeme := string(line[pos])
	whitespaces := []string{
		" ", "\n", "\r", "\t",
	}
	matched, currentLexeme, pos := l.skipWhitespaces(line, pos, whitespaces, currentLexeme)
	if matched && pos == len(line)-1 {
		return Token{}, pos, nil
	}
	if isNumber(currentLexeme) {
		return l.ExtractNumberLiteral(line, pos)
	}
	if currentLexeme == "\"" {
		return l.ExtractStringLiteral(line, pos)
	}
	var count int
	for count = CountPrefixInMapKeys(l.Lexemes, currentLexeme); count > 1; count = CountPrefixInMapKeys(l.Lexemes, currentLexeme) {
		if pos += 1; pos > len(line)-1 {
			if _, found := l.Lexemes[currentLexeme]; found {
				return NewToken(l.Lexemes[currentLexeme], currentLexeme, "null"), pos, nil
			}
			return Token{Lexeme: currentLexeme}, pos, errors.New(l.errors.unexpectedChar)
		}
		currentLexeme = currentLexeme + string(line[pos])
	}

	if count == 0 {
		previousLexeme := currentLexeme[:len(currentLexeme)-1]
		if _, found := l.Lexemes[previousLexeme]; found {
			return NewToken(l.Lexemes[previousLexeme], previousLexeme, "null"), pos - 1, nil
		}
		return Token{Lexeme: currentLexeme}, pos, errors.New(l.errors.unexpectedChar)
	}
	return NewToken(l.Lexemes[currentLexeme], currentLexeme, "null"), pos, nil
}

func (l Lexemes) skipWhitespaces(line string, pos int, whitespaces []string, currentLexeme string) (bool, string, int) {
	matched := slices.Contains(whitespaces, currentLexeme)
	for matched && pos < len(line)-1 {
		pos++
		currentLexeme = string(line[pos])
		matched = slices.Contains(whitespaces, currentLexeme)
	}
	return matched, currentLexeme, pos
}

func (l Lexemes) ExtractNumberLiteral(s string, pos int) (Token, int, error) {
	if !IsDigit(rune(s[pos])) {
		return Token{Lexeme: string(s[pos])}, pos, errors.New(l.errors.unexpectedChar)
	}

	res := string(s[pos])
	for pos = pos + 1; pos < len(s) && IsDigit(rune(s[pos])); pos++ {
		res += string(s[pos])
	}

	return NewToken("NUMBER", res, res+".0"), pos, nil
}

func (l Lexemes) ExtractStringLiteral(s string, pos int) (Token, int, error) {
	if s[pos] != '"' {
		return Token{Lexeme: string(s[pos])}, pos, errors.New(l.errors.unexpectedChar)
	}
	res := "\""
	for pos = pos + 1; pos < len(s) && s[pos] != '"'; pos++ {
		res += string(s[pos])
	}

	if pos >= len(s) {
		return Token{Lexeme: res}, pos, errors.New(l.errors.unterminatedString)
	}
	return NewToken("STRING", res+"\"", strings.TrimPrefix(res, "\"")), pos, nil
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
