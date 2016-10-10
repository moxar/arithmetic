package arithmetic

import (
	"errors"
)

// divide (/) operator.
type divide struct{}

func (o divide) String() string {
	return "/"
}

func (o divide) precedence() uint8 {
	return precedenceDivideMultiply
}

func (o divide) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The term can be of type "float" or "variable".
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	// Secure division by 0.
	if right == 0 {
		return nil, errors.New("division by 0")
	}

	return left / right, nil
}

// multiply (*) operator.
type multiply struct{}

func (o multiply) String() string {
	return "*"
}

func (o multiply) precedence() uint8 {
	return precedenceDivideMultiply
}

func (o multiply) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The term can be of type "float" or "variable".
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}

	return left * right, nil
}

// modulo (%) operator
type modulo struct{}

func (o modulo) String() string {
	return "%"
}

func (o modulo) precedence() uint8 {
	return precedenceDivideMultiply
}

func (o modulo) solve(st *stack) (interface{}, error) {

	// Retreive right and left terms from stack.
	// The term can be of type "float" or "variable".
	// The decimal part of the inputs must be 0.
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
