package htl

import (
	"bytes"
	"fmt"
)

type ArrayListIterator struct {
	vs  []interface{}
	idx int
}

func (it *ArrayListIterator) Next() interface{} {
	v := it.vs[it.idx]
	it.idx++
	return v
}

func (it *ArrayListIterator) HasNext() bool {
	return it.idx < len(it.vs)
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

type ArrayList struct {
	vs []interface{}
}

func (l ArrayList) Iterator() *ArrayListIterator {
	return &ArrayListIterator{vs: l.vs}
}

func (l ArrayList) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l ArrayList) String() string {
	if l.Empty() {
		return ""
	}

	var buf bytes.Buffer

	i := 0
	for ; i < len(l.vs)-1; i++ {
		buf.WriteString(fmt.Sprint(l.vs[i]))
		buf.WriteString(" | ")
	}
	buf.WriteString(fmt.Sprint(l.vs[i]))

	return buf.String()
}

func (l *ArrayList) Len() int {
	return len(l.vs)
}

func (l *ArrayList) Empty() bool {
	return len(l.vs) == 0
}

func (l *ArrayList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.vs[0]
}

func (l *ArrayList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.vs[len(l.vs)-1]
}

func (l *ArrayList) PushBack(v interface{}) {
	l.vs = append(l.vs, v)
}

func (l *ArrayList) PopBack() interface{} {
	if l.Empty() {
		return nil
	}
	v := l.vs[len(l.vs)-1]
	l.vs = l.vs[:len(l.vs)-1]
	return v
}

func (l *ArrayList) PushFront(v interface{}) {
	panic("PushFront is not support for performance reason, use LinkedList instead")
}

func (l *ArrayList) PopFront() interface{} {
	panic("PopFront is not support for performance reason, use LinkedList instead")
}
