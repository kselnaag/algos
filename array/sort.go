package array

import (
	"math/rand"

	I "github.com/kselnaag/algos/types"
)

func Quick3Sort[T any](arr []T) {
	Shuffle(arr)
	q3sort(arr, 0, len(arr)-1)
}

func q3sort[T any](arr []T, lo, hi int) {
	if hi <= lo {
		return
	}
	l, i, v, g := lo, lo+1, arr[lo], hi
	for i <= g {
		if I.LT(arr[i], v) {
			swap(arr, l, i)
			l++
			i++
		} else if I.GT(arr[i], v) {
			swap(arr, i, g)
			g--
		} else {
			i++
		}
	}
	q3sort(arr, lo, l-1)
	q3sort(arr, g+1, hi)
}

func HeapSort[T any](arr []T) {
	alen := len(arr)
	for k := alen / 2; k > 0; k-- {
		sink(arr, k, alen-1)
	}
	for n := alen - 1; n > 1; {
		swap(arr, 1, n)
		n--
		sink(arr, 1, n)
	}
}

func sink[T any](arr []T, k, n int) {
	for (2 * k) <= n {
		j := 2 * k
		if (j < n) && I.LT(arr[j], arr[j+1]) {
			j++
		}
		if !I.LT(arr[k], arr[j]) {
			break
		}
		swap(arr, k, j)
		k = j
	}
}

func QuickSort[T any](arr []T) {
	Shuffle(arr)
	qsort(arr, 0, len(arr)-1)
}

func qsort[T any](arr []T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSort(arr)
		return
	}
	j := pivot(arr, lo, hi)
	qsort(arr, lo, j-1)
	qsort(arr, j+1, hi)
}

func pivot[T any](arr []T, lo, hi int) int {
	i, j, v := lo, hi, lo
	for {
		for ; I.LT(arr[i], arr[v]); i++ {
			if i == hi {
				break
			}
		}
		for ; I.LT(arr[v], arr[j]); j-- {
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

func MergeSort[T any](arr []T) {
	alen := len(arr)
	aux := make([]T, alen)
	copy(aux, arr)
	mrgsort(aux, arr, 0, alen-1)
}

func mrgsort[T any](src []T, dst []T, lo, hi int) {
	if hi <= lo {
		return
	}
	if (hi - lo + 1) <= 12 {
		InsertSort(src[lo : hi+1])
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}
	mid := lo + (hi-lo)/2
	mrgsort(dst, src, lo, mid)
	mrgsort(dst, src, mid+1, hi)
	merge(src, dst, lo, mid, hi)
}

func merge[T any](src []T, dst []T, lo, mid, hi int) {
	if !I.GT(src[mid], src[mid+1]) {
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
		if I.LT(src[j], src[i]) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func ReverseSort[T any](arr []T) {
	alen := len(arr)
	mid := alen / 2
	for i := 0; i < mid; i++ {
		min := i
		max := i
		for j := i + 1; j < (alen - i); j++ {
			if I.LT(arr[j], arr[min]) {
				min = j
			}
			if I.GT(arr[j], arr[max]) {
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

func SelectSort[T any](arr []T) {
	alen := len(arr)
	for i := 0; i < alen; i++ {
		min := i
		for j := i + 1; j < alen; j++ {
			if I.LT(arr[j], arr[min]) {
				min = j
			}
		}
		swap(arr, i, min)
	}
}

func ShellSort[T any](arr []T) {
	alen := len(arr)
	h := 1
	for h < (alen / 3) {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < alen; i++ {
			for j := i; j >= h; j -= h {
				if I.LT(arr[j], arr[j-h]) {
					swap(arr, j, j-h)
				} else {
					break
				}
			}
		}
		h = h / 3
	}
}

func InsertSort[T any](arr []T) {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		for j := i; j > 0; j-- {
			if I.LT(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

func BinarySearch[T any](arr []T, elem T) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (low + ((high - low) / 2))
		val := arr[mid]
		if I.EQ(val, elem) {
			return mid
		} else if I.GT(val, elem) {
			high = mid
		} else if I.LT(val, elem) {
			low = mid + 1
		}
	}
	return -1
}

func Shuffle[T any](arr []T) {
	alen := len(arr)
	for i := 0; i < alen; i++ {
		j := rand.Intn(i + 1)
		swap(arr, i, j)
	}
}

func Reverse[T any](arr []T) {
	alen := len(arr)
	mid := alen / 2
	for i := 0; i < mid; i++ {
		swap(arr, i, alen-i-1)
	}
}

func Contained[T any](arr []T, elem T) int {
	for i, el := range arr {
		if I.EQ(el, elem) {
			return i
		}
	}
	return -1
}

func IsSorted[T any](arr []T) bool {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		if I.LT(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
