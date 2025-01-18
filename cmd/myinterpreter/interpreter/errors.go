package interpreter

type Errors struct {
	unexpectedChar string
}

func NewErrors() Errors {
	return Errors{unexpectedChar: "Unexpected character"}
}
