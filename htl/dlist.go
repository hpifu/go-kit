package htl

import (
	"bytes"
	"fmt"
)

type DListIterator struct {
	node *DListNode
}

func (it *DListIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.next
	return v
}

func (it *DListIterator) HasNext() bool {
	return it.node != nil
}

type DListRIterator struct {
	node *DListNode
}

func (it *DListRIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.prev
	return v
}

func (it *DListRIterator) HasNext() bool {
	return it.node != nil
}

type DListNode struct {
	val  interface{}
	next *DListNode
	prev *DListNode
}

func newDListNode(val interface{}, prev *DListNode, next *DListNode) *DListNode {
	return &DListNode{
		val:  val,
		prev: prev,
		next: next,
	}
}

func NewDList() *DList {
	return &DList{}
}

type DList struct {
	head *DListNode
	tail *DListNode
	size int
}

func (l DList) Iterator() *DListIterator {
	return &DListIterator{node: l.head}
}

func (l DList) RIterator() *DListIterator {
	return &DListIterator{node: l.tail}
}

func (l DList) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l DList) String() string {
	var buf bytes.Buffer

	node := l.head
	for node != l.tail {
		buf.WriteString(fmt.Sprint(node.val))
		buf.WriteString(" <-> ")
		node = node.next
	}
	buf.WriteString(fmt.Sprint(node.val))

	return buf.String()
}

func (l *DList) Len() int {
	return l.size
}

func (l *DList) Empty() bool {
	return l.size == 0
}

func (l *DList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.val
}

func (l *DList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.val
}

func (l *DList) PushFront(v interface{}) {
	if l.Empty() {
		node := newDListNode(v, nil, nil)
		l.head = node
		l.tail = node
	} else {
		node := newDListNode(v, nil, l.head)
		l.head.prev = node
		l.head = node
	}
	l.size++
}

func (l *DList) PushBack(v interface{}) {
	if l.Empty() {
		node := newDListNode(v, nil, nil)
		l.head = node
		l.tail = node
	} else {
		node := newDListNode(v, l.tail, nil)
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *DList) PopFront() interface{} {
	if l.Empty() {
		return nil
	}

	v := l.head.val
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--

	return v
}

func (l *DList) PopBack() interface{} {
	if l.Empty() {
		return nil
	}

	v := l.tail.val
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--

	return v
}
