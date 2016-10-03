package arithmetic

type RightParenthesis struct{}

func (o RightParenthesis) String() string {
	return ")"
}

func (o RightParenthesis) Value() (Operand, Operator) {
	return nil, o
}

func (o RightParenthesis) Kind() Kind {
	return KindRightParenthesis
}

func (o RightParenthesis) Precedence() uint8 {
	return 0
}

func (o RightParenthesis) Solve(st OperandStack) (Operand, error) {
	return nil, nil
}
