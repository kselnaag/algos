package list_test

import (
	"testing"

	"github.com/kselnaag/algos/list"

	"github.com/stretchr/testify/assert"
)

func TestHmap(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	hmap := list.NewHmap[int, string]()
	asrt.True(hmap.IsEmpty())
	asrt.Equal(0, hmap.Size())
	asrt.False(hmap.IsKey(1))
	asrt.False(hmap.IsKey(2))
	hmap.Set(1, "one")
	asrt.True(hmap.IsKey(1))
	hmap.Set(1, "oneone")
	asrt.True(hmap.IsKey(1))
	hmap.Set(2, "two")
	asrt.True(hmap.IsKey(2))
	asrt.False(hmap.IsEmpty())
	asrt.Equal(2, hmap.Size())
	asrt.Equal("oneone", hmap.Get(1))
	asrt.Equal("two", hmap.Get(2))
	asrt.Panics(func() { hmap.Get(0) }, "algos.list.(Hmap).Get():  the code is not panic when key is not found")
	asrt.Panics(func() { hmap.Get(3) }, "algos.list.(Hmap).Get():  the code is not panic when key is not found")
	asrt.Panics(func() { hmap.Del(3) }, "algos.list.(Hmap).Del():  the code is not panic when key is not found")
	hmap.Del(1)
	asrt.False(hmap.IsKey(1))
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
}
