package arithmetic

import (
	"strconv"
)

type Number float64

func (o Number) String() string {
	return strconv.FormatFloat(float64(o), 'f', 2, 64)
}

func (o Number) Value() (Operand, Operator) {
	return o, nil
}
