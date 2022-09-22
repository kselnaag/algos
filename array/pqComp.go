package array

import (
	I "github.com/kselnaag/algos/types"
)

// ===========================
type MinPQComp[T I.Comp] struct {
	pq   []*T
	size int
}

func NewMinPQComp[T I.Comp]() MinPQComp[T] {
	return MinPQComp[T]{
		pq:   make([]*T, 1, 8),
		size: 0,
	}
}

func (min *MinPQComp[T]) IsEmpty() bool {
	return min.size == 0
}

func (min *MinPQComp[T]) Size() int {
	return min.size
}

func (min *MinPQComp[T]) swim(k int) {
	for (k > 1) && gtComp(min.pq[k/2], min.pq[k]) {
		swap(min.pq, k/2, k)
		k /= 2
	}
}
func (min *MinPQComp[T]) sink(k int) {
	for (2 * k) <= min.size {
		j := 2 * k
		if (j < min.size) && gtComp(min.pq[j], min.pq[j+1]) {
			j++
		}
		if !gtComp(min.pq[k], min.pq[j]) {
			break
		}
		swap(min.pq, k, j)
		k = j
	}
}

func (min *MinPQComp[T]) Add(val *T) {
	min.size++
	if len(min.pq) > min.size {
		min.pq[min.size] = val
	} else {
		min.pq = append(min.pq, val)
	}
	min.swim(min.size)
}

func (min *MinPQComp[T]) Min() *T {
	if min.size < 1 {
		panic("algos.array.(MinPQ).Min(): Queue is empty ")
	}
	return min.pq[1]
}

func (min *MinPQComp[T]) GetMin() *T {
	if min.size < 1 {
		panic("algos.array.(MinPQ).GetMin(): Queue is empty ")
	}
	ret := min.pq[1]
	swap(min.pq, 1, min.size)
	min.size--
	min.sink(1)
	return ret
}

// ===========================
type MaxPQComp[T I.Comp] struct {
	pq   []*T
	size int
}

func NewMaxPQComp[T I.Comp]() MaxPQComp[T] {
	return MaxPQComp[T]{
		pq:   make([]*T, 1, 8),
		size: 0,
	}
}

func (max *MaxPQComp[T]) IsEmpty() bool {
	return max.size == 0
}

func (max *MaxPQComp[T]) Size() int {
	return max.size
}

func (max *MaxPQComp[T]) swim(k int) {
	for (k > 1) && ltComp(max.pq[k/2], max.pq[k]) {
		swap(max.pq, k/2, k)
		k /= 2
	}
}
func (max *MaxPQComp[T]) sink(k int) {
	for (2 * k) <= max.size {
		j := 2 * k
		if (j < max.size) && ltComp(max.pq[j], max.pq[j+1]) {
			j++
		}
		if !ltComp(max.pq[k], max.pq[j]) {
			break
		}
		swap(max.pq, k, j)
		k = j
	}
}

func (max *MaxPQComp[T]) Add(val *T) {
	max.size++
	if len(max.pq) > max.size {
		max.pq[max.size] = val
	} else {
		max.pq = append(max.pq, val)
	}
	max.swim(max.size)
}

func (max *MaxPQComp[T]) Max() *T {
	if max.size < 1 {
		panic("algos.array.(MaxPQ).Max(): Queue is empty ")
	}
	return max.pq[1]
}

func (max *MaxPQComp[T]) GetMax() *T {
	if max.size < 1 {
		panic("algos.array.(MaxPQ).GetMax(): Queue is empty ")
	}
	ret := max.pq[1]
	swap(max.pq, 1, max.size)
	max.size--
	max.sink(1)
	return ret
}

/*
//===========================//
type MinIdxPQComp[T I.Comp] struct {
	pq   []int
	qp   []int
	keys []*T
	size int
}

func NewMinIdxPQComp[T I.Comp]() MinIdxPQComp[T] {
	alen := 1
	acap := 8
	return MinIdxPQComp[T]{
		pq:   make([]int, alen, acap),
		qp:   make([]int, alen, acap),
		keys: make([]*T, alen, acap),
		size: 0,
	}
}

func (min *MinIdxPQComp[T]) gt(i, j int) bool {
	return min.keys[min.pq[i]] > min.keys[min.pq[j]]
}

func (min *MinIdxPQComp[T]) swap(i, j int) {
	min.pq[i], min.pq[j] = min.pq[j], min.pq[i]
	min.qp[min.pq[i]], min.qp[min.pq[j]] = i, j
}

func (min *MinIdxPQComp[T]) swim(k int) {
	for (k > 1) && min.gt(k/2, k) {
		min.swap(k/2, k)
		k /= 2
	}
}

func (min *MinIdxPQComp[T]) sink(k int) {
	for (2 * k) <= min.size {
		j := 2 * k
		if (j < min.size) && min.gt(j, j+1) {
			j++
		}
		if !min.gt(k, j) {
			break
		}
		min.swap(k, j)
		k = j
	}
}

func (min *MinIdxPQComp[T]) IsEmpty() bool {
	return min.size == 0
}

func (min *MinIdxPQComp[T]) Size() int {
	return min.size
}

func (min *MinIdxPQComp[T]) Contains(k int) bool {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).Contains(k): Index 'k' is out of range")
	}
	return min.qp[k] != 0
}

func (min *MinIdxPQComp[T]) InsertVal(k int, v T) {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).InsertVal(k, v): Index 'k' is out of range ")
	}
	if min.Contains(k) {
		panic("algos.array.(MinIdxPQ).InsertVal(k, v): Index 'k' is already exist ")
	}
	min.size++
	if len(min.pq) > min.size {
		min.qp[k] = min.size
		min.pq[min.size] = k
		min.keys[k] = v

	} else {
		min.qp = append(min.qp, min.size)
		min.pq = append(min.pq, k)
		min.keys = append(min.keys, v)
	}
	min.swim(min.size)
}

func (min *MinIdxPQComp[T]) MinIndex() int {
	if min.size < 1 {
		panic("algos.array.(MinIdxPQ).MinIndex(): No Index to return, zero length queqe")
	}
	return min.pq[1]
}

func (min *MinIdxPQComp[T]) MinVal() T {
	if min.size < 1 {
		panic("algos.array.(MinIdxPQ).MinVal(): No Val to return, zero length queqe")
	}
	return min.keys[min.pq[1]]
}

func (min *MinIdxPQComp[T]) DelMinIndex() int {
	if min.size < 1 {
		panic("algos.array.(MinIdxPQ).DelMinIndex(): No val to return, zero lenght queqe")
	}
	ret := min.pq[1]
	min.swap(1, min.size)
	min.size--
	min.sink(1)
	min.qp[ret] = 0
	min.pq[min.size+1] = 0
	return ret
}

func (min *MinIdxPQComp[T]) Val(k int) T {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).Val(k): Index 'k' is out of range ")
	}
	if !min.Contains(k) {
		panic("algos.array.(MinIdxPQ).Val(k): Index 'k' is not in queue ")
	}
	return min.keys[k]
}

func (min *MinIdxPQComp[T]) ChangeVal(k int, v T) {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).ChangeVal(k, v): Index 'k' is out of range ")
	}
	if !min.Contains(k) {
		panic("algos.array.(MinIdxPQ).ChangeVal(k, v): Index 'k' is not in queue ")
	}
	min.keys[k] = v
	min.swim(min.qp[k])
	min.sink(min.qp[k])
}

func (min *MinIdxPQComp[T]) IncreaseIndex(k int, v T) {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).IncreaseIndex(k, v): Index 'k' is out of range ")
	}
	if !min.Contains(k) {
		panic("algos.array.(MinIdxPQ).IncreaseIndex(k, v): Index 'k' is not in queue ")
	}
	if min.keys[k] >= v {
		panic("algos.array.(MinIdxPQ).IncreaseIndex(k, v): With such index 'k' and value 'v' queue would not change anything ")
	}
	min.keys[k] = v
	min.sink(min.qp[k])
}

func (min *MinIdxPQComp[T]) DecreaseIndex(k int, v T) {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).DecreaseIndex(k, v): Index 'k' is out of range ")
	}
	if !min.Contains(k) {
		panic("algos.array.(MinIdxPQ).DecreaseIndex(k, v): Index 'k' is not in queue ")
	}
	if min.keys[k] <= v {
		panic("algos.array.(MinIdxPQ).DecreaseIndex(k, v): With such index 'k' and value 'v' queue would not change anything ")
	}
	min.keys[k] = v
	min.swim(min.qp[k])
}

func (min *MinIdxPQComp[T]) DeleteIndex(k int) {
	if (k < 1) || (k > min.size) {
		panic("algos.array.(MinIdxPQ).DeleteIndex(k): Index 'k' is out of range ")
	}
	if !min.Contains(k) {
		panic("algos.array.(MinIdxPQ).DeleteIndex(k): Index 'k' is not in queue ")
	}
	index := min.qp[k]
	min.swap(index, min.size)
	min.size--
	min.swim(index)
	min.sink(index)
	min.qp[k] = 0
}

// ===========================
type MaxIdxPQComp[T I.Comp] struct {
	pq   []*T
	size int
}

func NewMaxIdxPQComp[T I.Comp]() MaxIdxPQComp[T] {
	return MaxIdxPQComp[T]{
		pq:   make([]*T, 1, 8),
		size: 0,
	}
}

// IsEmpty
// Size
// Insert
// Change
// Contains
// Delete
// MaxVal
// DelMaxIndex
// MaxIndex

*/
