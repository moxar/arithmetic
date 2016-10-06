package arithmetic

import (
	"fmt"
)

func shuntingYard(input []interface{}) ([]interface{}, error) {

	os := &stack{}
	as := &stack{}
	var output []interface{}

	for _, token := range input {

		switch v := token.(type) {

		case function:
			as.push(1)
			os.push(v)

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

		case leftParenthesis:
			os.push(v)

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

		default:
			output = append(output, v)
		}
	}

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
