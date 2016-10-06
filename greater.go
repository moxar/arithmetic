package arithmetic

type greater struct{}

func (o greater) String() string {
	return ">"
}

func (o greater) precedence() uint8 {
	return 0
}

func (o greater) solve(st *stack) (interface{}, error) {
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

	return b, nil
}

type greaterEqual struct{}

func (o greaterEqual) String() string {
	return ">="
}

func (o greaterEqual) precedence() uint8 {
	return 0
}

func (o greaterEqual) solve(st *stack) (interface{}, error) {
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
		return true, nil
	}

	return eq(left, right), nil
}
