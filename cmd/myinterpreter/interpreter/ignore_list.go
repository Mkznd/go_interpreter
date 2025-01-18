package interpreter

func NewIgnoreList() []string {
	return []string{
		" ", "\r", "\t",
	}
}
