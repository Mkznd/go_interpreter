//go:build wireinject
// +build wireinject

package di

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
	"github.com/google/wire"
)

func InitializeInterpreter() (interpreter.Interpreter, error) {
	wire.Build(interpreter.NewInterpreter, interpreter.NewParser, interpreter.NewTokenizer, interpreter.NewLexemes, interpreter.NewErrors)
	return interpreter.Interpreter{}, nil
}
