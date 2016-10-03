package arithmetic

type Token interface{
	Value() (Operand, Operator)
}
