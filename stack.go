package arithmetic

type stack struct{
	values []interface{}
}

func (s *stack) pop() (interface{}, bool) {
	if len(s.values) > 1 {
		v := s.values[len(s.values)-1]
		s.values = s.values[0:len(s.values)-1]
		return v, true
	}
	return nil, false
}

func (s *stack) push(v interface{}) {
	s.values = append(s.values, v)
}

type OperandStack struct{
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

type OperatorStack struct{
	stack
}

func (s *OperatorStack) Pop() (Operand, bool) {
	v, ok := s.stack.pop()
	if !ok {
		return nil, false
	}
	return v.(Operand), true
}

func (s *OperatorStack) Push(v Operand) {
	s.stack.push(v)
}
