package tree

import (
	"fmt"
	"math"

	I "github.com/kselnaag/algos/types"
)

// tree based Priority Queue
type PQnode[K I.Ord, V any] struct {
	Key     K
	Val     V
	P, L, R *PQnode[K, V]
}

// ===========================
type MinPQ[K I.Ord, V any] struct {
	root *PQnode[K, V]
	size int
}

func NewMinPQ[K I.Ord, V any]() MinPQ[K, V] {
	return MinPQ[K, V]{
		root: nil,
		size: 0,
	}
}

func (min *MinPQ[K, V]) Size() int {
	return min.size
}

func (min *MinPQ[K, V]) IsEmpty() bool {
	return min.Size() == 0
}

func (min *MinPQ[K, V]) Add(key K, val V) {
	if min.IsEmpty() {
		min.root = &PQnode[K, V]{Key: key, Val: val, P: nil, L: nil, R: nil}
		min.root.P = min.root
	} else {
		nodePar := min.lastRightNodePar(min.root.P)
		node := &PQnode[K, V]{Key: key, Val: val, P: nodePar, L: nil, R: nil}
		if nodePar.L == nil {
			nodePar.L = node
		} else {
			nodePar.R = node
		}
		min.root.P = node
		min.swimNode(node)
	}
	min.size++
}

func (min *MinPQ[K, V]) lastRightNodePar(lastnode *PQnode[K, V]) *PQnode[K, V] {
	if lastnode == nil {
		return nil
	}
	if lastnode == min.root {
		return min.root
	}
	node := lastnode.P
	if lastnode == node.R {
		for (node != min.root) && (lastnode == node.R) {
			lastnode = node
			node = node.P
		}
		if lastnode != node.R {
			node = node.R
		}
		for node.L != nil {
			node = node.L
		}
	}
	return node
}

func (min *MinPQ[K, V]) swapNode(i, j *PQnode[K, V]) {
	i.Key, i.Val, j.Key, j.Val = j.Key, j.Val, i.Key, i.Val
}

func (min *MinPQ[K, V]) swimNode(node *PQnode[K, V]) {
	if (node == nil) || (node.P == nil) {
		return
	}
	for (node != min.root) && I.LT(node.Key, node.P.Key) {
		min.swapNode(node, node.P)
		node = node.P
	}
}

func (min *MinPQ[K, V]) Min() (key K, val V) {
	return min.root.Key, min.root.Val
}

func (min *MinPQ[K, V]) DelMin() (key K, val V) {
	if min.IsEmpty() {
		panic("algos.tree.(MinPQ).DelMin(): No any node found, tree is empty")
	}
	key, val = min.root.Key, min.root.Val
	if min.size == 1 {
		min.root = nil
		min.size--
		return key, val
	}
	node := min.lastLeftNode(min.root.P)
	min.swapNode(min.root.P, min.root)
	if min.root.P.P.R == nil {
		min.root.P.P.L = nil
	} else {
		min.root.P.P.R = nil
	}
	min.root.P = node
	min.sinkNode(min.root)
	min.size--
	return key, val
}

func (min *MinPQ[K, V]) lastLeftNode(lastnode *PQnode[K, V]) *PQnode[K, V] {
	if lastnode == nil {
		return nil
	}
	if lastnode == min.root {
		return nil
	}
	node := lastnode.P
	if lastnode != node.L {
		return node.L
	} else {
		for (node != min.root) && (lastnode == node.L) {
			lastnode = node
			node = node.P
		}
		if lastnode != node.L {
			node = node.L
		}
		for node.R != nil {
			node = node.R
		}
	}
	return node
}

func (min *MinPQ[K, V]) sinkNode(node *PQnode[K, V]) {
	for {
		if (node == nil) || ((node.L == nil) && (node.R == nil)) {
			break
		}
		if (node.R == nil) && (node.L != nil) && I.GT(node.Key, node.L.Key) {
			min.swapNode(node, node.L)
			node = node.L
			continue
		}
		if (node.L == nil) && (node.R != nil) && I.GT(node.Key, node.R.Key) {
			min.swapNode(node, node.R)
			node = node.R
			continue
		}
		if I.GT(node.L.Key, node.R.Key) {
			min.swapNode(node.L, node.R)
		}
		if I.GT(node.Key, node.L.Key) {
			min.swapNode(node, node.L)
			node = node.L
			continue
		} else if I.GT(node.Key, node.R.Key) {
			min.swapNode(node, node.R)
			node = node.R
			continue
		}
		break
	}
}

func (min *MinPQ[K, V]) Iterate() []K {
	keysarr := make([]K, 0, min.Size())
	return min.iterate(min.root, keysarr)
}

func (min *MinPQ[K, V]) iterate(node *PQnode[K, V], keysarr []K) []K {
	if node == nil {
		return keysarr
	}
	keysarr = append(keysarr, node.Key)
	keysarr = min.iterate(node.L, keysarr)
	keysarr = min.iterate(node.R, keysarr)
	return keysarr
}

func (min *MinPQ[K, V]) HeightTreeCheck() int {
	if min.IsEmpty() {
		return 0
	}
	sqrt := math.Sqrt(float64(min.size))
	return int(math.Ceil(sqrt))
}

func (min *MinPQ[K, V]) IsCompleteBTcheck() bool {
	return min.iscompleteBTcheck(min.root, min.Size())
}

func (min *MinPQ[K, V]) iscompleteBTcheck(node *PQnode[K, V], size int) bool {
	if node == nil {
		return true
	}
	sqrt := math.Sqrt(float64(size))
	CBTlvl := math.Ceil(sqrt)
	maxCBTsize := int(math.Pow(2, CBTlvl))
	arr := make([]*PQnode[K, V], maxCBTsize)
	arr[1] = node
	for i := 1; i < maxCBTsize; i++ {
		if arr[i].L != nil {
			arr[2*i] = arr[i].L
		}
		if arr[i].R != nil {
			arr[2*i+1] = arr[i].R
		}
	}
	for i := 1; i <= size; i++ {
		if arr[i] == nil {
			return false
		}
	}
	return true
}

func (min *MinPQ[K, V]) PrintTreeCheck() {
	fmt.Printf("\n")
	fmt.Printf("<- left rotated tree pic <-\n")
	min.printtree(min.root, 0)
	fmt.Printf("=============================\n\n")
}

func (min *MinPQ[K, V]) printtree(node *PQnode[K, V], n int) {
	if node != nil {
		min.printtree(node.R, n+5)
		for i := 0; i < n; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%v:%v\n", node.Key, node.Val)
		min.printtree(node.L, n+5)
	}
}

/*
NewMinPQ Size IsEmpty HeightTree Add swimNode swapNode
Min DelMin sinkNode Iterate iterate
lastRightNodePar lastLeftNode

CompleteBTCheck HeightTreeCheck PrintTreeCheck printtree
*/

// ===========================
