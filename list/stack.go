package list

type node[T any] struct {
	val  T
	next *node[T]
}

func Reverse[T any](first *node[T]) *node[T] {
	if first == nil {
		return nil
	}
	if first.next == nil {
		return first
	}
	second := first.next
	root := Reverse(second)
	second.next = first
	first.next = nil
	return root
}

// ===========================
type Bag[T any] struct {
	first *node[T]
	size  int
	fwd   *node[T]
}

func NewBag[T any]() Bag[T] {
	return Bag[T]{
		first: nil,
		size:  0,
		fwd:   nil,
	}
}

func (b *Bag[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) Add(val T) {
	first := b.first
	b.first = &node[T]{val: val, next: first}
	b.size++
	b.fwd = b.first
}

func (b *Bag[T]) Next() T {
	ret := b.fwd.val
	if b.fwd.next == nil {
		b.fwd = b.first
	} else {
		b.fwd = b.fwd.next
	}
	return ret
}

func (b *Bag[T]) Reverse() {
	b.fwd = Reverse(b.first)
	b.first = b.fwd
}

func (b *Bag[T]) Drop() {
	b.first = nil
	b.size = 0
	b.fwd = nil
}

// ===========================
type Stack[T any] struct {
	first *node[T]
	size  int
	fwd   *node[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		first: nil,
		size:  0,
		fwd:   nil,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(val T) {
	first := s.first
	s.first = &node[T]{val: val, next: first}
	s.size++
	s.fwd = s.first
}

func (s *Stack[T]) Pop() T {
	ret := s.first.val
	s.first = s.first.next
	s.size--
	s.fwd = s.first
	return ret
}

func (s *Stack[T]) Next() T {
	ret := s.fwd.val
	if s.fwd.next == nil {
		s.fwd = s.first
	} else {
		s.fwd = s.fwd.next
	}
	return ret
}

func (s *Stack[T]) Reverse() {
	s.fwd = Reverse(s.first)
	s.first = s.fwd
}
func (s *Stack[T]) Drop() {
	s.first = nil
	s.size = 0
	s.fwd = nil
}

// ===========================
type Queue[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
	fwd   *node[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		first: nil,
		last:  nil,
		size:  0,
		fwd:   nil,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Enq(val T) {
	newnode := &node[T]{val: val, next: nil}
	if q.size == 0 {
		q.first = newnode
	} else {
		q.last.next = newnode
	}
	q.last = newnode
	q.size++
	q.fwd = q.first
}

func (q *Queue[T]) Deq() T {
	ret := q.first.val
	q.first = q.first.next
	q.size--
	if (q.size == 0) || (q.size == 1) {
		q.last = q.first
	}
	q.fwd = q.first
	return ret
}

func (q *Queue[T]) Next() T {
	ret := q.fwd.val
	if q.fwd.next == nil {
		q.fwd = q.first
	} else {
		q.fwd = q.fwd.next
	}
	return ret
}

func (q *Queue[T]) Reverse() {
	q.last = q.first
	q.fwd = Reverse(q.first)
	q.first = q.fwd
}

func (q *Queue[T]) Drop() {
	q.first = nil
	q.size = 0
	q.fwd = nil
}
