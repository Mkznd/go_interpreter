package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/di"
	"os"
	"slices"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	allowedCommands := []string{
		"tokenize", "parse",
	}

	if !slices.Contains(allowedCommands, command) {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	interpreter, err := di.InitializeInterpreter()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing the interpreter: %v\n", err)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	if len(fileContents) > 0 {
		switch command {
		case "tokenize":
			tokens, code := interpreter.Tokenize(fileContents)
			interpreter.DisplayTokens(tokens, os.Stdout)
			os.Exit(code)
		case "parse":
			tokens, code := interpreter.Tokenize(fileContents)
			expressions := interpreter.Parse(tokens)
			fmt.Println(expressions)
			os.Exit(code)
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
