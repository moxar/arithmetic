package arithmetic

// import "fmt"
//
// type UnaryMinus struct{}
//
// func (o UnaryMinus) String() string {
// 	return "-"
// }
//
// func (o UnaryMinus) Precedence() uint8 {
// 	return 4
// }
//
// func (o UnaryMinus) Solve(st *OperandStack) (Operand, error) {
// 	right, ok := st.Pop()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \"-\" must be followed by a valid operand or expression")
// 	}
//
// 	r, err := ToFloat(right)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid operand: %s", err)
// 	}
//
// 	return Number(-r), nil
// }
