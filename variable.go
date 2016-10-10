package arithmetic

import (
	"math"
	"strings"
)

func init() {
	RegisterVariable("e", math.E)
	RegisterVariable("pi", math.Pi)
	RegisterVariable("phi", math.Phi)
	RegisterVariable("sqrt2", math.Sqrt2)
	RegisterVariable("sqrte", math.SqrtE)
	RegisterVariable("sqrtpi", math.SqrtPi)
	RegisterVariable("sqrtphi", math.SqrtPhi)
	RegisterVariable("ln2", math.Ln2)
	RegisterVariable("log2e", math.Log2E)
	RegisterVariable("ln10", math.Ln10)
	RegisterVariable("ln10e", math.Log10E)
	RegisterVariable("true", true)
	RegisterVariable("false", false)
}

var variables = map[string]interface{}{}

// RegisterVariable associates a variable to a label.
// If the label is already defined, RegisterVariable panics.
// The label is not case sensitive.
func RegisterVariable(label string, value interface{}) {

	label = strings.ToLower(label)

	mustBeUnique(label)

	if v, ok := value.(int); ok {
		variables[label] = variable{label, float64(v)}
		return
	}
	variables[label] = variable{label, value}

}

// variable is a structure that holds a variable value and its label.
type variable struct {
	label string
	value interface{}
}

func (v variable) String() string {
	return v.label
}
