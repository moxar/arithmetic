package arithmetic

// minus (-) operator.
type minus struct{}

func (o minus) String() string {
	return "-"
}

func (o minus) precedence() uint8 {
	return precedenceMinusPlus
}

func (o minus) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The terms can be of type "float" or "variable".
	o2, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	o1, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, o2)
	}

	return o1 - o2, nil
}

// unary minus (-) operator.
type unaryMinus struct{}

func (o unaryMinus) String() string {
	return "-"
}

func (o unaryMinus) precedence() uint8 {
	return precedenceUnary
}

func (o unaryMinus) solve(st *stack) (interface{}, error) {

	// Retreive right term from stack.
	// The term can be of type "float" or "variable".
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	return -right, nil
}

// plus (+) operator.
type plus struct{}

func (o plus) String() string {
	return "+"
}

func (o plus) precedence() uint8 {
	return precedenceMinusPlus
}

func (o plus) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The terms can be of type "float" or "variable".
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

// unary plus (-) operator.
func (o unaryPlus) precedence() uint8 {
	return precedenceUnary
}

func (o unaryPlus) solve(st *stack) (interface{}, error) {

	// Retreive right term from stack.
	// The term can be of type "float" or "variable".
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	return right, nil
}
