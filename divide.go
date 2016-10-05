package arithmetic

// import "fmt"
//
// type Divide struct{}
//
// func (o Divide) String() string {
// 	return "/"
// }
//
// func (o Divide) Precedence() uint8 {
// 	return 2
// }
//
// func (o Divide) Solve(st *stack) (interface{}, error) {
// 	right, ok := st.popFloat()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \"/\" must be followed by a valid operand or expression")
// 	}
//
// 	left, ok := st.popFloat()
// 	if !ok {
// 		return nil, fmt.Errorf("invalid operation: \"/ %s\" must be preceeded by a valid operand or expression", right)
// 	}
//
// 	return l / r, nil
// }
