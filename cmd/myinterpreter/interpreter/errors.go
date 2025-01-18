package interpreter

type Errors struct {
	unexpectedChar     string
	unterminatedString string
}

func NewErrors() Errors {
	return Errors{unexpectedChar: "Unexpected character", unterminatedString: "Unterminated string"}
}
