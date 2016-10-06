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

type different struct{}

func (o different) String() string {
	return "!="
}

func (o different) precedence() uint8 {
	return 0
}

func (o different) solve(st *stack) (interface{}, error) {
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

	return !b, nil
}
