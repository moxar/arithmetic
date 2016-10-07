package arithmetic

import (
	"fmt"
	"strings"
)

type Function func(...interface{}) (interface{}, error)

type function struct {
	label string
	solve Function
}

func (f function) String() string {
	return f.label
}

var functions = map[string]function{}

func RegisterFunction(label string, f Function) {

	if f == nil {
		panic(fmt.Sprintf("function %s is nil", label))
	}

	label = strings.ToLower(label)

	mustBeUnique(label)

	functions[label] = function{label, f}
}
