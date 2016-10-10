package arithmetic

import (
	"errors"
	"fmt"
)

func init() {
	RegisterFunction("min", min)
}

// min is a function that returns the lowest of the provided inputs.
func min(args ...interface{}) (interface{}, error) {

	if len(args) == 0 {
		return nil, errors.New("min requires at least one argument")
	}

	// Ensure each argument is a float, or a "variable" float.
	var m, f float64
	var def bool
	for _, a := range args {
		switch t := a.(type) {
		case float64:
			f = t
		case variable:
			v, ok := t.value.(float64)
			if !ok {
				return nil, fmt.Errorf("min requires numeric arguments, %s given", t)
			}
			f = v
		default:
			return nil, fmt.Errorf("min requires numeric arguments, %v given", a)
		}

		if f < m || !def {
			def = true
			m = f
		}
	}

	return m, nil
}
