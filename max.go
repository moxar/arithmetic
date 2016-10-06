package arithmetic

import (
	"errors"
	"fmt"
)

func init() {
	RegisterFunction("max", max)
}

func max(args ...interface{}) (interface{}, error) {

	var m float64
	var def bool

	var f float64
	for _, a := range args {
		switch t := a.(type) {
		case float64:
			f = t
		case variable:
			v, ok := t.value.(float64)
			if !ok {
				return nil, fmt.Errorf("max requires numeric arguments, %s given", t)
			}
			f = v
		default:
			return nil, fmt.Errorf("max requires numeric arguments, %v given", a)
		}

		if f > m || !def {
			def = true
			m = f
		}
	}

	if !def {
		return nil, errors.New("max requires at least one argument")
	}

	return m, nil
}
