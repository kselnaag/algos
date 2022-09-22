package array

import (
	_ "github.com/kselnaag/algos/types"
)

// ===========================
type Bag[T any] struct {
	arr  []T
	size int
	fwd  int
}

func NewBag[T any]() Bag[T] {
	return Bag[T]{
		arr:  make([]T, 0, 8),
		size: 0,
		fwd:  -1,
	}
}

func (b *Bag[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) Add(val T) {
	b.arr = append(b.arr, val)
	b.size++
	b.fwd = -1
}

func (b *Bag[T]) Next() T {
	b.fwd++
	if b.fwd == b.size {
		b.fwd = 0
	}
	return b.arr[b.fwd]
}

func (b *Bag[T]) Reverse() {
	Reverse(b.arr)
	b.fwd = -1
}

func (b *Bag[T]) Drop() {
	b.arr = b.arr[:0]
	b.fwd = -1
	b.size = 0
}

// ===========================
type Stack[T any] struct {
	arr  []T
	size int
	fwd  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		arr:  make([]T, 0, 8),
		size: 0,
		fwd:  -1,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
	s.size++
	s.fwd = -1
}

func (s *Stack[T]) Pop() T {
	last := len(s.arr) - 1
	ret := s.arr[last]
	s.arr = s.arr[:last]
	s.size--
	s.fwd = -1
	return ret
}

func (s *Stack[T]) Next() T {
	s.fwd++
	if s.fwd == s.size {
		s.fwd = 0
	}
	return s.arr[s.fwd]
}

func (s *Stack[T]) Reverse() {
	Reverse(s.arr)
	s.fwd = -1
}

func (s *Stack[T]) Drop() {
	s.arr = s.arr[:0]
	s.fwd = -1
	s.size = 0
}

// ===========================
type Queue[T any] struct {
	arr  []T
	size int
	fwd  int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		arr:  make([]T, 0, 8),
		size: 0,
		fwd:  -1,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Enq(val T) {
	q.arr = append(q.arr, val)
	q.size++
	q.fwd = -1
}

func (q *Queue[T]) Deq() T {
	ret := q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	q.fwd = -1
	return ret
}

func (q *Queue[T]) Next() T {
	q.fwd++
	if q.fwd == q.size {
		q.fwd = 0
	}
	return q.arr[q.fwd]
}

func (q *Queue[T]) Reverse() {
	Reverse(q.arr)
	q.fwd = -1
}

func (q *Queue[T]) Drop() {
	q.arr = q.arr[:0]
	q.fwd = -1
	q.size = 0
}
