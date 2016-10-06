package arithmetic

type not struct{}

func (o not) String() string {
	return "!"
}

func (o not) precedence() uint8 {
	return 4
}

func (o not) solve(st *stack) (interface{}, error) {
	right, err := st.popBool()
	if err != nil {
		return nil, rightError(o)
	}

	return !right, nil
}
