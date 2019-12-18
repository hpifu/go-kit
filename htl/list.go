package htl

import (
	"bytes"
	"fmt"
)

type ListIterator struct {
	node *ListNode
}

func (it *ListIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.next
	return v
}

func (it *ListIterator) HasNext() bool {
	return it.node != nil
}

type ListNode struct {
	val  interface{}
	next *ListNode
}

func newListNode(val interface{}, next *ListNode) *ListNode {
	return &ListNode{
		val:  val,
		next: next,
	}
}

func NewList() *List {
	return &List{}
}

type List struct {
	head *ListNode
	tail *ListNode
	size int
}

func (l List) Iterator() *ListIterator {
	return &ListIterator{node: l.head}
}

func (l List) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l List) String() string {
	var buf bytes.Buffer

	node := l.head
	for node != l.tail {
		buf.WriteString(fmt.Sprintf("%v", node.val))
		buf.WriteString("->")
		node = node.next
	}
	buf.WriteString(fmt.Sprintf("%v", node.val))

	return buf.String()
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.val
}

func (l *List) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.val
}

func (l *List) PushFront(v interface{}) {
	if l.Empty() {
		node := newListNode(v, nil)
		l.head = node
		l.tail = node
	} else {
		node := newListNode(v, l.head)
		l.head = node
	}

	l.size++
}

func (l *List) PushBack(v interface{}) {
	if l.Empty() {
		node := newListNode(v, nil)
		l.head = node
		l.tail = node
	} else {
		l.tail.next = newListNode(v, nil)
		l.tail = l.tail.next
	}

	l.size++

}

func (l *List) PopFront() interface{} {
	if l.Empty() {
		return nil
	}

	v := l.head.val
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
	}

	l.size--
	return v
}

func (l *List) PopBack() interface{} {
	panic("PopBack is not support for performance reason, use DList instead")
}
