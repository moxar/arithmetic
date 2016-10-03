package arithmetic

import (
	"fmt"
	"math"
)

type Exponant struct{}

func (o Exponant) String() string {
	return "^"
}

func (o Exponant) Value() (Operand, Operator) {
	return nil, o
}

func (o Exponant) Kind() Kind {
	return KindOperation
}

func (o Exponant) Precedence() uint8 {
	return 3
}

func (o Exponant) Solve(st *OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"*\" must be followed by a valid operand or expression")
	}

	r, err := ToFloat(right)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	left, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"* %s\" must be preceeded by a valid operand or expression", right)
	}

	l, err := ToFloat(left)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	return Number(math.Pow(l,r)), nil
}
