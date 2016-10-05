package arithmetic

func ShuntingYard(input []interface{}) ([]interface{}, error) {

	os := &stack{}
	// 	as := &stack{}
	var output []interface{}

	for _, token := range input {

		switch v := token.(type) {

		// 		case Function:
		// 			as.push(1)
		// 			os.push(v)

		case operator:
			for {
				v, ok := os.pop()
				if !ok {
					break
				}

				// 				if v, ok := v.(function); ok {
				// 					output = append(output, as.pop())
				// 					output = append(output, v)
				// 					continue
				// 				}

				if _, ok := v.(operator); ok {
					output = append(output, v)
					continue
				}

				os.push(v)
				break
			}
			os.push(v)

			// 		case LeftParenthesis:
			// 			os.push(v)

			// 		case RightParethesis:
			// 			for {
			// 				v, ok := os.pop()
			// 				if !ok {
			// 					return nil, fmt.Errorf("invalid expression at position %d: %v...", i+1, input[:i+1])
			// 				}
			//
			// 				if v, ok := v.(LeftParenthesis); ok {
			// 					v, ok := os.pop()
			// 					if !ok {
			// 						break
			// 					}
			//
			// 					if v, ok := v.(function); ok {
			// 						output = append(output, as.pop(), v)
			// 						break
			// 					}
			//
			// 					os.push(v)
			// 					break
			// 				}
			//
			// 				output = append(output, v)
			// 			}

			// 		case Comma:
			// 			as.Inc()
			// 			for {
			// 				v, ok := os.pop()
			// 				if !ok {
			// 					return nil, fmt.Errorf("invalid expression at position %d: %v...", i+1, input[:i+1])
			// 				}
			//
			// 				if v, ok := v.(LeftParenthesis); ok {
			// 					os.push(v)
			// 					break
			// 				}
			//
			// 				if v, ok := v.(function); ok {
			// 					output = append(output, as.pop())
			// 				}
			// 				output = append(output, v)
			// 			}

		default:
			output = append(output, v)
		}
	}

	for {

		v, ok := os.pop()
		if !ok {
			break
		}

		// 		if v, ok := v.(LeftParenthesis); ok {
		// 			return nil, fmt.Errorf("mismatched parenthesis: %v", input)
		// 		}

		// 		if v, ok := v.(function); ok {
		// 			output = append(output, as.pop())
		// 		}
		output = append(output, v)
	}

	return output, nil
}
