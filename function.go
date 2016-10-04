package arithmetic

import (
	"fmt"
)

var functions = map[string]Token{}

func RegisterFunction(label string, f SolveFunc) {

	if f == nil {
		panic(fmt.Sprintf("function %s is nil", label))
	}

	_, ok := functions[label]
	if ok {
		panic(fmt.Sprintf("function %s defined twice", label))
	}

	var o function
	o.label = label
	o.solveFunc = f

	functions[label] = o
}

type SolveFunc func(*OperandStack) (Operand, error)

type function struct {
	label     string
	solveFunc SolveFunc
}

func (f function) String() string {
	return f.label
}

func (f function) Value() (Operand, Operator) {
	return nil, f
}

func (f function) Kind() Kind {
	return KindFunction
}

func (f function) Precedence() uint8 {
	return 5
}

func (f function) Solve(o *OperandStack) (Operand, error) {
	return f.solveFunc(o)
}
