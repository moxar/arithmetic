package arithmetic

type modulo struct{}

func (o modulo) String() string {
	return "%"
}

func (o modulo) precedence() uint8 {
	return 2
}

func (o modulo) solve(st *stack) (interface{}, error) {
	right, err := st.popInt()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popInt()
	if err != nil {
		return nil, leftError(o, right)
	}

	return left % right, nil
}
