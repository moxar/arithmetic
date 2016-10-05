package arithmetic

type leftParenthesis struct{}

func (o leftParenthesis) String() string {
	return "("
}

type rightParenthesis struct{}

func (o rightParenthesis) String() string {
	return ")"
}

type comma struct{}

func (o comma) String() string {
	return ","
}
