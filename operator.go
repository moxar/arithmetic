package arithmetic

type operator interface {
	precedence() uint8
	solve(*stack) (interface{}, error)
}

// precedence defines the priority between operations.
// For this implementation, a higher precedence means a higher priority.
// https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages
const (
	precedenceOr uint8 = iota
	precedenceAnd
	precedenceEqual
	precedenceGreater
	precedencePlus
	precedenceDivide
	precedenceExponant
	precedenceUnary
)
