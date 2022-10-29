package tree_test

import (
	"testing"

	"github.com/kselnaag/algos/tree"
	"github.com/stretchr/testify/assert"
)

func TestAAtree(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	aatree := tree.NewAAmap[string, int]()
	assert.Equal(aatree.IsEmpty(), true)
	assert.Equal(aatree.Size(), 0)
	assert.Panics(func() { aatree.Del("A") }, "algos.tree.(AAmap).Del(key): the code is not panic when tree is empty")

	arr := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"} // S E A R C H X M P L
	for i, el := range arr {
		aatree.Put(el, i)
	}
	assert.Equal(aatree.IsEmpty(), false)
	assert.Equal(aatree.Size(), 10)

	assert.Equal(aatree.IsKey("S"), true)
	assert.Equal(aatree.IsKey("E"), true)
	assert.Equal(aatree.IsKey("A"), true)
	assert.Equal(aatree.IsKey("R"), true)
	assert.Equal(aatree.IsKey("C"), true)
	assert.Equal(aatree.IsKey("H"), true)
	assert.Equal(aatree.IsKey("X"), true)
	assert.Equal(aatree.IsKey("M"), true)
	assert.Equal(aatree.IsKey("P"), true)
	assert.Equal(aatree.IsKey("L"), true)
	assert.Equal(aatree.IsKey("J"), false)

	assert.Equal(aatree.Get("S"), 0)
	assert.Equal(aatree.Get("E"), 12)
	assert.Equal(aatree.Get("A"), 8)
	assert.Equal(aatree.Get("R"), 3)
	assert.Equal(aatree.Get("C"), 4)
	assert.Equal(aatree.Get("H"), 5)
	assert.Equal(aatree.Get("X"), 7)
	assert.Equal(aatree.Get("M"), 9)
	assert.Equal(aatree.Get("P"), 10)
	assert.Equal(aatree.Get("L"), 11)
	assert.Panics(func() { aatree.Get("J") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")

	assert.Equal(aatree.IterateKeys(), []string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"})
	assert.Equal(aatree.IsEmpty(), false)
	assert.Equal(aatree.Size(), 10)

	aatree.PrintTreeCheck()
	aatree.Del("M")
	aatree.Del("R")
	aatree.Del("E")
	aatree.Del("C")
	aatree.Del("P")
	aatree.Put("P", 16)
	assert.Panics(func() { aatree.Get("M") }, "algos.tree.(RBmap).Get(key): the code is not panic when key is not found")
	assert.Panics(func() { aatree.Del("Z") }, "algos.tree.(RBmap).Del(key): the code is not panic when key is not found")

	assert.Equal(aatree.IterateKeys(), []string{"A", "H", "L", "P", "S", "X"})
	assert.Equal(aatree.IsEmpty(), false)
	assert.Equal(aatree.Size(), 6)
	aatree.PrintTreeCheck()
}

/*
func NewAAmap[K I.Ord, V any]() AAmap[K, V]
func (tm *AAmap[K, V]) Size() int
func (tm *AAmap[K, V]) IsEmpty() bool
func (tm *AAmap[K, V]) IsKey(key K) bool
func (tm *AAmap[K, V]) Get(key K) V
func (tm *AAmap[K, V]) Put(key K, val V)
func (tm *AAmap[K, V]) Del(key K)
func (tm *AAmap[K, V]) IterateKeys() []K
func (tm *AAmap[K, V]) PrintTreeCheck()

func (tm *AAmap[K, V]) put(node *AAnode[K, V], key K, val V) *AAnode[K, V]
func (tm *AAmap[K, V]) balance(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) skew(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) split(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) decrLvl(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) del(node *AAnode[K, V], key K) *AAnode[K, V]
func (tm *AAmap[K, V]) min(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) delmin(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) balanceDel(node *AAnode[K, V]) *AAnode[K, V]
func (tm *AAmap[K, V]) iteratekeys(node *AAnode[K, V], keysarr []K) []K
func (tm *AAmap[K, V]) printtree(node *AAnode[K, V], n int)
*/
