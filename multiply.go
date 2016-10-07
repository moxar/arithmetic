package arithmetic

type multiply struct{}

func (o multiply) String() string {
	return "*"
}

func (o multiply) precedence() uint8 {
	return precedenceDivide
}

func (o multiply) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	return left * right, nil
}
