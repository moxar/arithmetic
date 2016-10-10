package arithmetic

import (
	"fmt"
)

func ExampleRegisterExpression() {

	// Register an expression that checks if the input is foo:
	// f followed by any number of o.
	RegisterExpression(func(input string) (interface{}, bool) {
		if len(input) == 0 {
			return nil, false
		}

		// check the first rune is an f.
		if input[0] != 'f' {
			return nil, false
		}

		// Check that every other rune is an o.
		for i, r := range input {
			if i == 0 {
				continue
			}

			if r != 'o' {
				return nil, false
			}
		}

		return true, true
	})

	v, err := Parse("fooOOoOoooOOOooOo && true")
	if err != nil {
		// ...
	}

	fmt.Println(v)
	// Output: true

}
