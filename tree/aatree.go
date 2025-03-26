package tree

import (
	"fmt"

	"algos/amath"
	I "algos/types"
)

// right-linear Arne Andersson tree with intensive node deletion
type AAnode[K I.Ord, V any] struct {
	Key  K
	Val  V
	L, R *AAnode[K, V]
	Lvl  byte
}

type AAmap[K I.Ord, V any] struct {
	root *AAnode[K, V]
	size int
}

func NewAAmap[K I.Ord, V any]() AAmap[K, V] {
	return AAmap[K, V]{
		root: nil,
		size: 0,
	}
}

func (tm *AAmap[K, V]) Size() int {
	return tm.size
}

func (tm *AAmap[K, V]) IsEmpty() bool {
	return tm.Size() == 0
}

func (tm *AAmap[K, V]) IsKey(key K) bool {
	node := tm.root
	for node != nil {
		switch {
		case I.LT(key, node.Key):
			node = node.L
		case I.GT(key, node.Key):
			node = node.R
		default:
			return true
		}
	}
	return false
}

func (tm *AAmap[K, V]) Get(key K) V {
	node := tm.root
	for node != nil {
		switch {
		case I.LT(key, node.Key):
			node = node.L
		case I.GT(key, node.Key):
			node = node.R
		default:
			return node.Val
		}
	}
	panic("algos.tree.(AAmap).Get(key): No any key found, check first")
}

func (tm *AAmap[K, V]) Put(key K, val V) {
	tm.root = tm.put(tm.root, key, val)
}

func (tm *AAmap[K, V]) put(node *AAnode[K, V], key K, val V) *AAnode[K, V] {
	if node == nil {
		tm.size++
		node = &AAnode[K, V]{Key: key, Val: val, L: nil, R: nil, Lvl: 1}
	} else {
		switch {
		case I.LT(key, node.Key):
			node.L = tm.put(node.L, key, val)
		case I.GT(key, node.Key):
			node.R = tm.put(node.R, key, val)
		default:
			node.Val = val
		}
	}
	return tm.balance(node)
}

func (tm *AAmap[K, V]) balance(node *AAnode[K, V]) *AAnode[K, V] {
	if node == nil {
		return nil
	}
	node = tm.skew(node)
	node = tm.split(node)
	return node
}

func (tm *AAmap[K, V]) skew(node *AAnode[K, V]) *AAnode[K, V] { // rightRotation
	if node == nil {
		return nil
	}
	if (node.L != nil) && (node.L.Lvl == node.Lvl) {
		x := node.L
		node.L = x.R
		x.R = node
		return x
	}
	return node
}

func (tm *AAmap[K, V]) split(node *AAnode[K, V]) *AAnode[K, V] { // leftRotation
	if node == nil {
		return nil
	}
	if (node.R != nil) && (node.R.R != nil) && (node.Lvl == node.R.R.Lvl) {
		x := node.R
		node.R = x.L
		x.L = node
		x.Lvl++
		return x
	}
	return node
}

func (tm *AAmap[K, V]) Del(key K) {
	tm.root = tm.del(tm.root, key)
}

func (tm *AAmap[K, V]) del(node *AAnode[K, V], key K) *AAnode[K, V] {
	if node == nil {
		panic("algos.tree.(AAmap).Del(key): No any key found, check first")
	}
	switch {
	case I.LT(key, node.Key):
		node.L = tm.del(node.L, key)
	case I.GT(key, node.Key):
		node.R = tm.del(node.R, key)
	default:
		if tm.size > 0 {
			tm.size--
		}
		if node.R == nil {
			return node.L
		}
		minnode := tm.min(node.R)
		node.Key = minnode.Key
		node.Val = minnode.Val
		node.R = tm.delmin(node.R)
	}
	return tm.balanceDel(node)
}

func (tm *AAmap[K, V]) min(node *AAnode[K, V]) *AAnode[K, V] {
	if node == nil {
		return nil
	}
	if node.L == nil {
		return node
	}
	return tm.min(node.L)
}

func (tm *AAmap[K, V]) delmin(node *AAnode[K, V]) *AAnode[K, V] {
	if node == nil {
		return nil
	}
	if node.L == nil {
		return node.R
	}
	node.L = tm.delmin(node.L)
	return node
}

func (tm *AAmap[K, V]) balanceDel(node *AAnode[K, V]) *AAnode[K, V] {
	if node == nil {
		return nil
	}
	node = tm.decrLvl(node)
	node = tm.skew(node)
	if node.R != nil {
		node.R = tm.skew(node.R)
	}
	if (node.R != nil) && (node.R.R != nil) {
		node.R.R = tm.skew(node.R.R)
	}
	node = tm.split(node)
	if node.R != nil {
		node.R = tm.split(node.R)
	}
	return node
}

func (tm *AAmap[K, V]) decrLvl(node *AAnode[K, V]) *AAnode[K, V] {
	if node == nil {
		return nil
	}
	lvl := byte(1)
	if (node.L != nil) && (node.R != nil) {
		lvl = amath.Min(node.L.Lvl, node.R.Lvl) + 1
	}
	if lvl < node.Lvl {
		node.Lvl = lvl
		if (node.R != nil) && (lvl < node.R.Lvl) {
			node.R.Lvl = lvl
		}
	}
	return node
}

func (tm *AAmap[K, V]) IterateKeys() []K {
	keysarr := make([]K, 0, tm.Size())
	return tm.iteratekeys(tm.root, keysarr)
}

func (tm *AAmap[K, V]) iteratekeys(node *AAnode[K, V], keysarr []K) []K {
	if node == nil {
		return keysarr
	}
	keysarr = tm.iteratekeys(node.L, keysarr)
	keysarr = append(keysarr, node.Key)
	keysarr = tm.iteratekeys(node.R, keysarr)
	return keysarr
}

func (tm *AAmap[K, V]) PrintTreeCheck() {
	fmt.Printf("\n")
	fmt.Printf("<- left rotated tree pic <-\n")
	tm.printtree(tm.root, 0)
	fmt.Printf("=============================\n\n")
}

func (tm *AAmap[K, V]) printtree(node *AAnode[K, V], n int) {
	if node != nil {
		margin := 5
		tm.printtree(node.R, n+margin)
		for i := 0; i < n; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%v^%v:%v\n", node.Lvl, node.Key, node.Val)
		tm.printtree(node.L, n+margin)
	}
}
