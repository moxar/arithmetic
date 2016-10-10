package arithmetic

import (
	"fmt"
)

// mustBeUnique ensures a label is not registered in the functions, variables or aliases.
func mustBeUnique(label string) {

	if _, ok := functions[label]; ok {
		panic(fmt.Sprintf("%s already defined as function", label))
	}

	if _, ok := variables[label]; ok {
		panic(fmt.Sprintf("%s already defined as variable", label))
	}

	if _, ok := aliases[label]; ok {
		panic(fmt.Sprintf("%s already defined as alias", label))
	}
}

// leftError returns the error triggered by an invalid left operand.
func leftError(o fmt.Stringer, v interface{}) error {
	return fmt.Errorf("invalid operation: \"%s %v\" must be preceeded by a valid operand or expression", o, v)
}

// leftError returns the error triggered by an invalid right operand.
func rightError(o fmt.Stringer) error {
	return fmt.Errorf("invalid operation: \"%s\" must be followed by a valid operand or expression", o)
}

// invalidExpressionError returns the error triggered by an invalid expression.
func invalidExpressionError(o fmt.Stringer, left, right interface{}) error {
	return fmt.Errorf("invalid expression %v %s %v", left, o, right)
}

// eq checks if o1 and o2 are equals. It convers types.
func eq(o1, o2 interface{}) bool {

	f1, ok1 := ToFloat(o1)
	f2, ok2 := ToFloat(o2)
	if ok1 && ok2 {
		return f1 == f2
	}

	b1, ok1 := ToBool(o1)
	b2, ok1 := ToBool(o2)
	if ok1 && ok2 {
		return b1 == b2
	}

	return o1 == o2
}

// eq checks if o1 is greater than o2. It convers types.
func gt(o1, o2 interface{}) (bool, bool) {

	f1, ok := ToFloat(o1)
	if !ok {
		return false, false
	}

	f2, ok := ToFloat(o2)
	if !ok {
		return false, false
	}

	return f1 > f2, true
}

// floatToInt transforms a float to an int if the decimal part of the float is 0.
func floatToInt(o float64) (int, bool) {
	i := int(o)
	if float64(i) == o {
		return i, true
	}
	return 0, false
}

// ToFloat casts the input into a float. This helper should be used in the custom funcs to
// use both floats and named variables (such as "e")
func ToFloat(val interface{}) (float64, bool) {
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

// ToBool casts the input into a bool. This helper should be used in the custom funcs to
// use both bools and named variables (such as "true")
func ToBool(val interface{}) (bool, bool) {
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
