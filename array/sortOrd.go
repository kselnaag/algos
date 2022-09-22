package array

import I "github.com/kselnaag/algos/types"

func HeapSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	for k := alen / 2; k > 0; k-- {
		sinkOrd(arr, k, alen-1)
	}
	for n := alen - 1; n > 1; {
		swap(arr, 1, n)
		n--
		sinkOrd(arr, 1, n)
	}
}

func sinkOrd[T I.Ord](arr []T, k, n int) {
	for (2 * k) <= n {
		j := 2 * k
		if (j < n) && ltOrd(arr[j], arr[j+1]) {
			j++
		}
		if !ltOrd(arr[k], arr[j]) {
			break
		}
		swap(arr, k, j)
		k = j
	}
}

func QuickSortOrd[T I.Ord](arr []T) {
	Shuffle(arr)
	qsortOrd(arr, 0, len(arr)-1)
}

func qsortOrd[T I.Ord](arr []T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSortOrd(arr)
		return
	}
	j := pivotOrd(arr, lo, hi)
	qsortOrd(arr, lo, j-1)
	qsortOrd(arr, j+1, hi)
}

func pivotOrd[T I.Ord](arr []T, lo, hi int) int {
	i, j, v := lo, hi, lo
	for {
		for ; ltOrd(arr[i], arr[v]); i++ {
			if i == hi {
				break
			}
		}
		for ; ltOrd(arr[v], arr[j]); j-- {
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

func MergeSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	aux := make([]T, alen)
	copy(aux, arr)
	mrgsortOrd(aux, arr, 0, alen-1)
}

func mrgsortOrd[T I.Ord](src []T, dst []T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSortOrd(src[lo : hi+1])
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}
	mid := lo + (hi-lo)/2
	mrgsortOrd(dst, src, lo, mid)
	mrgsortOrd(dst, src, mid+1, hi)
	mergeOrd(src, dst, lo, mid, hi)
}

func mergeOrd[T I.Ord](src []T, dst []T, lo, mid, hi int) {
	if !gtOrd(src[mid], src[mid+1]) {
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
		if ltOrd(src[j], src[i]) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func ReverseSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	mid := alen / 2
	for i := 0; i < mid; i++ {
		min := i
		max := i
		for j := i + 1; j < (alen - i); j++ {
			if ltOrd(arr[j], arr[min]) {
				min = j
			}
			if gtOrd(arr[j], arr[max]) {
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

func SelectSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	for i := 0; i < alen; i++ {
		min := i
		for j := i + 1; j < alen; j++ {
			if ltOrd(arr[j], arr[min]) {
				min = j
			}
		}
		swap(arr, i, min)
	}
}

func ShellSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	h := 1
	for h < (alen / 3) {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < alen; i++ {
			for j := i; j >= h; j -= h {
				if ltOrd(arr[j], arr[j-h]) {
					swap(arr, j, j-h)
				} else {
					break
				}
			}
		}
		h = h / 3
	}
}

func InsertSortOrd[T I.Ord](arr []T) {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		for j := i; j > 0; j-- {
			if ltOrd(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

func BinarySearchOrd[T I.Ord](arr []T, elem T) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (low + ((high - low) / 2))
		val := arr[mid]
		if eqOrd(val, elem) {
			return mid
		} else if gtOrd(val, elem) {
			high = mid
		} else if ltOrd(val, elem) {
			low = mid + 1
		}
	}
	return -1
}

func ContainedOrd[T I.Ord](arr []T, elem T) int {
	for i, el := range arr {
		if eqOrd(el, elem) {
			return i
		}
	}
	return -1
}

func IsSortedOrd[T I.Ord](arr []T) bool {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		if ltOrd(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}
