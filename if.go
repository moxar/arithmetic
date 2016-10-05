package arithmetic

// import (
// 	"errors"
// 	"fmt"
// )
//
// func init() {
// 	RegisterFunction("if", If)
// }
//
// func If(st *OperandStack) (Operand, error) {
//
// 	_, err := st.PopInt()
// 	if err != nil {
// 		return nil, fmt.Errorf("if error: undefined argument len: %s", err)
// 	}
//
// 	fail, ok := st.Pop()
// 	if !ok {
// 		return nil, errors.New("if error: third argument must be operand, none found")
// 	}
//
// 	success, ok := st.Pop()
// 	if !ok {
// 		return nil, errors.New("if error: second argument must be operand, none found")
// 	}
//
// 	cond, err := st.PopBool()
// 	if err != nil {
// 		return nil, fmt.Errorf("if error: first argument must be bool: %s", err)
// 	}
//
// 	if cond {
// 		return success, nil
// 	}
//
// 	return fail, nil
// }
