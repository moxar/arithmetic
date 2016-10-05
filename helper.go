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

func invalidExpressionError(o fmt.Stringer, left, right interface{}) error {
	return fmt.Errorf("invalid expression %v %s %v", left, o, right)
}

func eq(o1, o2 interface{}) (bool, bool) {

	// Compare floats.
	f1, ok1 := o1.(float64)
	f2, ok2 := o2.(float64)
	if ok1 != ok2 {
		return false, false
	}
	if ok1 {
		return f1 == f2, true
	}

	// Compare other types...

	return false, false
}

func gt(o1, o2 interface{}) (bool, bool) {

	// Compare floats.
	f1, ok1 := o1.(float64)
	f2, ok2 := o2.(float64)
	if ok1 != ok2 {
		return false, false
	}
	if ok1 {
		return f1 > f2, true
	}

	// Compare other types...

	return false, false
}
