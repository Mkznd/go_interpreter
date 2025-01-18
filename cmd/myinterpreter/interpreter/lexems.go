package interpreter

import (
	"errors"
	"regexp"
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
	matched, _ := regexp.MatchString(`[\s\r]`, currentLexeme)
	for matched && pos < len(line)-1 {
		pos++
		currentLexeme = string(line[pos])
		matched, _ = regexp.MatchString(`[\s\r]`, currentLexeme)
	}
	if matched && pos == len(line)-1 {
		return Token{}, pos, nil
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
