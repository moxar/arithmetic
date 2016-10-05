package arithmetic

import "fmt"

type Minus struct{}

func (o Minus) String() string {
	return "-"
}

func (o Minus) Precedence() uint8 {
	return 1
}

func (o Minus) Solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, fmt.Errorf("invalid operation: \"-\" must be followed by a valid operand or expression")
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, fmt.Errorf("invalid operation: \"- %s\" must be preceeded by a valid operand or expression", right)
	}

	return left - right, nil
}
