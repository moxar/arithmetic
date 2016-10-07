package arithmetic

import(
	"errors"
)

type divide struct{}

func (o divide) String() string {
	return "/"
}

func (o divide) precedence() uint8 {
	return 2
}

func (o divide) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	left, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, right)
	}
	
	if right == 0 {
		return nil, errors.New("division by 0")
	}

	return left / right, nil
}
