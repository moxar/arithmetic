package arithmetic

import (
	"fmt"
)

func ExampleRegisterVariable() {

	// Register a new variable.
	RegisterVariable("dayInYear", 365)

	v, err := Parse("dayInYear * 2")
	if err != nil {
		// ...
	}

	fmt.Println(v)

	// Output: 730

}
