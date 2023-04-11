package list_test

import (
	"testing"

	"github.com/kselnaag/algos/array"
	"github.com/kselnaag/algos/list"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	n9 := &list.Snode[int]{9, nil}
	n8 := &list.Snode[int]{8, n9}
	n7 := &list.Snode[int]{7, n8}
	n6 := &list.Snode[int]{6, n7}
	n5 := &list.Snode[int]{5, n6}
	n4 := &list.Snode[int]{4, n5}
	n3 := &list.Snode[int]{3, n4}
	n2 := &list.Snode[int]{2, n3}
	n1 := &list.Snode[int]{1, n2}
	root := &list.Snode[int]{0, n1}

	t.Run("Reverse", func(t *testing.T) {
		root1 := list.Reverse(root)
		assert.Equal(root, list.Reverse(root1))
		assert.Equal((*list.Snode[int])(nil), list.Reverse[int](nil))
	})

	t.Run("ReverseRec", func(t *testing.T) {
		root1 := list.ReverseRec(root)
		assert.Equal(root, list.ReverseRec(root1))
		assert.Equal((*list.Snode[int])(nil), list.ReverseRec[int](nil))
	})

	t.Run("Map", func(t *testing.T) {
		bag := array.NewBag[int]()
		root1 := list.Map(root, func(x int) int { return x + 3 })
		list.Map(root1, func(x int) int { bag.Add(x); return x })
		assert.Equal(bag.Iterate(), []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
		root1 = list.Map(nil, func(x int) int { return x + 3 })
		assert.Equal(root1, (*list.Snode[int])(nil))
	})
	t.Run("MapA", func(t *testing.T) {
		bag := array.NewBag[int]()
		root1 := list.MapA(root, func(x int) int { return x + 3 })
		list.Map(root1, func(x int) int { bag.Add(x); return x })
		assert.Equal(bag.Iterate(), []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
		root1 = list.MapA(nil, func(x int) int { return x + 3 })
		assert.Equal(root1, (*list.Snode[int])(nil))
	})
	t.Run("Reduce", func(t *testing.T) {
		res := list.Reduce(root, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 81)
		res = list.Reduce(nil, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 6)
	})
	t.Run("ReduceR", func(t *testing.T) {
		res := list.ReduceR(root, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 81)
		res = list.ReduceR(nil, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 6)
	})
	t.Run("Filter", func(t *testing.T) {
		bag1 := array.NewBag[int]()
		res1 := list.Filter(root, func(x int) bool { return (x % 2) == 0 })
		list.Map(res1, func(x int) int { bag1.Add(x); return x })
		assert.Equal(bag1.Iterate(), []int{0, 2, 4, 6, 8})
		bag2 := array.NewBag[int]()
		res2 := list.Filter(root, func(x int) bool { return (x % 2) != 0 })
		list.Map(res2, func(x int) int { bag2.Add(x); return x })
		assert.Equal(bag2.Iterate(), []int{1, 3, 5, 7, 9})
		res2 = list.Filter(nil, func(x int) bool { return (x % 2) != 0 })
		assert.Equal(res2, (*list.Snode[int])(nil))
	})
}
