package arithmetic

// import "fmt"
//
// type GreaterEqual struct{}
//
// func (o GreaterEqual) String() string {
// 	return ">="
// }
//
// func (o GreaterEqual) Precedence() uint8 {
// 	return 0
// }
//
// func (o GreaterEqual) Solve(st *OperandStack) (Operand, error) {
// 	right, ok := st.Pop()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \">=\" must be followed by a valid operand or expression")
// 	}
//
// 	left, ok := st.Pop()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \">= %s\" must be preceeded by a valid operand or expression", right)
// 	}
//
// 	b, ok := greater(left, right)
// 	if !ok {
// 		return nil, fmt.Errorf("invalid expression %s >= %s", left, right)
// 	}
// 	if b {
// 		return Boolean(true), nil
// 	}
//
// 	b, _ = equals(left, right)
// 	return Boolean(b), nil
// }
