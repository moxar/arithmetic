package arithmetic

import (
	"fmt"
)

func ToFloat(o Operand) (float64, error) {
	v, ok := o.(Number)
	if !ok {
		return 0, fmt.Errorf("expecing float, having %v(%T)", o, o)
	}

	return float64(v), nil
}

func ToInt(o Operand) (int, error) {
	v, err := ToFloat(o)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func equals(o1, o2 Operand) (bool, bool) {
	
	// Compare floats.
	left, errLeft := ToFloat(o1)
	right, errRight := ToFloat(o2)
	if errLeft == nil {
		if errRight == nil {
			return left == right, true
		}
	}
	return false, false
}
