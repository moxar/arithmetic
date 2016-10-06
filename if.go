package arithmetic

import (
	"fmt"
)

func init() {
	RegisterFunction("if", if_)
}

func if_(args ...interface{}) (interface{}, error) {

	if len(args) != 3 {
		return nil, fmt.Errorf("if requires 3 arguments, %d provided", len(args))
	}

	cond, ok := toBool(args[0])
	if !ok {
		return nil, fmt.Errorf("invalid expression: if(%v, %v, %v)", args[0], args[1], args[2])
	}

	success := args[1]
	fail := args[2]

	if cond {
		return success, nil
	}

	return fail, nil
}
