package interpreter

type Expression struct {
	tokens []Token
}

func NewExpression(tokens []Token) Expression {
	return Expression{
		tokens: tokens,
	}
}

type Parser struct {
}

func (p Parser) Parse(tokens []Token) []Expression {
	return []Expression{NewExpression(tokens)}
}

func NewParser() Parser {
	return Parser{}
}
