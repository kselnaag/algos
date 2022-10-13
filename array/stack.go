package array

import (
	_ "github.com/kselnaag/algos/types"
)

// ===========================
type Bag[T any] struct {
	arr  []T
	size int
}

func NewBag[T any]() Bag[T] {
	return Bag[T]{
		arr:  make([]T, 0, 8),
		size: 0,
	}
}

func (b *Bag[T]) Drop() {
	b.arr = make([]T, 0, 8)
	b.size = 0
}

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) IsEmpty() bool {
	return b.Size() == 0
}

func (b *Bag[T]) Add(val T) {
	b.arr = append(b.arr, val)
	b.size++
}

func (b *Bag[T]) Iterate() []T {
	return b.arr
}

func (b *Bag[T]) Reverse() {
	Reverse(b.arr)
}

// ===========================
type Stack[T any] struct {
	arr  []T
	size int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		arr:  make([]T, 0, 8),
		size: 0,
	}
}

func (s *Stack[T]) Drop() {
	s.arr = make([]T, 0, 8)
	s.size = 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack[T]) Pop() T {
	last := len(s.arr) - 1
	ret := s.arr[last]
	s.arr = s.arr[:last]
	s.size--
	return ret
}

func (s *Stack[T]) Iterate() []T {
	return s.arr
}

func (s *Stack[T]) Reverse() {
	Reverse(s.arr)
}

// ===========================
type Queue[T any] struct {
	arr  []T
	size int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		arr:  make([]T, 0, 8),
		size: 0,
	}
}

func (q *Queue[T]) Drop() {
	q.arr = make([]T, 0, 8)
	q.size = 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Enq(val T) {
	q.arr = append(q.arr, val)
	q.size++
}

func (q *Queue[T]) Deq() T {
	ret := q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	return ret
}

func (q *Queue[T]) Iterate() []T {
	return q.arr
}

func (q *Queue[T]) Reverse() {
	Reverse(q.arr)
}
