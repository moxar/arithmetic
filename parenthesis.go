package arithmetic

// left parenthesis.
type leftParenthesis struct{}

func (o leftParenthesis) String() string {
	return "("
}

// right parenthesis.
type rightParenthesis struct{}

func (o rightParenthesis) String() string {
	return ")"
}

type comma struct{}

func (o comma) String() string {
	return ","
}
