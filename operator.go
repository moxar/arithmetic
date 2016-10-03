package arithmetic

type Operator interface {
	Kind() Kind
	Precedence() uint8
	Solve(OperandStack) (Operand, error)
}
