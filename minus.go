package arithmetic

type minus struct{}

func (o minus) String() string {
	return "-"
}

func (o minus) precedence() uint8 {
	return 1
}

func (o minus) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	return left - right, nil
}
