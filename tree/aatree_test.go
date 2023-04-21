package tree_test

import (
	"testing"

	"github.com/kselnaag/algos/tree"

	"github.com/stretchr/testify/assert"
)

func TestAAtree(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	aatree := tree.NewAAmap[string, int]()
	asrt.True(aatree.IsEmpty())
	asrt.Equal(0, aatree.Size())
	asrt.Panics(func() { aatree.Del("A") }, "algos.tree.(AAmap).Del(key): the code is not panic when tree is empty")

	arr := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"} // S E A R C H X M P L
	for i, el := range arr {
		aatree.Put(el, i)
	}
	asrt.False(aatree.IsEmpty())
	asrt.Equal(10, aatree.Size())

	asrt.True(aatree.IsKey("S"))
	asrt.True(aatree.IsKey("E"))
	asrt.True(aatree.IsKey("A"))
	asrt.True(aatree.IsKey("R"))
	asrt.True(aatree.IsKey("C"))
	asrt.True(aatree.IsKey("H"))
	asrt.True(aatree.IsKey("X"))
	asrt.True(aatree.IsKey("M"))
	asrt.True(aatree.IsKey("P"))
	asrt.True(aatree.IsKey("L"))
	asrt.False(aatree.IsKey("J"))

	asrt.Equal(0, aatree.Get("S"))
	asrt.Equal(12, aatree.Get("E"))
	asrt.Equal(8, aatree.Get("A"))
	asrt.Equal(3, aatree.Get("R"))
	asrt.Equal(4, aatree.Get("C"))
	asrt.Equal(5, aatree.Get("H"))
	asrt.Equal(7, aatree.Get("X"))
	asrt.Equal(9, aatree.Get("M"))
	asrt.Equal(10, aatree.Get("P"))
	asrt.Equal(11, aatree.Get("L"))
	asrt.Panics(func() { aatree.Get("J") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")

	asrt.Equal([]string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"}, aatree.IterateKeys())
	asrt.False(aatree.IsEmpty())
	asrt.Equal(10, aatree.Size())

	aatree.PrintTreeCheck()
	aatree.Del("M")
	aatree.Del("R")
	aatree.Del("E")
	aatree.Del("C")
	aatree.Del("P")
	aatree.Put("P", 16)
	asrt.Panics(func() { aatree.Get("M") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")
	asrt.Panics(func() { aatree.Del("Z") }, "algos.tree.(RBmap).Del(key): the code is not panic when key is not found")

	asrt.Equal([]string{"A", "H", "L", "P", "S", "X"}, aatree.IterateKeys())
	asrt.False(aatree.IsEmpty())
	asrt.Equal(6, aatree.Size())
	aatree.PrintTreeCheck()
}
