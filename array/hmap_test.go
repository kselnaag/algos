package array_test

import (
	"testing"

	"github.com/kselnaag/algos/array"
	"github.com/stretchr/testify/assert"
)

func TestHmap(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("Hmap", func(t *testing.T) {
		hmap := array.NewHmap[int, string]()
		assert.Equal(hmap.IsEmpty(), true)
		assert.Equal(hmap.Size(), 0)
		assert.Equal(hmap.IsKey(1), false)
		assert.Equal(hmap.IsKey(2), false)
		hmap.Set(1, "one")
		assert.Equal(hmap.IsKey(1), true)
		hmap.Set(1, "oneone")
		assert.Equal(hmap.IsKey(1), true)
		hmap.Set(2, "two")
		assert.Equal(hmap.IsKey(2), true)
		assert.Equal(hmap.IsEmpty(), false)
		assert.Equal(hmap.Size(), 2)
		assert.Equal(hmap.Get(1), "oneone")
		assert.Equal(hmap.Get(2), "two")
		assert.Panics(func() { hmap.Get(0) }, "The code is not panic")
		assert.Panics(func() { hmap.Get(3) }, "The code is not panic")
		assert.Panics(func() { hmap.Del(3) }, "The code is not panic")
		hmap.Del(1)
		assert.Equal(hmap.IsKey(1), false)
		assert.Equal(hmap.IsEmpty(), false)
		assert.Equal(hmap.Size(), 1)
		hmap.Set(3, "a")
		hmap.Set(4, "b")
		hmap.Set(5, "v")
		hmap.Set(3, "c")
		hmap.Set(6, "az")
		hmap.Set(7, "ad")
		hmap.Set(8, "ae")
		hmap.Set(9, "at")
		assert.Equal(hmap.IsEmpty(), false)
		assert.Equal(hmap.Size(), 8)
		assert.Equal(hmap.IterateKeys(), []int{2, 3, 4, 5, 6, 7, 8, 9})
		hmap.Drop()
		assert.Equal(hmap.IsEmpty(), true)
		assert.Equal(hmap.Size(), 0)
	})
}
