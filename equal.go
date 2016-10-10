package arithmetic

// equal (==) operator.
type equal struct{}

func (o equal) String() string {
	return "=="
}

func (o equal) precedence() uint8 {
	return precedenceEqual
}

func (o equal) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in the proper type (float, bool) and
	// test them.
	return eq(left, right), nil
}

// different (!=) operator.
type different struct{}

func (o different) String() string {
	return "!="
}

func (o different) precedence() uint8 {
	return precedenceEqual
}

func (o different) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms.
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	// cast the left and right terms in the proper type (float, bool) and
	// test them.
	return !eq(left, right), nil
}
