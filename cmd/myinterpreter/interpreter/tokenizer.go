package interpreter

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Tokenizer struct {
	lexemes Lexemes
	errors  Errors
}

func (t Tokenizer) Tokenize(buf []byte) ([]Token, int) {
	code := 0
	start := time.Now()
	var tokens []Token
	lines := strings.Split(string(buf[:]), "\n")
	for i, line := range lines {
		for pos := 0; pos < len(line); pos++ {
			token, newPos, err := t.lexemes.ResolveLexemes(line, pos)
			pos = newPos
			if err != nil {
				if err.Error() == t.errors.unexpectedChar {
					fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", i+1, t.errors.unexpectedChar, token.Lexeme)
					code = 65
					continue
				} else if err.Error() == t.errors.unterminatedString {
					fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", i+1, t.errors.unterminatedString)
					code = 65
					break
				}
			}
			if token.Lexeme == "" {
				break
			}
			if token.Lexeme == "//" {
				break
			}
			tokens = append(tokens, token)
		}
	}
	elapsed := time.Since(start)
	fmt.Fprintln(os.Stderr, "Tokenization took", elapsed.Milliseconds(), "millisecs")
	return tokens, code
}

func NewTokenizer(lexemes Lexemes, errors Errors, ignore []string) Tokenizer {
	return Tokenizer{
		lexemes: lexemes,
		errors:  errors,
	}
}

func (t Tokenizer) Display(tokens []Token, writer io.Writer) {
	for _, token := range tokens {
		_, err := writer.Write([]byte(token.String() + "\n"))
		if err != nil {
			return
		}
	}
	fmt.Println("EOF  null")
}

type Token struct {
	TokenType string
	Lexeme    string
	Literal   string
}

func (token Token) String() string {
	return fmt.Sprintf("%s %s %s", token.TokenType, token.Lexeme, token.Literal)

}

func NewToken(tokenType string, lexeme string, literal string) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
	}
}
