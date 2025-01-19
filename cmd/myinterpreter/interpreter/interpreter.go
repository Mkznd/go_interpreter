package interpreter

import "io"

type Interpreter struct {
	tokenizer Tokenizer
	parser    Parser
}

func (i Interpreter) Parse(tokens []Token) []Expression {
	return i.parser.Parse(tokens)
}

func (i Interpreter) DisplayTokens(tokens []Token, writer io.Writer) {
	i.tokenizer.Display(tokens, writer)
}

func (i Interpreter) Tokenize(buf []byte) ([]Token, int) {
	return i.tokenizer.Tokenize(buf)
}

func NewInterpreter(t Tokenizer, p Parser) Interpreter {
	return Interpreter{
		tokenizer: t,
		parser:    p,
	}
}
