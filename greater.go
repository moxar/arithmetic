package arithmetic

import "fmt"

type Greater struct{}

func (o Greater) String() string {
	return ">"
}

func (o Greater) Value() (Operand, Operator) {
	return nil, o
}

func (o Greater) Kind() Kind {
	return KindOperation
}

func (o Greater) Precedence() uint8 {
	return 0
}

func (o Greater) Solve(st *OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \">\" must be followed by a valid operand or expression")
	}

	left, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"> %s\" must be preceeded by a valid operand or expression", right)
	}

	b, ok := greater(left, right)
	if !ok {
		return nil, fmt.Errorf("invalid expression %s > %s", left, right)
	}

	return Boolean(b), nil
}
