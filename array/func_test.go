package array_test

import (
	"testing"

	"algos/array"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("Map", func(t *testing.T) {
		arr := array.Map([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) int { return x + 3 })
		assert.Equal(arr, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
		arr = array.Map([]int{}, func(x int) int { return x + 3 })
		assert.Equal(arr, []int{})
	})
	t.Run("MapA", func(t *testing.T) {
		arr := array.MapA([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) int { return x + 3 })
		assert.Equal(arr, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
		arr = array.MapA([]int{}, func(x int) int { return x + 3 })
		assert.Equal(arr, []int{})
	})
	t.Run("Reduce", func(t *testing.T) {
		res := array.Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 91)
		res = array.Reduce([]int{}, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 6)
	})
	t.Run("ReduceR", func(t *testing.T) {
		res := array.ReduceR([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 91)
		res = array.ReduceR([]int{}, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 6)
	})
	t.Run("Filter", func(t *testing.T) {
		arr := array.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return (x % 2) != 0 })
		assert.Equal(arr, []int{1, 3, 5, 7, 9})
		arr = array.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return (x % 2) == 0 })
		assert.Equal(arr, []int{2, 4, 6, 8, 10})
		arr = array.Filter([]int{}, func(x int) bool { return (x % 2) == 0 })
		assert.Equal(arr, []int{})
	})
}
