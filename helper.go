package arithmetic

import (
	"fmt"
)

func leftError(o fmt.Stringer, v interface{}) error {
	return fmt.Errorf("invalid operation: \"%s %v\" must be preceeded by a valid operand or expression", o, v)
}

func rightError(o fmt.Stringer) error {
	return fmt.Errorf("invalid operation: \"%s\" must be followed by a valid operand or expression", o)
}

// func equals(o1, o2 Operand) (bool, bool) {
//
// 	// Compare floats.
// 	left, errLeft := ToFloat(o1)
// 	right, errRight := ToFloat(o2)
// 	if errLeft == nil {
// 		if errRight == nil {
// 			return left == right, true
// 		}
// 	}
// 	return false, false
// }
//
// func greater(o1, o2 Operand) (bool, bool) {
//
// 	// Compare floats.
// 	left, errLeft := ToFloat(o1)
// 	right, errRight := ToFloat(o2)
// 	if errLeft == nil {
// 		if errRight == nil {
// 			return left > right, true
// 		}
// 	}
// 	return false, false
// }
