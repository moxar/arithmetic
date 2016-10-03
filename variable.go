package arithmetic

import (
	"fmt"
)

func init() {
	RegisterVariable("pi", 3.142)
	RegisterVariable("e", 2.718)
}

var variables = map[string]variable{}

func RegisterVariable(label string, value interface{}) {

	_, ok := variables[label]
	if ok {
		panic(fmt.Sprintf("variable %s defined twice", label))
	}

	var v variable
	v.label = label

	switch t := value.(type) {
	case float64:
		v.operand = Number(t)
	default:
		panic(fmt.Sprintf("forbidden variable %s type %T", label, value))
	}

	variables[label] = v
}

type variable struct {
	label   string
	operand Operand
}

func (v variable) String() string {
	return v.label
}

func (v variable) Value() (Operand, Operator) {
	return v.operand, nil
}
