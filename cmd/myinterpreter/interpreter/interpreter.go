package interpreter

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Parser struct {
	lexems Lexems
	errors Errors
	ignore []string
}

func (p Parser) Scan(buf []byte) {
	code := 0
	lines := strings.Split(string(buf[:]), "\n")
	for i, line := range lines {
		for _, letter := range line {
			if value, found := p.lexems[string(letter)]; found {
				fmt.Println(value, string(letter), "null")
			} else if slices.Contains(p.ignore, string(letter)) {
				continue
			} else {
				_, err := fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", i+1, p.errors.unexpectedChar, string(letter))
				if err != nil {
					return
				}
				code = 65
			}
		}
	}
	fmt.Println("EOF  null")
	os.Exit(code)
}

func NewParser(lexems Lexems, errors Errors, ignore []string) Parser {
	return Parser{
		lexems: lexems,
		errors: errors,
		ignore: ignore,
	}
}
