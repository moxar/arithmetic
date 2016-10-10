package arithmetic

import (
	"errors"
)

// solve the postfix input.
func solve(input []interface{}) (interface{}, error) {
	st := &stack{}

	// read the input. For each token...
	for _, t := range input {

		// If it is a function, retreive the functions arguments: pop the operand
		// stack to get the number of arguments. Then, get the last n elements
		// from the stack (and remove them from it), and injects them to the
		// function. Finaly, solve the function, and put the output to the operator stack.
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

		// If it is an operator, solve it and push its output to the operand stack.
		// To solve an operator, inject the stack to the solve method. The solve method
		// will pop the number of operands required by the operator.
		case operator:
			o, err := v.solve(st)
			if err != nil {
				return nil, err
			}
			st.push(o)

		// If it is an operand, push it to the operand stack.
		default:
			st.push(v)
		}
	}

	// The output is the last operand on the stack. If there are more than one, the input
	// postfix is non valid.
	out, _ := st.pop()

	if _, ok := st.pop(); ok {
		return nil, errors.New("operator stack should be empty")
	}

	return out, nil
}
