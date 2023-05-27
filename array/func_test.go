package array_test

import (
	"strconv"
	"testing"

	"github.com/kselnaag/algos/array"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("Map", func(t *testing.T) {
		arr := array.Map([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) int { return x + 3 })
		asrt.Equal([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, arr)
		arr = array.Map([]int{}, func(x int) int { return x + 3 })
		asrt.Equal([]int{}, arr)
	})
	t.Run("MapA", func(t *testing.T) {
		arr := array.MapA([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) int { return x + 3 })
		asrt.Equal([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, arr)
		arr = array.MapA([]int{}, func(x int) int { return x + 3 })
		asrt.Equal([]int{}, arr)
	})
	t.Run("Reduce", func(t *testing.T) {
		res := array.Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc int, x int) int { return acc + x + 3 }, 6)
		asrt.Equal(91, res)
		res = array.Reduce([]int{}, func(acc int, x int) int { return acc + x + 3 }, 6)
		asrt.Equal(6, res)
	})
	t.Run("ReduceR", func(t *testing.T) {
		res := array.ReduceR([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc int, x int) int { return acc + x + 3 }, 6)
		asrt.Equal(91, res)
		res = array.ReduceR([]int{}, func(acc int, x int) int { return acc + x + 3 }, 6)
		asrt.Equal(6, res)
	})
	t.Run("Filter", func(t *testing.T) {
		arr := array.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return (x % 2) != 0 })
		asrt.Equal([]int{1, 3, 5, 7, 9}, arr)
		arr = array.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return (x % 2) == 0 })
		asrt.Equal([]int{2, 4, 6, 8, 10}, arr)
		arr = array.Filter([]int{}, func(x int) bool { return (x % 2) == 0 })
		asrt.Equal([]int{}, arr)
	})
	t.Run("Perm", func(t *testing.T) {
		resStr := make([]string, 0, 6)
		array.Perm([]byte("abc"), func(a []byte) {
			resStr = append(resStr, string(a))
		})
		asrt.Equal([]string{"abc", "acb", "bac", "bca", "cba", "cab"}, resStr)

		resInt := make([]string, 0, 6)
		array.Perm([]int{1, 3, 2}, func(a []int) {
			str := ""
			for _, el := range a {
				str += strconv.Itoa(el) + " "
			}
			resInt = append(resInt, str)
		})
		asrt.Equal([]string{"1 3 2 ", "1 2 3 ", "3 1 2 ", "3 2 1 ", "2 3 1 ", "2 1 3 "}, resInt)
	})
}
