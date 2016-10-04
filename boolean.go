package arithmetic

import (
	"fmt"
	"strconv"
)

type Boolean bool

func (o Boolean) String() string {
	return strconv.FormatBool(bool(o))
}

func (o Boolean) Value() (Operand, Operator) {
	return o, nil
}

func ToBool(o Operand) (bool, error) {
	v, ok := o.(Boolean)
	if !ok {
		return false, fmt.Errorf("expecing bool, having %v(%T)", o, o)
	}

	return bool(v), nil
}
