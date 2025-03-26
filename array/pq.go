package array

import (
	I "algos/types"
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

func (minpq *MinPQ[T]) Size() int {
	return minpq.size
}

func (minpq *MinPQ[T]) IsEmpty() bool {
	return minpq.Size() == 0
}

func (minpq *MinPQ[T]) swim(k int) {
	swimLT(minpq.pq, k)
}

func (minpq *MinPQ[T]) sink(k int) {
	sinkLT(minpq.pq, k, minpq.size)
}

func (minpq *MinPQ[T]) Add(val T) {
	minpq.size++
	if len(minpq.pq) > minpq.size {
		minpq.pq[minpq.size] = val
	} else {
		minpq.pq = append(minpq.pq, val)
	}
	minpq.swim(minpq.size)
}

func (minpq *MinPQ[T]) Min() T {
	if minpq.size < 1 {
		panic("algos.array.(MinPQ).Min(): Queue is empty ")
	}
	return minpq.pq[1]
}

func (minpq *MinPQ[T]) GetMin() T {
	if minpq.size < 1 {
		panic("algos.array.(MinPQ).GetMin(): Queue is empty ")
	}
	ret := minpq.pq[1]
	swap(minpq.pq, 1, minpq.size)
	minpq.size--
	minpq.sink(1)
	return ret
}

func (minpq *MinPQ[T]) Iterate() []T {
	return minpq.pq[1:(minpq.size + 1)]
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

func (maxpq *MaxPQ[T]) Size() int {
	return maxpq.size
}

func (maxpq *MaxPQ[T]) IsEmpty() bool {
	return maxpq.Size() == 0
}

func (maxpq *MaxPQ[T]) swim(k int) {
	swimGT(maxpq.pq, k)
}

func (maxpq *MaxPQ[T]) sink(k int) {
	sinkGT(maxpq.pq, k, maxpq.size)
}

func (maxpq *MaxPQ[T]) Add(val T) {
	maxpq.size++
	if len(maxpq.pq) > maxpq.size {
		maxpq.pq[maxpq.size] = val
	} else {
		maxpq.pq = append(maxpq.pq, val)
	}
	maxpq.swim(maxpq.size)
}

func (maxpq *MaxPQ[T]) Max() T {
	if maxpq.size < 1 {
		panic("algos.array.(MaxPQ).Max(): Queue is empty ")
	}
	return maxpq.pq[1]
}

func (maxpq *MaxPQ[T]) GetMax() T {
	if maxpq.size < 1 {
		panic("algos.array.(MaxPQ).GetMax(): Queue is empty ")
	}
	ret := maxpq.pq[1]
	swap(maxpq.pq, 1, maxpq.size)
	maxpq.size--
	maxpq.sink(1)
	return ret
}

func (maxpq *MaxPQ[T]) Iterate() []T {
	return maxpq.pq[1:(maxpq.size + 1)]
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
