//go:build wireinject
// +build wireinject

package di

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
	"github.com/google/wire"
)

func InitializeParser() (interpreter.Parser, error) {
	wire.Build(interpreter.NewParser, interpreter.NewParenthesesScanner, interpreter.NewLexems)
	return interpreter.Parser{}, nil
}
