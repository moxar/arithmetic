package arithmetic

type Operator interface{
	Precedence() uint8
	Solve(OperandStack) Operand
}
