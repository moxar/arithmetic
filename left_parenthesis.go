package arithmetic

type LeftParenthesis struct{}

func (o LeftParenthesis) String() string {
	return "("
}

func (o LeftParenthesis) Value() (Operand, Operator) {
	return nil, o
}

func (o LeftParenthesis) Kind() Kind {
	return KindLeftParenthesis
}

func (o LeftParenthesis) Precedence() uint8 {
	return 0
}

func (o LeftParenthesis) Solve(st *OperandStack) (Operand, error) {
	return nil, nil
}
