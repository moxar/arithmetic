package arithmetic

import "fmt"

type Minus struct{}

func (o Minus) String() string {
	return "-"
}

func (o Minus) Value() (Operand, Operator) {
	return nil, o
}

func (o Minus) Kind() Kind {
	return KindOperation
}

func (o Minus) Precedence() uint8 {
	return 1
}

func (o Minus) Solve(st *OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"-\" must be followed by a valid operand or expression")
	}

	r, err := ToFloat(right)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	left, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"- %s\" must be preceeded by a valid operand or expression", right)
	}

	l, err := ToFloat(left)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	return Number(l - r), nil
}
