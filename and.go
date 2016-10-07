package arithmetic

type and struct{}

func (o and) String() string {
	return "&&"
}

func (o and) precedence() uint8 {
	return precedenceAnd
}

func (o and) solve(st *stack) (interface{}, error) {
	right, err := st.popBool()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popBool()
	if err != nil {
		return nil, leftError(o, right)
	}

	return right && left, nil
}
