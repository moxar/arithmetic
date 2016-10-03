package arithmetic

import (
	"fmt"
)

type Token interface {
	fmt.Stringer
	Value() (Operand, Operator)
}

func tokensToString(input []Token) string {
	var output string
	for _, v := range input {
		output += " " + v.String()
	}
	return output
}
