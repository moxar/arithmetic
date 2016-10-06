package arithmetic

import (
	"errors"
)

func solve(input []interface{}) (interface{}, error) {
	st := &stack{}

	for _, t := range input {

		switch v := t.(type) {
		case function:
			size, err := st.popInt()
			if err != nil {
				return nil, err
			}

			args, err := st.slice(size)
			if err != nil {
				return nil, err
			}

			o, err := v.solve(args...)
			if err != nil {
				return nil, err
			}

			st.push(o)

		case operator:
			o, err := v.solve(st)
			if err != nil {
				return nil, err
			}
			st.push(o)

		default:
			st.push(v)
		}
	}

	out, _ := st.pop()

	if _, ok := st.pop(); ok {
		return nil, errors.New("operator stack should be empty")
	}

	return out, nil
}
