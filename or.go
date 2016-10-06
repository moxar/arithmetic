package arithmetic

type or struct{}

func (o or) String() string {
	return "||"
}

func (o or) precedence() uint8 {
	return 3
}

func (o or) solve(st *stack) (interface{}, error) {
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
