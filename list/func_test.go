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

	t.Run("Map", func(t *testing.T) {
		n4 := &list.Node[int]{4, nil}
		n3 := &list.Node[int]{3, n4}
		n2 := &list.Node[int]{2, n3}
		n1 := &list.Node[int]{1, n2}
		root := &list.Node[int]{0, n1}
		bag := array.NewBag[int]()
		root1 := list.Map(root, func(x int) int { return x + 3 })
		list.Map(root1, func(x int) int { bag.Add(x); return x })
		assert.Equal(bag.Iterate(), []int{3, 4, 5, 6, 7})
	})
	t.Run("MapA", func(t *testing.T) {
		n4 := &list.Node[int]{4, nil}
		n3 := &list.Node[int]{3, n4}
		n2 := &list.Node[int]{2, n3}
		n1 := &list.Node[int]{1, n2}
		root := &list.Node[int]{0, n1}
		bag := array.NewBag[int]()
		root1 := list.MapA(root, func(x int) int { return x + 3 })
		list.Map(root1, func(x int) int { bag.Add(x); return x })
		assert.Equal(bag.Iterate(), []int{3, 4, 5, 6, 7})
	})
	t.Run("Reduce", func(t *testing.T) {
		n4 := &list.Node[int]{4, nil}
		n3 := &list.Node[int]{3, n4}
		n2 := &list.Node[int]{2, n3}
		n1 := &list.Node[int]{1, n2}
		root := &list.Node[int]{0, n1}
		res := list.Reduce(root, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 31)
	})
	t.Run("ReduceR", func(t *testing.T) {
		n4 := &list.Node[int]{4, nil}
		n3 := &list.Node[int]{3, n4}
		n2 := &list.Node[int]{2, n3}
		n1 := &list.Node[int]{1, n2}
		root := &list.Node[int]{0, n1}
		res := list.Reduce(root, func(acc int, x int) int { return acc + x + 3 }, 6)
		assert.Equal(res, 31)
	})
	t.Run("Filter", func(t *testing.T) {
		n9 := &list.Node[int]{9, nil}
		n8 := &list.Node[int]{8, n9}
		n7 := &list.Node[int]{7, n8}
		n6 := &list.Node[int]{6, n7}
		n5 := &list.Node[int]{5, n6}
		n4 := &list.Node[int]{4, n5}
		n3 := &list.Node[int]{3, n4}
		n2 := &list.Node[int]{2, n3}
		n1 := &list.Node[int]{1, n2}
		root := &list.Node[int]{0, n1}
		bag1 := array.NewBag[int]()
		res1 := list.Filter(root, func(x int) bool { return (x % 2) == 0 })
		list.Map(res1, func(x int) int { bag1.Add(x); return x })
		assert.Equal(bag1.Iterate(), []int{0, 2, 4, 6, 8})
		bag2 := array.NewBag[int]()
		res2 := list.Filter(root, func(x int) bool { return (x % 2) != 0 })
		list.Map(res2, func(x int) int { bag2.Add(x); return x })
		assert.Equal(bag2.Iterate(), []int{1, 3, 5, 7, 9})
	})
}
