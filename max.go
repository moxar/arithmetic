package arithmetic

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	RegisterFunction("max", Max)
}

func Max(args ...interface{}) (interface{}, error) {

	var def bool
	var max float64
	for _, a := range args {
		o, ok := a.(float64)
		if !ok {
			return nil, fmt.Errorf("max error: argument must be float, having %v(%T)", a)
		}

		if !def {
			max = o
			def = true
		}

		max = math.Max(max, o)
	}

	if !def {
		return nil, errors.New("max error: no argument provided")
	}

	return max, nil
}
