package list_test

import (
	"testing"

	"algos/list"

	"github.com/stretchr/testify/assert"
)

func TestHmap(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	hmap := list.NewHmap[int, string]()
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
	assert.Panics(func() { hmap.Get(0) }, "algos.list.(Hmap).Get():  the code is not panic when key is not found")
	assert.Panics(func() { hmap.Get(3) }, "algos.list.(Hmap).Get():  the code is not panic when key is not found")
	assert.Panics(func() { hmap.Del(3) }, "algos.list.(Hmap).Del():  the code is not panic when key is not found")
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
	assert.Equal(hmap.IterateKeys(), []int{5, 6, 9, 4, 8, 2, 3, 7})
}
