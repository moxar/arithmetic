package arithmetic

import (
	"fmt"
)

// shunting yard algorithm: transforms an input from infix to postfix notation.
func shuntingYard(input []interface{}) ([]interface{}, error) {

	os := &stack{}
	as := &stack{}
	var output []interface{}

	// Read each token and look at the type.
	for _, token := range input {

		switch v := token.(type) {

		// If it is a function, push it to the operator stack and increment the arity stack.
		case function:
			as.push(1)
			os.push(v)

		// If it is an operator, pop each operator and function with
		// until an operator with a lower precedence is encountered.
		// Put those values to the output.
		// If a function is popped, pop the arity stack to the output before the operator.
		// Then, push the operator to the stack.
		case operator:
			op := v
			for {
				v, ok := os.pop()
				if !ok {
					break
				}

				if v, ok := v.(operator); ok {
					if v.precedence() >= op.precedence() {
						output = append(output, v)
						continue
					}
				}

				if v, ok := v.(function); ok {
					p, _ := as.pop()
					output = append(output, p)
					output = append(output, v)
					continue
				}

				os.push(v)
				break

			}
			os.push(v)

		// If a left parenthesis is encoutered, push it to the operator stack.
		case leftParenthesis:
			os.push(v)

		// If a right parenthesis is encountered, pop the operator stack until a left parenthesis
		// is encountered. If the parenthesis belongs to a function, pop the function too.
		// Then, push the operator to the stack.
		case rightParenthesis:
			for {
				v, ok := os.pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression: %v", input)
				}

				if _, ok := v.(leftParenthesis); ok {
					v, ok := os.pop()
					if !ok {
						break
					}

					if v, ok := v.(function); ok {
						p, _ := as.pop()
						output = append(output, p, v)
						break
					}

					os.push(v)
					break
				}

				output = append(output, v)
			}

		// If a comma is encountered, pop the operator stack until a left panrethesis is encountered. If the parenthesis belongs to a function, pop the function too.
		case comma:
			as.inc()
			for {
				v, ok := os.pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression: %v", input)
				}

				if v, ok := v.(leftParenthesis); ok {
					os.push(v)
					break
				}

				if _, ok := v.(function); ok {
					p, _ := as.pop()
					output = append(output, p)
				}
				output = append(output, v)
			}

		// If an operand is encountered, append it to the output.
		default:
			output = append(output, v)
		}
	}

	// Once the input has been read, pop the operator stack. When a function is encountered, pop
	// the arity stack.
	for {

		v, ok := os.pop()
		if !ok {
			break
		}

		if _, ok := v.(leftParenthesis); ok {
			return nil, fmt.Errorf("mismatched parenthesis: %v", input)
		}

		if _, ok := v.(function); ok {
			p, _ := as.pop()
			output = append(output, p)
		}
		output = append(output, v)
	}

	return output, nil
}
