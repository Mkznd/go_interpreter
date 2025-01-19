//go:build wireinject
// +build wireinject

package di

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
	"github.com/google/wire"
)

func InitializeParser() (interpreter.Tokenizer, error) {
	wire.Build(interpreter.NewTokenizer, interpreter.NewLexemes, interpreter.NewErrors, interpreter.NewIgnoreList)
	return interpreter.Tokenizer{}, nil
}
