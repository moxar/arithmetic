package arithmetic

import (
	"fmt"
)

func init() {
	RegisterFunction("if", if_)
}

// if_ is a function that takes 3 arguments: if(cond, success, failure).
// If cond is true, return success. Else, return failure.
//
// NOTE: transform this to a switch with variadic (2n+1) arguments ?
func if_(args ...interface{}) (interface{}, error) {

	// Ensure there are 3 arguments.
	if len(args) != 3 {
		return nil, fmt.Errorf("if requires 3 arguments, %d provided", len(args))
	}

	// Ensure the first arg is a boolean.
	cond, ok := ToBool(args[0])
	if !ok {
		return nil, fmt.Errorf("invalid expression: if(%v, %v, %v)", args[0], args[1], args[2])
	}

	if cond {
		return args[1], nil
	}

	return args[2], nil
}
