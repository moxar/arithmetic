package arithmetic

type leftParenthesis struct{}

func (o leftParenthesis) String() string {
	return "("
}

type rightParenthesis struct{}

func (o rightParenthesis) String() string {
	return ")"
}

//
// type Comma struct{}
//
// func (o Comma) String() string {
// 	return ","
// }
