package arithmetic

import "fmt"

type Operand interface{}

func ToFloat(o Operand) (float64, error) {
	v, ok := o.(float64)
	if !ok {
		return 0, fmt.Errorf("expecing float, having %v(%T)", o, o)
	}

	return v, nil
}

func ToInt(o Operand) (int, error) {
	v, err := ToFloat(o)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}
