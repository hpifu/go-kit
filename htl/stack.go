package htl

func NewStack() *Stack {
	return &Stack{
		l: NewArrayList(),
	}
}

type Stack struct {
	l *ArrayList
}

func (s *Stack) Push(v interface{}) {
	s.l.PushBack(v)
}

func (s *Stack) Top() interface{} {
	return s.l.Back()
}

func (s *Stack) Pop() interface{} {
	return s.l.PopBack()
}

func (s *Stack) Empty() bool {
	return s.l.Empty()
}

func (s *Stack) Len() int {
	return s.l.Len()
}
