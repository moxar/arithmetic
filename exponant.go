package arithmetic

import (
	"math"
)

// exponant (^) operator.
type exponant struct{}

func (o exponant) String() string {
	return "^"
}

func (o exponant) precedence() uint8 {
	return precedenceExponant
}

func (o exponant) solve(st *stack) (interface{}, error) {

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

	return math.Pow(left, right), nil
}
