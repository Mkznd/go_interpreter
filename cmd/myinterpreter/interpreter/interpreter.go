package interpreter

type Lexems = map[string]string

func NewLexems() Lexems {
	return Lexems{
		"(": "LEFT_PAREN",
		")": "RIGHT_PAREN",
	}
}

type Parser struct {
	lexems             Lexems
	parenthesesScanner ParenthesesScanner
}

func (p Parser) ScanParentheses(buf []byte) {
	p.parenthesesScanner.scan(buf)
}

func NewParser(lexems Lexems, parenthesesScanner ParenthesesScanner) Parser {
	return Parser{
		lexems:             lexems,
		parenthesesScanner: parenthesesScanner,
	}
}
