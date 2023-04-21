package array

import (
	I "github.com/kselnaag/algos/types"
)

func swimLT[T any](arr []T, k int) {
	for (k > 1) && I.GT(arr[k/2], arr[k]) {
		swap(arr, k/2, k)
		k /= 2
	}
}

func sinkLT[T any](arr []T, k, size int) {
	for (2 * k) <= size {
		j := 2 * k
		if (j < size) && I.GT(arr[j], arr[j+1]) {
			j++
		}
		if !I.GT(arr[k], arr[j]) {
			break
		}
		swap(arr, k, j)
		k = j
	}
}

func swimGT[T any](arr []T, k int) {
	for (k > 1) && I.LT(arr[k/2], arr[k]) {
		swap(arr, k/2, k)
		k /= 2
	}
}

func sinkGT[T any](arr []T, k, size int) {
	for (2 * k) <= size {
		j := 2 * k
		if (j < size) && I.LT(arr[j], arr[j+1]) {
			j++
		}
		if !I.LT(arr[k], arr[j]) {
			break
		}
		swap(arr, k, j)
		k = j
	}
}

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
	swimLT(min.pq, k)
}

func (min *MinPQ[T]) sink(k int) {
	sinkLT(min.pq, k, min.size)
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
	swimGT(max.pq, k)
}

func (max *MaxPQ[T]) sink(k int) {
	sinkGT(max.pq, k, max.size)
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

// ===========================
type LRU[T any] struct {
	arr []T
}

func NewLRU[T any](size int) LRU[T] {
	return LRU[T]{
		arr: make([]T, 0, size),
	}
}

func (lru *LRU[T]) Cap() int {
	return cap(lru.arr)
}

func (lru *LRU[T]) Size() int {
	return len(lru.arr)
}

func (lru *LRU[T]) IsEmpty() bool {
	return lru.Size() == 0
}

func (lru *LRU[T]) Iterate() []T {
	return lru.arr
}

func (lru *LRU[T]) binsrchLT(arr []T, pidx int) int {
	low := 0
	high := len(arr) - 1
	elem := arr[pidx]
	mid := 0
	for low < high {
		mid = (low + ((high - low) / 2))
		val := arr[mid]
		switch {
		case I.EQ(val, elem):
			return mid
		case I.GT(val, elem):
			high = mid
		case I.LT(val, elem):
			low = mid + 1
		default:
			panic("algos.array.(LRU).binsrchLT(arr, pidx): comparison failed")
		}
	}
	return mid
}

func (lru *LRU[T]) sortpoint(pidx int) {
	val := lru.arr[pidx]
	inspos := lru.binsrchLT(lru.arr, pidx)
	if I.LT(lru.arr[inspos], val) {
		inspos++
	}
	for i := lru.Size() - 2; i >= inspos; i-- {
		lru.arr[i+1] = lru.arr[i]
	}
	lru.arr[inspos] = val
}

func (lru *LRU[T]) Set(val T) {
	if lru.Size() == lru.Cap() {
		if !I.LT(val, lru.arr[lru.Size()-1]) {
			return
		}
		lru.arr[lru.Size()-1] = val
	} else {
		lru.arr = append(lru.arr, val)
	}
	lru.sortpoint(lru.Size() - 1)
}

// ===========================
type MRU[T any] struct {
	arr []T
}

func NewMRU[T any](size int) MRU[T] {
	return MRU[T]{
		arr: make([]T, 0, size),
	}
}

func (mru *MRU[T]) Cap() int {
	return cap(mru.arr)
}

func (mru *MRU[T]) Size() int {
	return len(mru.arr)
}

func (mru *MRU[T]) IsEmpty() bool {
	return mru.Size() == 0
}

func (mru *MRU[T]) Iterate() []T {
	return mru.arr
}

func (mru *MRU[T]) binsrchGT(arr []T, pidx int) int {
	low := 0
	high := len(arr) - 1
	elem := arr[pidx]
	mid := 0
	for low < high {
		mid = (low + ((high - low) / 2))
		val := arr[mid]
		switch {
		case I.EQ(val, elem):
			return mid
		case I.LT(val, elem):
			high = mid
		case I.GT(val, elem):
			low = mid + 1
		default:
			panic("algos.array.(LRU).binsrchLT(arr, pidx): comparison failed")
		}
	}
	return mid
}

func (mru *MRU[T]) sortpoint(pidx int) {
	val := mru.arr[pidx]
	inspos := mru.binsrchGT(mru.arr, pidx)
	if I.GT(mru.arr[inspos], val) {
		inspos++
	}
	for i := mru.Size() - 2; i >= inspos; i-- {
		mru.arr[i+1] = mru.arr[i]
	}
	mru.arr[inspos] = val
}

func (mru *MRU[T]) Set(val T) {
	if mru.Size() == mru.Cap() {
		if !I.GT(val, mru.arr[mru.Size()-1]) {
			return
		}
		mru.arr[mru.Size()-1] = val
	} else {
		mru.arr = append(mru.arr, val)
	}
	mru.sortpoint(mru.Size() - 1)
}
