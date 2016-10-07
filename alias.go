package arithmetic

import (
	"fmt"
	"strings"
)

func RegisterAlias(label string, value string) {

	label = strings.ToLower(label)

	mustBeUnique(label)

	if v, ok := aliases[value]; ok {
		aliases[label] = v
		return
	}

	if v, ok := variables[value]; ok {
		variables[label] = v
		return
	}

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
