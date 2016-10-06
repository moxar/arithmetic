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

func eq(o1, o2 interface{}) bool {

	f1, ok1 := toFloat(o1)
	f2, ok2 := toFloat(o2)
	if ok1 && ok2 {
		return f1 == f2
	}
	
	return o1 == o2
}

func gt(o1, o2 interface{}) (bool, bool) {

	f1, ok := toFloat(o1)
	if !ok {
		return false, false
	}
	
	f2, ok := toFloat(o2)
	if !ok {
		return false, false
	}
	
	return f1 > f2, true
}

func floatToInt(o float64) (int, bool) {
	i := int(o)
	if float64(i) == o {
		return i, true
	}
	return 0, false
}

func toFloat(val interface{}) (float64, bool) {
	switch t := val.(type) {
		
	case float64:
		return t, true
		
	case variable:
		v, ok := t.value.(float64)
		if !ok {
			return 0, false
		}
		return v, true
		
	default:
		return 0, false
	}
}

func toBool(val interface{}) (bool, bool) {
	switch t := val.(type) {
		
	case bool:
		return t, true
		
	case variable:
		v, ok := t.value.(bool)
		if !ok {
			return false, false
		}
		return v, true
		
	default:
		return false, false
	}
}
