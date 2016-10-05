package arithmetic

import "fmt"

type minus struct{}

func (o minus) String() string {
	return "-"
}

func (o minus) precedence() uint8 {
	return 1
}

func (o minus) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, fmt.Errorf("invalid operation: \"-\" must be followed by a valid operand or expression")
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, fmt.Errorf("invalid operation: \"- %v\" must be preceeded by a valid operand or expression", right)
	}

	return left - right, nil
}
