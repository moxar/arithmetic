package arithmetic

import (
	"strconv"
)

type String string

func (o String) String() string {
	return strconv.Quote(string(o))
}

func (o String) Value() (Operand, Operator) {
	return o, nil
}
