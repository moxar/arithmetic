package arithmetic

type equal struct{}

func (o equal) String() string {
	return "=="
}

func (o equal) precedence() uint8 {
	return 0
}

func (o equal) solve(st *stack) (interface{}, error) {
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	b, ok := eq(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	return b, nil
}
