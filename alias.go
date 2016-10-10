package arithmetic

import (
	"fmt"
	"strings"
)

// RegisterAlias aliases a variable or a function. This is useful if you want to use locale language. Aliases are case insensitive.
func RegisterAlias(label string, value string) {

	// Make case insensitive.
	label = strings.ToLower(label)
	value = strings.ToLower(value)

	// Ensure the label is not already registered.
	mustBeUnique(label)

	// Associate the alias to another, if any.
	if v, ok := aliases[value]; ok {
		aliases[label] = v
		return
	}

	// Associate the alias to a variable, if any.
	if v, ok := variables[value]; ok {
		variables[label] = v
		return
	}

	// Associate the alias to a function, if any.
	if v, ok := functions[value]; ok {
		functions[label] = v
		return
	}

	panic(fmt.Sprintf("no function, variable or alias matching %s", value))
}

var aliases = map[string]interface{}{
	"and": and{},
	"or":  or{},
	"not": not{},
}
