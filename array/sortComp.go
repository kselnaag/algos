package array

import I "github.com/kselnaag/algos/types"

func HeapSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	for k := alen / 2; k > 0; k-- {
		sinkComp(arr, k, alen-1)
	}
	for n := alen - 1; n > 1; {
		swap(arr, 1, n)
		n--
		sinkComp(arr, 1, n)
	}
}

func sinkComp[T I.Comp](arr []*T, k, n int) {
	for (2 * k) <= n {
		j := 2 * k
		if (j < n) && ltComp(arr[j], arr[j+1]) {
			j++
		}
		if !ltComp(arr[k], arr[j]) {
			break
		}
		swap(arr, k, j)
		k = j
	}
}

func QuickSortComp[T I.Comp](arr []*T) {
	Shuffle(arr)
	qsortComp(arr, 0, len(arr)-1)
}

func qsortComp[T I.Comp](arr []*T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSortComp(arr)
		return
	}
	j := pivotComp(arr, lo, hi)
	qsortComp(arr, lo, j-1)
	qsortComp(arr, j+1, hi)
}

func pivotComp[T I.Comp](arr []*T, lo, hi int) int {
	i, j, v := lo, hi, lo
	for {
		for ; ltComp(arr[i], arr[v]); i++ {
			if i == hi {
				break
			}
		}
		for ; ltComp(arr[v], arr[j]); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		swap(arr, i, j)
	}
	swap(arr, lo, j)
	return j
}

func MergeSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	aux := make([]*T, alen)
	copy(aux, arr)
	mrgsortComp(aux, arr, 0, alen-1)
}

func mrgsortComp[T I.Comp](src []*T, dst []*T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSortComp(src[lo : hi+1])
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}
	mid := lo + (hi-lo)/2
	mrgsortComp(dst, src, lo, mid)
	mrgsortComp(dst, src, mid+1, hi)
	mergeComp(src, dst, lo, mid, hi)
}

func mergeComp[T I.Comp](src []*T, dst []*T, lo, mid, hi int) {
	if !gtComp(src[mid], src[mid+1]) {
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
			continue
		}
		if j > hi {
			dst[k] = src[i]
			i++
			continue
		}
		if ltComp(src[j], src[i]) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func ReverseSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	mid := alen / 2
	for i := 0; i < mid; i++ {
		min := i
		max := i
		for j := i + 1; j < (alen - i); j++ {
			if ltComp(arr[j], arr[min]) {
				min = j
			}
			if gtComp(arr[j], arr[max]) {
				max = j
			}
		}
		if max == i {
			if min == (alen - i - 1) {
				swap(arr, max, min)
			} else {
				swap(arr, (alen - i - 1), max)
				swap(arr, i, min)
			}
		} else {
			swap(arr, i, min)
			swap(arr, (alen - i - 1), max)
		}
	}
}

func SelectSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	for i := 0; i < alen; i++ {
		min := i
		for j := i + 1; j < alen; j++ {
			if ltComp(arr[j], arr[min]) {
				min = j
			}
		}
		swap(arr, i, min)
	}
}

func ShellSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	h := 1
	for h < (alen / 3) {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < alen; i++ {
			for j := i; j >= h; j -= h {
				if ltComp(arr[j], arr[j-h]) {
					swap(arr, j, j-h)
				} else {
					break
				}
			}
		}
		h = h / 3
	}
}

func InsertSortComp[T I.Comp](arr []*T) {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		for j := i; j > 0; j-- {
			if ltComp(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

func BinarySearchComp[T I.Comp](arr []*T, elem *T) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (low + ((high - low) / 2))
		val := arr[mid]
		if eqComp(val, elem) {
			return mid
		} else if gtComp(val, elem) {
			high = mid
		} else if ltComp(val, elem) {
			low = mid + 1
		}
	}
	return -1
}

func ContainedComp[T I.Comp](arr []*T, elem *T) int {
	for i, el := range arr {
		if eqComp(el, elem) {
			return i
		}
	}
	return -1
}

func IsSortedComp[T I.Comp](arr []*T) bool {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		if ltComp(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}
