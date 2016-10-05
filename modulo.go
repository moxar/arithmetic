package arithmetic

// type Modulo struct{}
//
// func (o Modulo) String() string {
// 	return "%"
// }
//
// func (o Modulo) Precedence() uint8 {
// 	return 2
// }
//
// func (o Modulo) Solve(st *OperandStack) (Operand, error) {
// 	right, ok := st.Pop()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \"%\" must be followed by a valid operand or expression")
// 	}
//
// 	r, err := ToInt(right)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid operand: %s", err)
// 	}
//
// 	left, ok := st.Pop()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \"% %s\" must be preceeded by a valid operand or expression", right)
// 	}
//
// 	l, err := ToInt(left)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid operand: %s", err)
// 	}
//
// 	return Number(l % r), nil
// }
