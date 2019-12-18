package htl

import (
	"bytes"
	"fmt"
)

type DLinkedListIterator struct {
	node *DLinkedListNode
}

func (it *DLinkedListIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.next
	return v
}

func (it *DLinkedListIterator) HasNext() bool {
	return it.node != nil
}

type DLinkedListRIterator struct {
	node *DLinkedListNode
}

func (it *DLinkedListRIterator) Next() interface{} {
	v := it.node.val
	it.node = it.node.prev
	return v
}

func (it *DLinkedListRIterator) HasNext() bool {
	return it.node != nil
}

type DLinkedListNode struct {
	val  interface{}
	next *DLinkedListNode
	prev *DLinkedListNode
}

func newDLinkedListNode(val interface{}, prev *DLinkedListNode, next *DLinkedListNode) *DLinkedListNode {
	return &DLinkedListNode{
		val:  val,
		prev: prev,
		next: next,
	}
}

func NewDLinkedList() *DLinkedList {
	return &DLinkedList{}
}

type DLinkedList struct {
	head *DLinkedListNode
	tail *DLinkedListNode
	size int
}

func (l DLinkedList) Iterator() *DLinkedListIterator {
	return &DLinkedListIterator{node: l.head}
}

func (l DLinkedList) RIterator() *DLinkedListIterator {
	return &DLinkedListIterator{node: l.tail}
}

func (l DLinkedList) ForEach(op func(interface{})) {
	it := l.Iterator()
	for it.HasNext() {
		op(it.Next())
	}
}

func (l DLinkedList) String() string {
	if l.Empty() {
		return ""
	}

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

func (l *DLinkedList) Len() int {
	return l.size
}

func (l *DLinkedList) Empty() bool {
	return l.size == 0
}

func (l *DLinkedList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.val
}

func (l *DLinkedList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.val
}

func (l *DLinkedList) PushFront(v interface{}) {
	if l.Empty() {
		node := newDLinkedListNode(v, nil, nil)
		l.head = node
		l.tail = node
	} else {
		node := newDLinkedListNode(v, nil, l.head)
		l.head.prev = node
		l.head = node
	}
	l.size++
}

func (l *DLinkedList) PushBack(v interface{}) {
	if l.Empty() {
		node := newDLinkedListNode(v, nil, nil)
		l.head = node
		l.tail = node
	} else {
		node := newDLinkedListNode(v, l.tail, nil)
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *DLinkedList) PopFront() interface{} {
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

func (l *DLinkedList) PopBack() interface{} {
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
