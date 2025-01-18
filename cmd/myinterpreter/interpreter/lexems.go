package interpreter

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
		"=": "EQUAL",
	}
}
