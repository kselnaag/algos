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
}

func NewBag[T any]() Bag[T] {
	return Bag[T]{
		first: nil,
		size:  0,
	}
}

func (b *Bag[T]) Drop() {
	b.first = nil
	b.size = 0
}

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) IsEmpty() bool {
	return b.Size() == 0
}

func (b *Bag[T]) Add(val T) {
	b.first = &node[T]{val: val, next: b.first}
	b.size++
}

func (b *Bag[T]) Iterate() []T {
	res := make([]T, 0, b.size)
	for node := b.first; node != nil; node = node.next {
		res = append(res, node.val)
	}
	return res
}

func (b *Bag[T]) Reverse() {
	b.first = Reverse(b.first)
}

// ===========================
type Stack[T any] struct {
	first *node[T]
	size  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		first: nil,
		size:  0,
	}
}

func (s *Stack[T]) Drop() {
	s.first = nil
	s.size = 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(val T) {
	s.first = &node[T]{val: val, next: s.first}
	s.size++
}

func (s *Stack[T]) Pop() T {
	ret := s.first.val
	s.first = s.first.next
	s.size--
	return ret
}

func (s *Stack[T]) Iterate() []T {
	res := make([]T, 0, s.size)
	for node := s.first; node != nil; node = node.next {
		res = append(res, node.val)
	}
	return res
}

func (s *Stack[T]) Reverse() {
	s.first = Reverse(s.first)
}

// ===========================
type Queue[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *Queue[T]) Drop() {
	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
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
}

func (q *Queue[T]) Deq() T {
	ret := q.first.val
	q.first = q.first.next
	q.size--
	if (q.size == 0) || (q.size == 1) {
		q.last = q.first
	}
	return ret
}

func (q *Queue[T]) Iterate() []T {
	res := make([]T, 0, q.size)
	for node := q.first; node != nil; node = node.next {
		res = append(res, node.val)
	}
	return res
}

func (q *Queue[T]) Reverse() {
	q.last = q.first
	q.first = Reverse(q.first)
}
