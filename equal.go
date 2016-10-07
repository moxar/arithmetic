package arithmetic

type equal struct{}

func (o equal) String() string {
	return "=="
}

func (o equal) precedence() uint8 {
	return precedenceEqual
}

func (o equal) solve(st *stack) (interface{}, error) {
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	return eq(left, right), nil
}

type different struct{}

func (o different) String() string {
	return "!="
}

func (o different) precedence() uint8 {
	return precedenceEqual
}

func (o different) solve(st *stack) (interface{}, error) {
	right, ok := st.pop()
	if !ok {
		return nil, rightError(o)
	}

	left, ok := st.pop()
	if !ok {
		return nil, leftError(o, right)
	}

	return !eq(left, right), nil
}
