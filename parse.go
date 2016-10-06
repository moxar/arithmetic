package arithmetic

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
