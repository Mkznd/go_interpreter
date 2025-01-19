// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/interpreter"
)

// Injectors from wire.go:

func InitializeTokenizer() (interpreter.Tokenizer, error) {
	errors := interpreter.NewErrors()
	v := interpreter.NewIgnoreList()
	lexemes := interpreter.NewLexemes(errors, v)
	parser := interpreter.NewTokenizer(lexemes, errors, v)
	return parser, nil
}
