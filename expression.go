package arithmetic

type Expression func(string) (interface{}, bool)

var expressions []Expression

func RegisterExpression(e Expression) {

	if e == nil {
		panic("provided expression is nil")
	}

	expressions = append(expressions, e)
}
