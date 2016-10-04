package arithmetic

import "fmt"

type Different struct{}

func (o Different) String() string {
	return "!="
}

func (o Different) Value() (Operand, Operator) {
	return nil, o
}

func (o Different) Kind() Kind {
	return KindOperation
}

func (o Different) Precedence() uint8 {
	return 0
}

func (o Different) Solve(st *OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"!=\" must be followed by a valid operand or expression")
	}

	left, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"!= %s\" must be preceeded by a valid operand or expression", right)
	}

	b, ok := equals(left, right)
	if !ok {
		return nil, fmt.Errorf("invalid expression %s != %s", left, right)
	}

	return Boolean(!b), nil
}
