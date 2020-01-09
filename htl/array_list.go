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

func (it *ArrayListIterator) Prev() interface{} {
	v := it.vs[it.idx]
	it.idx--
	return v
}

func (it *ArrayListIterator) HasPrev() bool {
	return it.idx >= 0
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

type ArrayList struct {
	vs []interface{}
}

func (l ArrayList) String() string {
	if l.IsEmpty() {
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

func (l ArrayList) Iterator() Iterator {
	return &ArrayListIterator{vs: l.vs}
}

func (l ArrayList) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l *ArrayList) Size() int {
	return len(l.vs)
}

func (l *ArrayList) IsEmpty() bool {
	return len(l.vs) == 0
}

func (l *ArrayList) Contains(v interface{}) bool {
	return l.IndexOf(v) >= 0
}

func (l *ArrayList) IndexOf(v interface{}) int {
	for i, val := range l.vs {
		if val == v {
			return i
		}
	}
	return -1
}

func (l *ArrayList) ToArray() []interface{} {
	return l.vs
}

func (l *ArrayList) Add(v interface{}) {
	l.AddLast(v)
}

func (l *ArrayList) Remove(v interface{}) bool {
	i := l.IndexOf(v)
	if i == -1 {
		return false
	}
	for ; i < len(l.vs); i++ {
		l.vs[i] = l.vs[i+1]
	}
	l.vs = l.vs[:len(l.vs)-1]

	return true
}

func (l *ArrayList) Clear() {
	l.vs = l.vs[:0]
}

func (l *ArrayList) GetFirst() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.vs[0]
}

func (l *ArrayList) GetLast() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.vs[len(l.vs)-1]
}

func (l *ArrayList) AddLast(v interface{}) {
	l.vs = append(l.vs, v)
}

func (l *ArrayList) RemoveLast() interface{} {
	if l.IsEmpty() {
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
