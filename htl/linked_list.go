package htl

import (
	"bytes"
	"fmt"
)

type LinkedListIterator struct {
	node *LinkedListNode
}

func (it *LinkedListIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.next
	return v
}

func (it *LinkedListIterator) HasNext() bool {
	return it.node != nil
}

type LinkedListNode struct {
	val  interface{}
	next *LinkedListNode
}

func newLinkedListNode(val interface{}, next *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{
		val:  val,
		next: next,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func (l LinkedList) Iterator() *LinkedListIterator {
	return &LinkedListIterator{node: l.head}
}

func (l LinkedList) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l LinkedList) String() string {
	if l.Empty() {
		return ""
	}

	var buf bytes.Buffer

	node := l.head
	for node != l.tail {
		buf.WriteString(fmt.Sprint(node.val))
		buf.WriteString(" -> ")
		node = node.next
	}
	buf.WriteString(fmt.Sprint(node.val))

	return buf.String()
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.val
}

func (l *LinkedList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.val
}

func (l *LinkedList) PushFront(v interface{}) {
	if l.Empty() {
		node := newLinkedListNode(v, nil)
		l.head = node
		l.tail = node
	} else {
		node := newLinkedListNode(v, l.head)
		l.head = node
	}

	l.size++
}

func (l *LinkedList) PushBack(v interface{}) {
	if l.Empty() {
		node := newLinkedListNode(v, nil)
		l.head = node
		l.tail = node
	} else {
		l.tail.next = newLinkedListNode(v, nil)
		l.tail = l.tail.next
	}

	l.size++

}

func (l *LinkedList) PopFront() interface{} {
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

func (l *LinkedList) PopBack() interface{} {
	panic("PopBack is not support for performance reason, use DLinkedList instead")
}
