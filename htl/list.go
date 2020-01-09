package htl

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type ListIterator interface {
	Next() interface{}
	HasNext() bool
	Prev() interface{}
	HasPrev() bool
}

type Predicate interface {
	Test() bool
}

type Collection interface {
	Size() int
	IsEmpty() bool
	Contains(v interface{}) bool
	Iterator() Iterator
	ToArray() []interface{}
	Add(v interface{})
	Remove(v interface{}) bool
	ContainsAll(collection Collection) bool
	AddAll(collection Collection) bool
	RemoveAll(collection Collection) bool
	RemoveIf(predicate Predicate) bool
	RetainAll(collection Collection) bool
	Clear()
	Equals(v interface{}) bool
}

type List interface {
	Collection
	//Size() int
	//IsEmpty() bool
	//Contains(v interface{}) bool
	//Iterator() Iterator
	//ToArray() []interface{}
	//Add(v interface{}) bool
	//Remove(v interface{}) bool
	//ContainsAll(collection Collection) bool
	//AddAll(collection Collection) bool
	//RemoveAll(collection Collection) bool
	//RetainAll(collection Collection) bool
	//Clear()
	Get(index int) interface{}
	Set(index int, v interface{}) interface{}
	AddAllIndex(index int, collection Collection) bool
	AddIndex(index int, v interface{}) bool
	RemoveIndex(index int) bool
	IndexOf(v interface{}) int
	LastIndexOf(v interface{}) int
	ListIterator() ListIterator
	SubList(fromIndex int, toIndex int)
}

type Queue interface {
	Add(v interface{})
	Offer(v interface{}) bool

	Remove() interface{}
	Poll() interface{}

	Element() interface{}
	Peek() interface{}
}

type DQueue interface {
	Queue

	AddFirst(v interface{})
	AddLast(v interface{})
	OfferFirst(v interface{}) bool
	OfferLast(v interface{}) bool

	RemoveFirst() interface{}
	RemoveLast() interface{}
	PollFirst() interface{}
	PollLast() interface{}

	RemoveFirstOccurrence(v interface{}) bool
	RemoveLastOccurrence(v interface{}) bool

	GetFirst() interface{}
	GetLast() interface{}
	PeekFirst() interface{}
	PeekLast() interface{}
}
