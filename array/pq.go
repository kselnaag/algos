package array

import (
	I "github.com/kselnaag/algos/types"
)

// ===========================
type MinPQ[T any] struct {
	pq   []T
	size int
}

func NewMinPQ[T any]() MinPQ[T] {
	return MinPQ[T]{
		pq:   make([]T, 1, 8),
		size: 0,
	}
}

func (min *MinPQ[T]) Size() int {
	return min.size
}

func (min *MinPQ[T]) IsEmpty() bool {
	return min.Size() == 0
}

func (min *MinPQ[T]) swim(k int) {
	for (k > 1) && I.GT(min.pq[k/2], min.pq[k]) {
		swap(min.pq, k/2, k)
		k /= 2
	}
}
func (min *MinPQ[T]) sink(k int) {
	for (2 * k) <= min.size {
		j := 2 * k
		if (j < min.size) && I.GT(min.pq[j], min.pq[j+1]) {
			j++
		}
		if !I.GT(min.pq[k], min.pq[j]) {
			break
		}
		swap(min.pq, k, j)
		k = j
	}
}

func (min *MinPQ[T]) Add(val T) {
	min.size++
	if len(min.pq) > min.size {
		min.pq[min.size] = val
	} else {
		min.pq = append(min.pq, val)
	}
	min.swim(min.size)
}

func (min *MinPQ[T]) Min() T {
	if min.size < 1 {
		panic("algos.array.(MinPQ).Min(): Queue is empty ")
	}
	return min.pq[1]
}

func (min *MinPQ[T]) GetMin() T {
	if min.size < 1 {
		panic("algos.array.(MinPQ).GetMin(): Queue is empty ")
	}
	ret := min.pq[1]
	swap(min.pq, 1, min.size)
	min.size--
	min.sink(1)
	return ret
}

func (min *MinPQ[T]) Iterate() []T {
	return min.pq[1:(min.size + 1)]
}

// ===========================
type MaxPQ[T any] struct {
	pq   []T
	size int
}

func NewMaxPQ[T any]() MaxPQ[T] {
	return MaxPQ[T]{
		pq:   make([]T, 1, 8),
		size: 0,
	}
}

func (max *MaxPQ[T]) Size() int {
	return max.size
}

func (max *MaxPQ[T]) IsEmpty() bool {
	return max.Size() == 0
}

func (max *MaxPQ[T]) swim(k int) {
	for (k > 1) && I.LT(max.pq[k/2], max.pq[k]) {
		swap(max.pq, k/2, k)
		k /= 2
	}
}
func (max *MaxPQ[T]) sink(k int) {
	for (2 * k) <= max.size {
		j := 2 * k
		if (j < max.size) && I.LT(max.pq[j], max.pq[j+1]) {
			j++
		}
		if !I.LT(max.pq[k], max.pq[j]) {
			break
		}
		swap(max.pq, k, j)
		k = j
	}
}

func (max *MaxPQ[T]) Add(val T) {
	max.size++
	if len(max.pq) > max.size {
		max.pq[max.size] = val
	} else {
		max.pq = append(max.pq, val)
	}
	max.swim(max.size)
}

func (max *MaxPQ[T]) Max() T {
	if max.size < 1 {
		panic("algos.array.(MaxPQ).Max(): Queue is empty ")
	}
	return max.pq[1]
}

func (max *MaxPQ[T]) GetMax() T {
	if max.size < 1 {
		panic("algos.array.(MaxPQ).GetMax(): Queue is empty ")
	}
	ret := max.pq[1]
	swap(max.pq, 1, max.size)
	max.size--
	max.sink(1)
	return ret
}

func (max *MaxPQ[T]) Iterate() []T {
	return max.pq[1:(max.size + 1)]
}
