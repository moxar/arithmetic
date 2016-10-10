package arithmetic

// unary not (!) operator
type not struct{}

func (o not) String() string {
	return "!"
}

func (o not) precedence() uint8 {
	return precedenceUnary
}

func (o not) solve(st *stack) (interface{}, error) {

	// Retreive right term from stack.
	// The term can be of type "bool" or "variable".
	right, err := st.popBool()
	if err != nil {
		return nil, rightError(o)
	}

	return !right, nil
}
