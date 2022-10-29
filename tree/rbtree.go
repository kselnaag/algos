package tree

import (
	"fmt"

	amath "github.com/kselnaag/algos/math"
	I "github.com/kselnaag/algos/types"
)

// left-linear red-black tree with lazy node deletion
type RBnode[K I.Ord, V any] struct {
	Key          K
	Val          V
	L, R         *RBnode[K, V]
	IsRed, IsDel bool
}

type RBmap[K I.Ord, V any] struct {
	root *RBnode[K, V]
	size int
}

func NewRBmap[K I.Ord, V any]() RBmap[K, V] {
	return RBmap[K, V]{
		root: nil,
		size: 0,
	}
}

func (tm *RBmap[K, V]) Size() int {
	return tm.size
}

func (tm *RBmap[K, V]) IsEmpty() bool {
	return tm.Size() == 0
}

func (tm *RBmap[K, V]) IsKey(key K) bool {
	node := tm.root
	for node != nil {
		if I.LT(key, node.Key) {
			node = node.L
		} else if I.GT(key, node.Key) {
			node = node.R
		} else {
			return !node.IsDel
		}
	}
	return false
}

func (tm *RBmap[K, V]) Get(key K) V {
	node := tm.root
	for node != nil {
		if I.LT(key, node.Key) {
			node = node.L
		} else if I.GT(key, node.Key) {
			node = node.R
		} else {
			if !node.IsDel {
				return node.Val
			}
			node = nil
		}
	}
	panic("algos.tree.(RBmap).Get(key): No any key found, check first")
}

func (tm *RBmap[K, V]) Del(key K) {
	if tm.IsEmpty() {
		panic("algos.tree.(RBmap).Del(): No any key found, tree is empty")
	}
	node := tm.root
	for node != nil {
		if I.LT(key, node.Key) {
			node = node.L
		} else if I.GT(key, node.Key) {
			node = node.R
		} else {
			node.IsDel = true
			if tm.size > 0 {
				tm.size--
			}
			return
		}
	}
	panic("algos.tree.(RBmap).Del(key): No any key found, check first")
}

func (tm *RBmap[K, V]) Put(key K, val V) {
	tm.root = tm.put(tm.root, key, val)
	if !tm.IsEmpty() {
		tm.root.IsRed = false
	}
}

func (tm *RBmap[K, V]) put(node *RBnode[K, V], key K, val V) *RBnode[K, V] {
	if node == nil {
		tm.size++
		node = &RBnode[K, V]{Key: key, Val: val, L: nil, R: nil, IsRed: true, IsDel: false}
	} else {
		if I.LT(key, node.Key) {
			node.L = tm.put(node.L, key, val)
		} else if I.GT(key, node.Key) {
			node.R = tm.put(node.R, key, val)
		} else {
			node.Val = val
			if node.IsDel {
				node.IsDel = false
				tm.size++
			}
		}
	}
	return tm.balance(node)
}

func (tm *RBmap[K, V]) balance(node *RBnode[K, V]) *RBnode[K, V] {
	if node == nil {
		return nil
	}
	if (node.L == nil || !node.L.IsRed) && (node.R != nil) && node.R.IsRed {
		node = tm.rotateLeft(node)
	}
	if (node.L != nil) && node.L.IsRed && (node.L.L != nil) && node.L.L.IsRed {
		node = tm.rotateRight(node)
	}
	if (node.L != nil) && node.L.IsRed && (node.R != nil) && node.R.IsRed {
		tm.flipColors(node)
	}
	return node
}

func (tm *RBmap[K, V]) rotateLeft(node *RBnode[K, V]) *RBnode[K, V] {
	if node == nil {
		return nil
	}
	x := node.R
	node.R = x.L
	x.L = node
	x.IsRed = node.IsRed
	node.IsRed = true
	return x
}

func (tm *RBmap[K, V]) rotateRight(node *RBnode[K, V]) *RBnode[K, V] {
	if node == nil {
		return nil
	}
	x := node.L
	node.L = x.R
	x.R = node
	x.IsRed = node.IsRed
	node.IsRed = true
	return x
}

func (tm *RBmap[K, V]) flipColors(node *RBnode[K, V]) {
	if node != nil {
		node.IsRed = true
		if node.L != nil {
			node.L.IsRed = false
		}
		if node.R != nil {
			node.R.IsRed = false
		}
	}
}

func (tm *RBmap[K, V]) IterateKeys() []K {
	keysarr := make([]K, 0, tm.Size())
	return tm.iteratekeys(tm.root, keysarr)
}

func (tm *RBmap[K, V]) iteratekeys(node *RBnode[K, V], keysarr []K) []K {
	if node == nil {
		return keysarr
	}
	keysarr = tm.iteratekeys(node.L, keysarr)
	if !node.IsDel {
		keysarr = append(keysarr, node.Key)
	}
	keysarr = tm.iteratekeys(node.R, keysarr)
	return keysarr
}

func (tm *RBmap[K, V]) IsBSTcheck() bool {
	return tm.isbst(tm.root, nil, nil)
}

func (tm *RBmap[K, V]) isbst(node *RBnode[K, V], min *K, max *K) bool {
	if node == nil {
		return true
	}
	if (min != nil) && (node.Key <= *min) {
		return false
	}
	if (max != nil) && (node.Key >= *max) {
		return false
	}
	return tm.isbst(node.L, min, &node.Key) && tm.isbst(node.R, &node.Key, max)
}

func (tm *RBmap[K, V]) BSTheightCheck() int {
	return tm.bstheight(tm.root)
}

func (tm *RBmap[K, V]) bstheight(node *RBnode[K, V]) int {
	if node == nil {
		return -1
	}
	return 1 + amath.Max(tm.bstheight(node.L), tm.bstheight(node.R))
}

func (tm *RBmap[K, V]) IsRedBlackCheck() bool {
	return tm.isredblack(tm.root)
}

func (tm *RBmap[K, V]) isredblack(node *RBnode[K, V]) bool {
	if node == nil {
		return true
	}
	if (node.R != nil) && node.R.IsRed {
		return false
	}
	if (node != tm.root) && node.IsRed && (node.L != nil) && node.L.IsRed {
		return false
	}
	return tm.isredblack(node.L) && tm.isredblack(node.R)
}

func (tm *RBmap[K, V]) IsBalancedCheck() bool {
	blacknum := 0
	node := tm.root
	for node != nil {
		if !node.IsRed {
			blacknum++
		}
		node = node.L
	}
	return tm.isbalanced(tm.root, blacknum)
}

func (tm *RBmap[K, V]) isbalanced(node *RBnode[K, V], blacknum int) bool {
	if node == nil {
		return blacknum == 0
	}
	if !node.IsRed {
		blacknum--
	}
	return tm.isbalanced(node.L, blacknum) && tm.isbalanced(node.R, blacknum)
}

func (tm *RBmap[K, V]) PrintTreeCheck() {
	fmt.Printf("\n")
	fmt.Printf("<- left rotated tree pic <-\n")
	tm.printtree(tm.root, 0)
	fmt.Printf("(+ Red node, x Del node, * Red and Del node )\n\n")
}

func (tm *RBmap[K, V]) printtree(node *RBnode[K, V], n int) {
	if node != nil {
		tm.printtree(node.R, n+5)
		for i := 0; i < n; i++ {
			fmt.Printf(" ")
		}
		if node.IsRed && node.IsDel {
			fmt.Printf("*")
		} else {
			if node.IsRed {
				fmt.Printf("+")
			}
			if node.IsDel {
				fmt.Printf("x")
			}
		}
		fmt.Printf("%v:%v\n", node.Key, node.Val)
		tm.printtree(node.L, n+5)
	}
}
