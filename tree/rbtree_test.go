package tree_test

import (
	"testing"

	"github.com/kselnaag/algos/tree"

	"github.com/stretchr/testify/assert"
)

func TestRBtree(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	rbtree := tree.NewRBmap[string, int]()
	assert.Equal(rbtree.IsEmpty(), true)
	assert.Equal(rbtree.Size(), 0)
	assert.Panics(func() { rbtree.Del("A") }, "algos.tree.(RBmap).Del(key): the code is not panic when tree is empty")

	arr := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"} // S E A R C H X M P L
	for i, el := range arr {
		rbtree.Put(el, i)
	}
	assert.Equal(rbtree.IsEmpty(), false)
	assert.Equal(rbtree.Size(), 10)

	assert.Equal(rbtree.IsKey("S"), true)
	assert.Equal(rbtree.IsKey("E"), true)
	assert.Equal(rbtree.IsKey("A"), true)
	assert.Equal(rbtree.IsKey("R"), true)
	assert.Equal(rbtree.IsKey("C"), true)
	assert.Equal(rbtree.IsKey("H"), true)
	assert.Equal(rbtree.IsKey("X"), true)
	assert.Equal(rbtree.IsKey("M"), true)
	assert.Equal(rbtree.IsKey("P"), true)
	assert.Equal(rbtree.IsKey("L"), true)
	assert.Equal(rbtree.IsKey("J"), false)

	assert.Equal(rbtree.Get("S"), 0)
	assert.Equal(rbtree.Get("E"), 12)
	assert.Equal(rbtree.Get("A"), 8)
	assert.Equal(rbtree.Get("R"), 3)
	assert.Equal(rbtree.Get("C"), 4)
	assert.Equal(rbtree.Get("H"), 5)
	assert.Equal(rbtree.Get("X"), 7)
	assert.Equal(rbtree.Get("M"), 9)
	assert.Equal(rbtree.Get("P"), 10)
	assert.Equal(rbtree.Get("L"), 11)
	assert.Panics(func() { rbtree.Get("J") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")

	assert.Equal(rbtree.IsBSTcheck(), true)
	assert.Equal(rbtree.BSTheightCheck(), 3)
	assert.Equal(rbtree.IsRedBlackCheck(), true)
	assert.Equal(rbtree.IsBalancedCheck(), true)

	assert.Equal(rbtree.IterateKeys(), []string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"})
	assert.Equal(rbtree.IsEmpty(), false)
	assert.Equal(rbtree.Size(), 10)

	rbtree.Del("R")
	rbtree.Del("S")
	rbtree.Del("X")
	rbtree.Del("P")
	rbtree.Put("P", 16)
	assert.Panics(func() { rbtree.Get("X") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")
	assert.Panics(func() { rbtree.Del("Z") }, "algos.tree.(RBmap).Del(key): the code is not panic when key is not found")

	assert.Equal(rbtree.IterateKeys(), []string{"A", "C", "E", "H", "L", "M", "P"})
	assert.Equal(rbtree.IsEmpty(), false)
	assert.Equal(rbtree.Size(), 7)

	assert.Equal(rbtree.IsBSTcheck(), true)
	assert.Equal(rbtree.BSTheightCheck(), 3)
	assert.Equal(rbtree.IsRedBlackCheck(), true)
	assert.Equal(rbtree.IsBalancedCheck(), true)

	rbtree.PrintTreeCheck()
}

/*
NewRBmap Size IsEmpty IsKey Get Put
Del IterateKeys

put iteratekeys printtree isbalanced isredblack bstheight isbst
balance rotateLeft rotateRight flipColors

IsBSTcheck IsRedBlackCheck IsBalancedCheck
BSTheightCheck PrintTreeCheck
*/
