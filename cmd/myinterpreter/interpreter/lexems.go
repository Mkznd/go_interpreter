package interpreter

import (
	"errors"
	"slices"
	"strings"
)

type Lexemes struct {
	Lexemes       LexemesMap
	ReservedWords LexemesMap
	errors        Errors
}

func NewLexemes(errors Errors) Lexemes {
	m := GetMiscLexemes()
	reservedWords := GetReservedWords()

	return Lexemes{
		Lexemes:       m,
		ReservedWords: reservedWords,
		errors:        errors,
	}
}

func (l Lexemes) ResolveLexemes(line string, pos int) (Token, int, error) {
	currentLexeme := string(line[pos])
	whitespaces := []string{
		" ", "\n", "\r", "\t",
	}
	matched, currentLexeme, pos := l.skipWhitespaces(line, pos, whitespaces, currentLexeme)
	if matched && pos == len(line)-1 {
		return Token{}, pos, nil
	}
	if IsNumber(currentLexeme) {
		return l.ExtractNumberLiteral(line, pos)
	}
	if currentLexeme == "\"" {
		return l.ExtractStringLiteral(line, pos)
	}
	if IsIdentifier(currentLexeme) {
		return l.ExtractIdentifierLiteral(line, pos)
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

func (l Lexemes) ExtractIdentifierLiteral(s string, pos int) (Token, int, error) {
	if !IsIdentifierSymbol(rune(s[pos])) {
		return Token{Lexeme: string(s[pos])}, pos, errors.New(l.errors.unexpectedChar)
	}
	res := string(s[pos])
	for pos = pos + 1; pos < len(s) && IsIdentifier(res+string(s[pos])); pos++ {
		res += string(s[pos])
	}

	if idType, found := l.ReservedWords[res]; found {
		return NewToken(idType, res, "null"), pos - 1, nil
	}

	return NewToken("IDENTIFIER", res, "null"), pos - 1, nil
}

func (l Lexemes) ExtractNumberLiteral(s string, pos int) (Token, int, error) {
	hadDot := false
	if !IsDigit(rune(s[pos])) {
		return Token{Lexeme: string(s[pos])}, pos, errors.New(l.errors.unexpectedChar)
	}

	res := string(s[pos])
	for pos = pos + 1; pos < len(s); pos++ {
		if IsDigit(rune(s[pos])) {
			res += string(s[pos])
		} else if s[pos] == '.' && !hadDot {
			res += "."
			hadDot = true
		} else {
			break
		}
	}
	if hadDot {
		literal := strings.TrimRight(res, "0")
		if literal[len(literal)-1] == '.' {
			literal += "0"
		}
		return NewToken("NUMBER", res, literal), pos - 1, nil
	} else {
		return NewToken("NUMBER", res, res+".0"), pos - 1, nil
	}
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

func CountPrefixInMapKeys(m LexemesMap, prefix string) int {
	count := 0
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			count++
		}
	}
	return count
}
