package arithmetic

import "fmt"

type UnaryPlus struct{}

func (o UnaryPlus) String() string {
	return "+"
}

func (o UnaryPlus) Value() (Operand, Operator) {
	return nil, o
}

func (o UnaryPlus) Kind() Kind {
	return KindOperation
}

func (o UnaryPlus) Precedence() uint8 {
	return 4
}

func (o UnaryPlus) Solve(st *OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"+\" must be followed by a valid operand or expression")
	}

	r, err := ToFloat(right)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	return Number(r), nil
}
