package interpreter

import (
	"fmt"
)

type ParenthesesScanner struct {
	lexems map[string]string
}

func NewParenthesesScanner(lexems map[string]string) ParenthesesScanner {
	return ParenthesesScanner{lexems: lexems}
}

func (s ParenthesesScanner) scan(buf []byte) {
	content := string(buf[:])
	for _, letter := range content {
		value, found := s.lexems[string(letter)]
		if !found {
			continue
		}
		fmt.Println(value, string(letter), "null")
	}
	fmt.Println("EOF  null")
}
