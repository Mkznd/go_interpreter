package interpreter

import "fmt"

type Lexems = map[string]string

func NewLexems() Lexems {
	return Lexems{
		"(": "LEFT_PAREN",
		")": "RIGHT_PAREN",
		"{": "LEFT_BRACE",
		"}": "RIGHT_BRACE",
		"*": "STAR",
		".": "DOT",
		",": "COMMA",
		"+": "PLUS",
		"-": "MINUS",
		";": "SEMICOLON",
		":": "COLON",
	}
}

type Parser struct {
	lexems Lexems
}

func (p Parser) Scan(buf []byte) {
	for _, letter := range buf {
		if value, found := p.lexems[string(letter)]; found {
			fmt.Println(value, string(letter), "null")
		}
	}
	fmt.Println("EOF  null")
}

func NewParser(lexems Lexems) Parser {
	return Parser{
		lexems: lexems,
	}
}
