package arithmetic

// operator designates an object that combines one or two operands to another.
type operator interface {

	// precedence or priority of the operator.
	precedence() uint8

	// solve is the function the operator uses to transform the operand. The stack
	// contains the operands to use. solve returns an transformed operand and an error.
	solve(*stack) (interface{}, error)
}

// precedence defines the priority between operations.
// For this implementation, a higher precedence means a higher priority.
// https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages
const (
	precedenceOr uint8 = iota
	precedenceAnd
	precedenceEqual
	precedenceLowerGreater
	precedenceMinusPlus
	precedenceDivideMultiply
	precedenceExponant
	precedenceUnary
)
