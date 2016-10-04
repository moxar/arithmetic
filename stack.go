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

type OperandStack struct {
	stack
}

func (s *OperandStack) Pop() (Operand, bool) {
	v, ok := s.stack.pop()
	if !ok {
		return nil, false
	}
	return v.(Operand), true
}

func (s *OperandStack) Push(v Operand) {
	s.stack.push(v)
}

func (s *OperandStack) PopInt() (int, error) {
	op, ok := s.stack.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	v, ok := op.(Number)
	if !ok {
		return 0, fmt.Errorf("%s is not a numeric value", op)
	}

	return int(v), nil
}

func (s *OperandStack) PopFloat() (float64, error) {
	op, ok := s.stack.pop()
	if !ok {
		return 0, errors.New("empty stack")
	}

	v, ok := op.(Number)
	if !ok {
		return 0, fmt.Errorf("%s is not a numeric value", op)
	}

	return float64(v), nil
}

func (s *OperandStack) PopBool() (bool, error) {
	op, ok := s.stack.pop()
	if !ok {
		return false, errors.New("empty stack")
	}

	v, ok := op.(Boolean)
	if !ok {
		return false, fmt.Errorf("%s is not a boolean", op)
	}

	return bool(v), nil
}

type OperatorStack struct {
	stack
}

func (s *OperatorStack) Pop() (Operator, bool) {
	v, ok := s.stack.pop()
	if !ok {
		return nil, false
	}
	return v.(Operator), true
}

func (s *OperatorStack) Push(v Operator) {
	s.stack.push(v)
}

type ArityStack struct {
	stack
}

func (s *ArityStack) Pop() int {
	v, _ := s.stack.pop()
	return v.(int)
}

func (s *ArityStack) Push(v int) {
	s.stack.push(v)
}

func (s *ArityStack) Inc() {
	s.stack.push(s.Pop() + 1)
}
