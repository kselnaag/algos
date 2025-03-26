package array_test

import (
	"strconv"
	"testing"

	"algos/array"

	"github.com/stretchr/testify/assert"
)

func TestHmap(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("Hmap", func(t *testing.T) {
		hmap := array.NewHMap[int, string]()
		asrt.True(hmap.IsEmpty())
		asrt.Equal(0, hmap.Size())
		asrt.Nil(hmap.Get(1))
		asrt.Nil(hmap.Get(2))
		hmap.Set(1, "one")
		asrt.NotNil(hmap.Get(1))
		hmap.Set(1, "oneone")
		asrt.NotNil(hmap.Get(1))
		hmap.Set(2, "two")
		asrt.NotNil(hmap.Get(2))
		asrt.False(hmap.IsEmpty())
		asrt.Equal(2, hmap.Size())
		asrt.Equal("oneone", *hmap.Get(1))
		asrt.Equal("two", *hmap.Get(2))
		hmap.Del(1)
		asrt.Nil(hmap.Get(1))
		asrt.False(hmap.IsEmpty())
		asrt.Equal(1, hmap.Size())
		hmap.Set(3, "a")
		hmap.Set(4, "b")
		hmap.Set(5, "v")
		hmap.Set(3, "c")
		hmap.Set(6, "az")
		hmap.Set(7, "ad")
		hmap.Set(8, "ae")
		hmap.Set(9, "at")
		asrt.False(hmap.IsEmpty())
		asrt.Equal(8, hmap.Size())
		asrt.Equal([]int{5, 4, 7, 6, 3, 2, 9, 8}, hmap.IterateKeys())

		for i := 0; i < 101; i++ {
			hmap.Set(i, strconv.Itoa(i))
		}
		asrt.Equal(400, hmap.Buckets())
		hmap.PrintAll()
	})
}
