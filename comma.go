package arithmetic

type Comma struct{}

func (o Comma) String() string {
	return ","
}

func (o Comma) Value() (Operand, Operator) {
	return nil, o
}

func (o Comma) Kind() Kind {
	return KindComma
}

func (o Comma) Precedence() uint8 {
	return 0
}

func (o Comma) Solve(st OperandStack) (Operand, error) {
	return nil, nil
}
