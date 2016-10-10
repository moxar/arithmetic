package arithmetic

import (
	"errors"
	"fmt"
)

// stack implementation.
type stack struct {
	values []interface{}
}

// pop returns the last inserted value from the stack and removes it.
func (s *stack) pop() (interface{}, bool) {
	last := len(s.values) - 1
	if last >= 0 {
		v := s.values[last]
		s.values = s.values[:last]
		return v, true
	}
	return nil, false
}

// push appends a value to the stack.
func (s *stack) push(v interface{}) {
	s.values = append(s.values, v)
}

// inc increments the last inserted value. This value must be an int.
func (s *stack) inc() {
	v, _ := s.popInt()
	s.push(v + 1)
}

// slice returns an slice of the last n elements of the stack. The elements are remove
// from the stack.
func (s *stack) slice(size int) ([]interface{}, error) {
	l := len(s.values)
	if l < size {
		return nil, fmt.Errorf("stack too small: %d element required out of %d", size, l)
	}

	out := s.values[l-size : l]
	s.values = s.values[:l-size]
	return out, nil
}

// popFloat is a helper that pops the stack and returns the element as float.
func (s *stack) popFloat() (float64, error) {
	v, ok := s.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	if f, ok := v.(variable); ok {
		f, ok := f.value.(float64)
		if ok {
			return f, nil
		}
	}

	f, ok := v.(float64)
	if !ok {
		return 0, fmt.Errorf("expected float, having %v (%T)", v, v)
	}

	return f, nil
}

// popInt is a helper that pops the stack and returns the element as int.
func (s *stack) popInt() (int, error) {
	v, ok := s.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	if i, ok := v.(variable); ok {
		i, ok := i.value.(int)
		if ok {
			return i, nil
		}
	}

	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("expected int, having %v (%T)", v, v)
	}

	return i, nil
}

// popBool is a helper that pops the stack and returns the element as bool.
func (s *stack) popBool() (bool, error) {
	v, ok := s.pop()
	if !ok {
		return false, errors.New("empty stack")
	}

	if b, ok := v.(variable); ok {
		b, ok := b.value.(bool)
		if ok {
			return b, nil
		}
	}

	b, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("expected boolean, having %v (%T)", v, v)
	}

	return b, nil
}
