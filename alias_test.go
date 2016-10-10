package arithmetic

import "fmt"

func ExampleRegisterAlias() {

	// Register french aliases.
	RegisterAlias("et", "and")
	RegisterAlias("si", "if")
	RegisterAlias("vrai", "true")

	v, err := Parse("si(vrai et vrai, \"ok\", \"ko\")")
	if err != nil {
		// ...
	}

	fmt.Println(v)

	// Output: ok
}
