package interpreter

type LexemesMap = map[string]string

func GetMiscLexemes() LexemesMap {
	return LexemesMap{
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
}

func GetReservedWords() LexemesMap {
	return LexemesMap{
		"and":    "AND",
		"class":  "CLASS",
		"else":   "ELSE",
		"false":  "FALSE",
		"for":    "FOR",
		"fun":    "FUN",
		"if":     "IF",
		"nil":    "NIL",
		"or":     "OR",
		"print":  "PRINT",
		"return": "RETURN",
		"super":  "SUPER",
		"this":   "THIS",
		"true":   "TRUE",
		"var":    "VAR",
		"while":  "WHILE",
	}
}
