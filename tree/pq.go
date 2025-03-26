package tree

import (
	"fmt"
	"math"

	I "algos/types"
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

func (minpq *MinPQ[K, V]) Size() int {
	return minpq.size
}

func (minpq *MinPQ[K, V]) IsEmpty() bool {
	return minpq.Size() == 0
}

func (minpq *MinPQ[K, V]) Add(key K, val V) {
	if minpq.IsEmpty() {
		minpq.root = &PQnode[K, V]{Key: key, Val: val, P: nil, L: nil, R: nil}
		minpq.root.P = minpq.root
	} else {
		nodePar := minpq.lastRightNodePar(minpq.root.P)
		node := &PQnode[K, V]{Key: key, Val: val, P: nodePar, L: nil, R: nil}
		if nodePar.L == nil {
			nodePar.L = node
		} else {
			nodePar.R = node
		}
		minpq.root.P = node
		minpq.swimNode(node)
	}
	minpq.size++
}

func (minpq *MinPQ[K, V]) lastRightNodePar(lastnode *PQnode[K, V]) *PQnode[K, V] {
	if lastnode == nil {
		return nil
	}
	if lastnode == minpq.root {
		return minpq.root
	}
	node := lastnode.P
	if lastnode == node.R {
		for (node != minpq.root) && (lastnode == node.R) {
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

func (minpq *MinPQ[K, V]) swapNode(i, j *PQnode[K, V]) {
	i.Key, i.Val, j.Key, j.Val = j.Key, j.Val, i.Key, i.Val
}

func (minpq *MinPQ[K, V]) swimNode(node *PQnode[K, V]) {
	if (node == nil) || (node.P == nil) {
		return
	}
	for (node != minpq.root) && I.LT(node.Key, node.P.Key) {
		minpq.swapNode(node, node.P)
		node = node.P
	}
}

func (minpq *MinPQ[K, V]) Min() (key K, val V) {
	return minpq.root.Key, minpq.root.Val
}

func (minpq *MinPQ[K, V]) DelMin() (key K, val V) {
	if minpq.IsEmpty() {
		panic("algos.tree.(MinPQ).DelMin(): No any node found, tree is empty")
	}
	key, val = minpq.root.Key, minpq.root.Val
	if minpq.size == 1 {
		minpq.root = nil
		minpq.size--
		return key, val
	}
	node := minpq.lastLeftNode(minpq.root.P)
	minpq.swapNode(minpq.root.P, minpq.root)
	if minpq.root.P.P.R == nil {
		minpq.root.P.P.L = nil
	} else {
		minpq.root.P.P.R = nil
	}
	minpq.root.P = node
	minpq.sinkNode(minpq.root)
	minpq.size--
	return key, val
}

func (minpq *MinPQ[K, V]) lastLeftNode(lastnode *PQnode[K, V]) *PQnode[K, V] {
	if lastnode == nil {
		return nil
	}
	if lastnode == minpq.root {
		return nil
	}
	node := lastnode.P
	if lastnode != node.L {
		return node.L
	} else {
		for (node != minpq.root) && (lastnode == node.L) {
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

func (minpq *MinPQ[K, V]) sinkNode(node *PQnode[K, V]) {
	for {
		if (node == nil) || ((node.L == nil) && (node.R == nil)) { //nolint:staticcheck // in case of consistency
			break
		}
		if (node.R == nil) && (node.L != nil) && I.GT(node.Key, node.L.Key) {
			minpq.swapNode(node, node.L)
			node = node.L
			continue
		}
		if (node.L == nil) && (node.R != nil) && I.GT(node.Key, node.R.Key) {
			minpq.swapNode(node, node.R)
			node = node.R
			continue
		}
		if I.GT(node.L.Key, node.R.Key) {
			minpq.swapNode(node.L, node.R)
		}
		if I.GT(node.Key, node.L.Key) {
			minpq.swapNode(node, node.L)
			node = node.L
			continue
		} else if I.GT(node.Key, node.R.Key) {
			minpq.swapNode(node, node.R)
			node = node.R
			continue
		}
		break
	}
}

func (minpq *MinPQ[K, V]) Iterate() []K {
	keysarr := make([]K, 0, minpq.Size())
	return minpq.iterate(minpq.root, keysarr)
}

func (minpq *MinPQ[K, V]) iterate(node *PQnode[K, V], keysarr []K) []K {
	if node == nil {
		return keysarr
	}
	keysarr = append(keysarr, node.Key)
	keysarr = minpq.iterate(node.L, keysarr)
	keysarr = minpq.iterate(node.R, keysarr)
	return keysarr
}

func (minpq *MinPQ[K, V]) HeightTreeCheck() int {
	if minpq.IsEmpty() {
		return 0
	}
	sqrt := math.Sqrt(float64(minpq.size))
	return int(math.Ceil(sqrt))
}

func (minpq *MinPQ[K, V]) IsCompleteBTcheck() bool {
	return minpq.iscompleteBTcheck(minpq.root, minpq.Size())
}

func (minpq *MinPQ[K, V]) iscompleteBTcheck(node *PQnode[K, V], size int) bool {
	if node == nil {
		return true
	}
	sqrt := math.Sqrt(float64(size))
	CBTlvl := math.Ceil(sqrt)
	base := 2.0
	maxCBTsize := int(math.Pow(base, CBTlvl))
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

func (minpq *MinPQ[K, V]) PrintTreeCheck() {
	fmt.Printf("\n")
	fmt.Printf("<- left rotated tree pic <-\n")
	minpq.printtree(minpq.root, 0)
	fmt.Printf("=============================\n\n")
}

func (minpq *MinPQ[K, V]) printtree(node *PQnode[K, V], n int) {
	if node != nil {
		margin := 5
		minpq.printtree(node.R, n+margin)
		for range n {
			fmt.Printf(" ")
		}
		fmt.Printf("%v:%v\n", node.Key, node.Val)
		minpq.printtree(node.L, n+margin)
	}
}

/*
NewMinPQ Size IsEmpty HeightTree Add swimNode swapNode
Min DelMin sinkNode Iterate iterate
lastRightNodePar lastLeftNode

CompleteBTCheck HeightTreeCheck PrintTreeCheck printtree
*/

// ===========================
