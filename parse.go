package arithmetic

// Parse the input and calculates the value. The parser uses a shunting-yard algorithm
// that transforms the input from infix to postfix notation (aka reverse-polish notation).
// The input can contain spaces. They will be used to separate tokens, but are not mandatory.
//
// More about the shunting-yard algorithm here:
// http://wcipeg.com/wiki/Shunting_yard_algorithm
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func Parse(input string) (interface{}, error) {
	infix, err := tokenize(input)
	if err != nil {
		return nil, err
	}

	postfix, err := shuntingYard(infix)
	if err != nil {
		return nil, err
	}

	value, err := solve(postfix)
	if err != nil {
		return nil, err
	}

	return value, nil
}
