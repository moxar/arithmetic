package arithmetic

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	RegisterFunction("max", max)
}

func max(args ...interface{}) (interface{}, error) {

	var def bool
	var m float64
	for _, a := range args {
		o, ok := a.(float64)
		if !ok {
			return nil, fmt.Errorf("max error: argument must be float, having %v(%T)", a)
		}

		if !def {
			m = o
			def = true
		}

		m = math.Max(m, o)
	}

	if !def {
		return nil, errors.New("max error: no argument provided")
	}

	return m, nil
}
