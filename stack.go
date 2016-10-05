package arithmetic

import (
	"errors"
	"fmt"
)

type stack struct {
	values []interface{}
}

func (s *stack) pop() (interface{}, bool) {
	last := len(s.values) - 1
	if last >= 0 {
		v := s.values[last]
		s.values = s.values[:last]
		return v, true
	}
	return nil, false
}

func (s *stack) push(v interface{}) {
	s.values = append(s.values, v)
}

func (s *stack) inc() {
	v, _ := s.popInt()
	s.push(v + 1)
}

func (s *stack) slice(size int) ([]interface{}, error) {
	l := len(s.values)
	if l < size {
		return nil, fmt.Errorf("stack too small: %d element required out of %d", size, l)
	}

	out := s.values[l-size:l]
	s.values = s.values[:l-size]
	return out, nil
}

func (s *stack) popFloat() (float64, error) {
	v, ok := s.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	f, ok := v.(float64)
	if !ok {
		return 0, fmt.Errorf("expected float, having %v (%T)", v, v)
	}

	return f, nil
}

func (s *stack) popInt() (int, error) {
	v, ok := s.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("expected int, having %v (%T)", v, v)
	}

	return i, nil
}
