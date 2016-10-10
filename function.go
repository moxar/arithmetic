package arithmetic

import (
	"fmt"
	"strings"
)

// Function is a mathematical function that takes a (defined or not) number of arguments
// and returns a value.
//
// Note that the incomming argument can be of various types: bool, float, string.
// Use the ToFloat and ToBool helper if the function requires to used boolean or floats. Otherwise,
// variables (such as e or pi) would not be handled by the function.
type Function func(...interface{}) (interface{}, error)

type function struct {
	label string
	solve Function
}

func (f function) String() string {
	return f.label
}

var functions = map[string]function{}

// RegisterFunction associates a function to a label.
// If the label is already defined, RegisterFunction panics.
// If the function is nil, RegisterFunction also panics.
// The label is not case sensitive.
func RegisterFunction(label string, f Function) {

	// Ensure the function is not nil.
	if f == nil {
		panic(fmt.Sprintf("function %s is nil", label))
	}

	// Make case insensitive.
	label = strings.ToLower(label)

	// Ensure the label is not already registered.
	mustBeUnique(label)

	functions[label] = function{label, f}
}
