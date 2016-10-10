package arithmetic

// greater (>) operator.
type greater struct{}

func (o greater) String() string {
	return ">"
}

func (o greater) precedence() uint8 {
	return precedenceLowerGreater
}

func (o greater) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in floats and
	// test them.
	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	return b, nil
}

type greaterEqual struct{}

// greater or equal (>=) operator.
func (o greaterEqual) String() string {
	return ">="
}

func (o greaterEqual) precedence() uint8 {
	return precedenceLowerGreater
}

func (o greaterEqual) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in floats and
	// test them.
	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}
	if b {
		return true, nil
	}

	return eq(left, right), nil
}

// lower (<) operator.
type lower struct{}

func (o lower) String() string {
	return "<"
}

func (o lower) precedence() uint8 {
	return precedenceLowerGreater
}

func (o lower) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in floats and
	// test them.
	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	if b {
		return false, nil
	}

	return !eq(left, right), nil
}

// lower or equal (<=) operator.
type lowerEqual struct{}

func (o lowerEqual) String() string {
	return "<="
}

func (o lowerEqual) precedence() uint8 {
	return precedenceLowerGreater
}

func (o lowerEqual) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in floats and
	// test them.
	b, ok := gt(left, right)
	if !ok {
		return nil, invalidExpressionError(o, left, right)
	}

	return !b, nil
}
