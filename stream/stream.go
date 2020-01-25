package stream

import (
	"sort"
)

func NewIterable(supplier Supplier) *Iterable {
	return &Iterable{
		Next: supplier,
	}
}

type Iterable struct {
	Next func() interface{}
}

func NewStream(iterable *Iterable) *Stream {
	return &Stream{
		iter: iterable,
	}
}

func Of(vs ...interface{}) *Stream {
	i := -1
	return NewStream(NewIterable(func() interface{} {
		i++
		if i < len(vs) {
			return vs[i]
		}
		return nil
	}))
}

type Stream struct {
	iter *Iterable
	ops  []func(interface{}) (interface{}, bool)
	eop  func(iterable Iterable) interface{}
	idx  int
}

func (s *Stream) Map(function func(interface{}) interface{}) *Stream {
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		return function(v), true
	})
	return s
}

func (s *Stream) Filter(predicate func(interface{}) bool) *Stream {
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		return v, predicate(v)
	})
	return s
}

func (s *Stream) Limit(num int) *Stream {
	idx := -1
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		idx++
		if idx >= num {
			return nil, false
		}
		return v, true
	})

	return s
}

func (s *Stream) Skip(num int) *Stream {
	idx := -1
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		idx++
		if idx < num {
			return v, false
		}

		return v, true
	})

	return s
}

type Predicate func(interface{}) bool
type Consumer func(interface{})
type Comparator func(interface{}, interface{}) int
type Supplier func() interface{}
type BinaryOperator func(interface{}, interface{}) interface{}

func (s *Stream) TakeWhile(predicate Predicate) *Stream {
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		if predicate(v) {
			return v, true
		}
		return nil, false
	})
	return s
}

func (s *Stream) DropWhile(predicate Predicate) *Stream {
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		return v, !predicate(v)
	})
	return s
}

func (s *Stream) Distinct() *Stream {
	set := map[interface{}]struct{}{}
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		if _, ok := set[v]; ok {
			return v, false
		}

		set[v] = struct{}{}
		return v, true
	})
	return s
}

func (s *Stream) Peek(consumer Consumer) *Stream {
	s.ops = append(s.ops, func(v interface{}) (interface{}, bool) {
		consumer(v)
		return v, true
	})

	return s
}

func (s *Stream) Sorted(comparator Comparator) *Stream {
	vs := s.ToSlice()
	sort.Slice(vs, func(i, j int) bool {
		return comparator(vs[i], vs[j]) < 0
	})

	i := -1
	return NewStream(NewIterable(func() interface{} {
		i++
		if i < len(vs) {
			return vs[i]
		}
		return nil
	}))
}

func (s *Stream) Next() interface{} {
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		return v
	}

	return nil
}

func (s *Stream) ForEach(consumer Consumer) {
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		consumer(v)
	}
}

func (s *Stream) ToSlice() []interface{} {
	var vs []interface{}
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		vs = append(vs, v)
	}

	return vs
}

func (s *Stream) Max(comparator Comparator) interface{} {
	var max interface{}
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		if max == nil || comparator(max, v) < 0 {
			max = v
		}
	}

	return max
}

func (s *Stream) Min(comparator Comparator) interface{} {
	var min interface{}
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		if min == nil || comparator(min, v) > 0 {
			min = v
		}
	}

	return min
}

func (s *Stream) AnyMatch(predicate Predicate) bool {
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		if predicate(v) {
			return true
		}
	}

	return false
}

func (s *Stream) AllMatch(predicate Predicate) bool {
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		if !predicate(v) {
			return false
		}
	}

	return true
}

func (s *Stream) NoneMatch(predicate Predicate) bool {
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		if predicate(v) {
			return false
		}
	}

	return true
}

func (s *Stream) Count() int {
	num := 0
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		num++
	}

	return num
}

func (s *Stream) Reduce(binaryOperator BinaryOperator, initialValue interface{}) interface{} {
	x := initialValue
	for v := s.iter.Next(); v != nil; v = s.iter.Next() {
		ok := true
		for _, op := range s.ops {
			v, ok = op(v)
			if !ok {
				break
			}
		}
		if v == nil {
			break
		}
		if !ok {
			continue
		}

		x = binaryOperator(x, v)
	}

	return x
}
