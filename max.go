package arithmetic

import (
	"errors"
	"fmt"
	"math"
)

func Max(st *OperandStack) (Operand, error) {
	s, ok := st.Pop()
	if !ok {
		return nil, errors.New("stack error: first element must be argument number")
	}

	var m float64
	var def bool
	
	size, ok := s.(Number)
	if !ok {
		return nil, errors.New("stack error: first element must be integer")
	}
	for i := 0; i < int(size); i++ {
		op, ok := st.Pop()
		if !ok {
			return nil, fmt.Errorf("\"max() \" %d arguments, having %d:", int(size), i)
		}
		
		num, ok := op.(Number)
		if !ok {
			return nil, fmt.Errorf("\"max() \" expects arguments, numeric arguments, having %s:", op)
		}
		
		if !def {
			m = float64(num)
			def = true
		}
		m = math.Max(m, float64(num))
	}
	
	if !def {
		return nil, errors.New("invalid usage of \"max()\": at least one argument is required")
	}
	
	return Number(m), nil
}
