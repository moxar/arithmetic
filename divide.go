package arithmetic

type divide struct{}

func (o divide) String() string {
	return "/"
}

func (o divide) precedence() uint8 {
	return 2
}

func (o divide) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	return left / right, nil
}
