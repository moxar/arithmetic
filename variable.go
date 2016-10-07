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

func RegisterVariable(label string, value interface{}) {

	label = strings.ToLower(label)

	mustBeUnique(label)

	variables[label] = variable{label, value}

}

type variable struct {
	label string
	value interface{}
}

func (v variable) String() string {
	return v.label
}
