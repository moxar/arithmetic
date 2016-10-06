package arithmetic

type modulo struct{}

func (o modulo) String() string {
	return "%"
}

func (o modulo) precedence() uint8 {
	return 2
}

func (o modulo) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	r, ok := floatToInt(right)
	if !ok {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	l, ok := floatToInt(left)
	if !ok {
		return nil, leftError(o, right)
	}

	return float64(l % r), nil
}
