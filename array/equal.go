package array

import (
	"math/rand"

	I "github.com/kselnaag/algos/types"
)

func gtOrd[T I.Ord](i, j T) bool {
	return i > j
}

func ltOrd[T I.Ord](i, j T) bool {
	return i < j
}

func eqOrd[T I.Ord](i, j T) bool {
	return i == j
}

func gtComp[T I.Comp](i, j *T) bool {
	return (*i).CompareTo(*j) > 0
}

func ltComp[T I.Comp](i, j *T) bool {
	return (*i).CompareTo(*j) < 0
}

func eqComp[T I.Comp](i, j *T) bool {
	return (*i).CompareTo(*j) == 0
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

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
func median3Ord[T I.Ord](arr []T, i, j, k int) int {
	if ltOrd(arr, i, j) {
		if ltOrd(arr, j, k) {
			return j
		} else {
			if ltOrd(arr, i, k) {
				return k
			} else {
				return i
			}
		}
	} else {
		if ltOrd(arr, k, j) {
			return j
		} else {
			if ltOrd(arr, k, i) {
				return k
			} else {
				return i
			}
		}
	}
}
*/
