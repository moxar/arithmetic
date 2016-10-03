package arithmetic

import (
	"errors"
)

func Solve(input []Token) (Operand, error) {
	st := &OperandStack{}

	for _, t := range input {
		operand, operator := t.Value()
		if operand != nil {
			st.Push(operand)
			continue
		}

		o, err := operator.Solve(st)
		if err != nil {
			return nil, err
		}

		st.Push(o)
	}

	out, ok := st.Pop()
	if !ok {
		return nil, errors.New("empty postfix input")
	}

	_, ok = st.Pop()
	if ok {
		return nil, errors.New("missing operand in postfix input")
	}

	return out, nil
}
