package arithmetic

type plus struct{}

func (o plus) String() string {
	return "+"
}

func (o plus) precedence() uint8 {
	return 1
}

func (o plus) solve(st *stack) (interface{}, error) {
	o2, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	o1, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, o2)
	}

	return o1 + o2, nil
}

type unaryPlus struct{}

func (o unaryPlus) String() string {
	return "+"
}

func (o unaryPlus) precedence() uint8 {
	return 4
}

func (o unaryPlus) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	return right, nil
}
