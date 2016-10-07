package arithmetic

type lower struct{}

func (o lower) String() string {
	return "<"
}

func (o lower) precedence() uint8 {
	return precedenceGreater
}

func (o lower) solve(st *stack) (interface{}, error) {
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	if b {
		return false, nil
	}

	return !eq(left, right), nil
}

type lowerEqual struct{}

func (o lowerEqual) String() string {
	return "<="
}

func (o lowerEqual) precedence() uint8 {
	return precedenceGreater
}

func (o lowerEqual) solve(st *stack) (interface{}, error) {
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	return !b, nil
}
