package arithmetic

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	RegisterFunction("max", Max)
}

func Max(st *OperandStack) (Operand, error) {

	var max float64
	var def bool

	size, err := st.PopInt()
	if err != nil {
		return nil, fmt.Errorf("max error: undefined argument len: %s", err)
	}

	for i := 0; i < size; i++ {
		challenger, err := st.PopFloat()
		if err != nil {
			return nil, fmt.Errorf("max error: argument must be float: %s", err)
		}

		if !def {
			max = challenger
			def = true
		}
		max = math.Max(max, challenger)
	}

	if !def {
		return nil, errors.New("max error: no argument provided")
	}

	return Number(max), nil
}
