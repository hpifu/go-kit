package htl

func NewQueue() *Queue {
	return &Queue{
		l: NewLinkedList(),
	}
}

type Queue struct {
	l *LinkedList
}

func (q *Queue) Top() interface{} {
	return q.l.Front()
}

func (q *Queue) Push(v interface{}) {
	q.l.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	return q.l.PopFront()
}

func (q *Queue) Empty() bool {
	return q.l.Empty()
}

func (q *Queue) Len() int {
	return q.l.Len()
}

func (q Queue) String() string {
	return q.l.String()
}
