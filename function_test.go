package arithmetic

import (
	"fmt"
)

func ExampleRegisterFunction() {

	// Register a new function, that increments the value by one.
	RegisterFunction("increment", func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("increment requires one argument, %d provided", len(args))
		}

		f, ok := ToFloat(args[0])
		if !ok {
			return nil, fmt.Errorf("increment requires integer argument, %v (%T) provided", args[0], args[0])
		}

		return int(f) + 1, nil
	})

	v, err := Parse("increment(2)")
	if err != nil {
		// ...
	}

	fmt.Println(v)
	// Output: 3

}
