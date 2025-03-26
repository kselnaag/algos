package tree_test

import (
	"testing"

	"algos/tree"

	"github.com/stretchr/testify/assert"
)

func TestRBtree(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	rbtree := tree.NewRBmap[string, int]()
	asrt.True(rbtree.IsEmpty())
	asrt.Equal(0, rbtree.Size())
	asrt.Panics(func() { rbtree.Del("A") }, "algos.tree.(RBmap).Del(key): the code is not panic when tree is empty")

	arr := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"} // S E A R C H X M P L
	for i, el := range arr {
		rbtree.Put(el, i)
	}
	asrt.False(rbtree.IsEmpty())
	asrt.Equal(10, rbtree.Size())

	asrt.True(rbtree.IsKey("S"))
	asrt.True(rbtree.IsKey("E"))
	asrt.True(rbtree.IsKey("A"))
	asrt.True(rbtree.IsKey("R"))
	asrt.True(rbtree.IsKey("C"))
	asrt.True(rbtree.IsKey("H"))
	asrt.True(rbtree.IsKey("X"))
	asrt.True(rbtree.IsKey("M"))
	asrt.True(rbtree.IsKey("P"))
	asrt.True(rbtree.IsKey("L"))
	asrt.False(rbtree.IsKey("J"))

	asrt.Equal(0, rbtree.Get("S"))
	asrt.Equal(12, rbtree.Get("E"))
	asrt.Equal(8, rbtree.Get("A"))
	asrt.Equal(3, rbtree.Get("R"))
	asrt.Equal(4, rbtree.Get("C"))
	asrt.Equal(5, rbtree.Get("H"))
	asrt.Equal(7, rbtree.Get("X"))
	asrt.Equal(9, rbtree.Get("M"))
	asrt.Equal(10, rbtree.Get("P"))
	asrt.Equal(11, rbtree.Get("L"))
	asrt.Panics(func() { rbtree.Get("J") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")

	asrt.True(rbtree.IsBSTcheck())
	asrt.Equal(3, rbtree.BSTheightCheck())
	asrt.True(rbtree.IsRedBlackCheck())
	asrt.True(rbtree.IsBalancedCheck())

	asrt.Equal([]string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"}, rbtree.IterateKeys())
	asrt.False(rbtree.IsEmpty())
	asrt.Equal(10, rbtree.Size())

	rbtree.Del("R")
	rbtree.Del("S")
	rbtree.Del("X")
	rbtree.Del("P")
	rbtree.Put("P", 16)
	asrt.Panics(func() { rbtree.Get("X") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")
	asrt.Panics(func() { rbtree.Del("Z") }, "algos.tree.(RBmap).Del(key): the code is not panic when key is not found")

	asrt.Equal([]string{"A", "C", "E", "H", "L", "M", "P"}, rbtree.IterateKeys())
	asrt.False(rbtree.IsEmpty())
	asrt.Equal(7, rbtree.Size())

	asrt.True(rbtree.IsBSTcheck())
	asrt.Equal(3, rbtree.BSTheightCheck())
	asrt.True(rbtree.IsRedBlackCheck())
	asrt.True(rbtree.IsBalancedCheck())

	rbtree.PrintTreeCheck()
}
