package interpreter

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Parser struct {
	lexemes Lexemes
	errors  Errors
}

func (p Parser) Scan(buf []byte) {
	code := 0
	start := time.Now()
	lines := strings.Split(string(buf[:]), "\n")
	for i, line := range lines {
		for pos := 0; pos < len(line); pos++ {
			token, newPos, err := p.lexemes.ResolveLexems(line, pos)
			pos = newPos
			if err != nil {
				if err.Error() == p.errors.unexpectedChar {
					fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", i+1, p.errors.unexpectedChar, token.Lexeme)
					code = 65
					continue
				} else if err.Error() == p.errors.unterminatedString {
					fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", i+1, p.errors.unterminatedString)
					code = 65
					break
				}
			}
			if token.Lexeme == "" {
				break
			}
			if token.Lexeme == "//" {
				break
			}
			fmt.Println(token.TokenType, token.Lexeme, token.Literal)
		}
	}
	fmt.Println("EOF  null")
	elapsed := time.Since(start)
	fmt.Fprintln(os.Stderr, "Parsing took", elapsed.Nanoseconds(), "nanosecs")
	os.Exit(code)
}

func NewParser(lexemes Lexemes, errors Errors, ignore []string) Parser {
	return Parser{
		lexemes: lexemes,
		errors:  errors,
	}
}
