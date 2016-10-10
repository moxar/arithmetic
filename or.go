package arithmetic

// and (||) operator.
type or struct{}

func (o or) String() string {
	return "||"
}

func (o or) precedence() uint8 {
	return precedenceOr
}

func (o or) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The term can be of type "bool" or "variable".
	right, err := st.popBool()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popBool()
	if err != nil {
		return nil, leftError(o, right)
	}

	return right || left, nil
}
