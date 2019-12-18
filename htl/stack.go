package htl

func NewStack() *Stack {
	return &Stack{
		vs: make([]interface{}, 0, 8),
	}
}

type Stack struct {
	vs []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.vs = append(s.vs, v)
}

func (s *Stack) Top() interface{} {
	return s.vs[len(s.vs)-1]
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.vs[len(s.vs)-1]
	s.vs = s.vs[:len(s.vs)-1]
	return v
}

func (s *Stack) Empty() bool {
	return len(s.vs) == 0
}

func (s *Stack) Len() int {
	return len(s.vs)
}
