package arithmetic

import(
	"fmt"
)

func ShuntingYard(input []Token) ([]Token, error) {
	
	st := &OperatorStack{}
	var output []Token
	
	for i, t := range input {
		_, op := t.Value()
		
		if op == nil {
			output = append(output, t)
			continue
		}
		
		if op.Kind() == KindFunction {
			st.Push(op)
			continue
		}
		
		if op.Kind() == KindLeftParenthesis {
			st.Push(op)
			continue
		}
		
		if op.Kind() == KindOperation {
			for v, ok := st.Pop(); ok; {
				
				if v.Precedence() <= op.Precedence() {
					output = append(output, v.(Token))
				}
			}
			
			output = append(output, op.(Token))
			continue
		}
		
		if op.Kind() == KindComma {
			for {
				v, ok := st.Pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression at position %d: %s...", i+1, tokensToString(input[:i+1]))
				}
				
				if v.Kind() == KindLeftParenthesis {
					break
				}
				
				output = append(output, v.(Token))
			}
			continue
		}
		
		if op.Kind() == KindRightParenthesis {
			for {
				v, ok := st.Pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression at position %d: %s...", i+1, tokensToString(input[:i+1]))
				}
				
				// NOTE: What if there is a function ?
				if v.Kind() == KindLeftParenthesis {
					break
				}
				
				output = append(output, v.(Token))
				
			}
			continue
		}
	}
	
	for v, ok := st.Pop(); ok; {
		if v.Kind() == KindLeftParenthesis {
			return nil, fmt.Errorf("mismatched parenthesis: %s", tokensToString(input))
		}
		output = append(output, v.(Token))
	}
	
	return output, nil
}
