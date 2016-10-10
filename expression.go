package arithmetic

// Expression is a function that takes a string, analyses it and returns a value if the expression
// is recognized. If The second (bool) argument is true if the expression is recognized, false otherwise.
//
// When the parser detects an input, each registered expression is tested.
// The function should return false matches very quickly to improve performances.
type Expression func(string) (interface{}, bool)

var expressions []Expression

// RegisterExpression saves an expression matcher.
func RegisterExpression(e Expression) {

	if e == nil {
		panic("provided expression is nil")
	}

	expressions = append(expressions, e)
}
