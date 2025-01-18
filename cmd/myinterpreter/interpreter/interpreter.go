package interpreter

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Parser struct {
	lexemes Lexemes
	errors  Errors
	ignore  []string
}

func (p Parser) Scan(buf []byte) {
	code := 0
	buf = p.clean(buf)
	lines := strings.Split(string(buf[:]), "\n")
	for i, line := range lines {
		for pos := 0; pos < len(line); pos++ {
			token, newPos, err := p.lexemes.ResolveLexems(line, pos)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", i+1, p.errors.unexpectedChar, token.Lexeme)
				code = 65
				continue
			}
			if token.Lexeme == "//" {
				break
			}
			pos = newPos
			fmt.Println(token.TokenType, token.Lexeme, token.Literal)
		}
	}
	fmt.Println("EOF  null")
	os.Exit(code)
}

func (p Parser) clean(buf []byte) []byte {
	return slices.DeleteFunc(buf, func(b byte) bool {
		return slices.Contains(p.ignore, string(b))
	})
}

func NewParser(lexemes Lexemes, errors Errors, ignore []string) Parser {
	return Parser{
		lexemes: lexemes,
		errors:  errors,
		ignore:  ignore,
	}
}
