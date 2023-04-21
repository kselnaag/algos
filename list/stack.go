package list

type Snode[T any] struct {
	Val  T
	Next *Snode[T]
}

// ===========================
type Bag[T any] struct {
	first *Snode[T]
	size  int
}

func NewBag[T any]() Bag[T] {
	return Bag[T]{
		first: nil,
		size:  0,
	}
}

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) IsEmpty() bool {
	return b.Size() == 0
}

func (b *Bag[T]) Add(val T) {
	b.first = &Snode[T]{Val: val, Next: b.first}
	b.size++
}

func (b *Bag[T]) Iterate() []T {
	res := make([]T, 0, b.size)
	for node := b.first; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	return res
}

func (b *Bag[T]) Reverse() {
	b.first = Reverse(b.first)
}

// ===========================
type Stack[T any] struct {
	first *Snode[T]
	size  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		first: nil,
		size:  0,
	}
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(val T) {
	s.first = &Snode[T]{Val: val, Next: s.first}
	s.size++
}

func (s *Stack[T]) Pop() T {
	ret := s.first.Val
	s.first = s.first.Next
	s.size--
	return ret
}

func (s *Stack[T]) Iterate() []T {
	res := make([]T, 0, s.size)
	for node := s.first; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	return res
}

func (s *Stack[T]) Reverse() {
	s.first = Reverse(s.first)
}

// ===========================
type Queue[T any] struct {
	first *Snode[T]
	last  *Snode[T]
	size  int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Enq(val T) {
	newnode := &Snode[T]{Val: val, Next: nil}
	if q.size == 0 {
		q.first = newnode
	} else {
		q.last.Next = newnode
	}
	q.last = newnode
	q.size++
}

func (q *Queue[T]) Deq() T {
	ret := q.first.Val
	q.first = q.first.Next
	q.size--
	if (q.size == 0) || (q.size == 1) {
		q.last = q.first
	}
	return ret
}

func (q *Queue[T]) Iterate() []T {
	res := make([]T, 0, q.size)
	for node := q.first; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	return res
}

func (q *Queue[T]) Reverse() {
	q.last = q.first
	q.first = Reverse(q.first)
}

// ===========================
type Dnode[T any] struct {
	Val  T
	Prev *Dnode[T]
	Next *Dnode[T]
}

type Deque[T any] struct {
	first *Dnode[T]
	last  *Dnode[T]
	size  int
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Deque[T]) Front() T {
	return d.first.Val
}

func (d *Deque[T]) Back() T {
	return d.last.Val
}

func (d *Deque[T]) PushFront(val T) {
	newnode := &Dnode[T]{Val: val, Prev: nil, Next: d.first}
	if d.size == 0 {
		d.last = newnode
	} else {
		d.first.Prev = newnode
	}
	d.first = newnode
	d.size++
}

func (d *Deque[T]) PopFront() T {
	val := d.first.Val
	if d.size == 1 {
		d.last = nil
		d.first = nil
	} else {
		d.first = d.first.Next
		d.first.Prev = nil
	}
	d.size--
	return val
}

func (d *Deque[T]) PushBack(val T) {
	newnode := &Dnode[T]{Val: val, Prev: d.last, Next: nil}
	if d.size == 0 {
		d.first = newnode
	} else {
		d.last.Next = newnode
	}
	d.last = newnode
	d.size++
}

func (d *Deque[T]) PopBack() T {
	val := d.last.Val
	if d.size == 1 {
		d.last = nil
		d.first = nil
	} else {
		d.last = d.last.Prev
		d.last.Next = nil
	}
	d.size--
	return val
}

func (d *Deque[T]) Iterate() []T {
	res := make([]T, 0, d.size)
	for node := d.first; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	return res
}

func (d *Deque[T]) Reverse() {
	d.first, d.last = d.last, d.first
	for node := d.first; node != nil; node = node.Next {
		node.Next, node.Prev = node.Prev, node.Next
	}
}
