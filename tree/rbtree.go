package tree

import (
	// amath "github.com/kselnaag/algos/math"
	I "github.com/kselnaag/algos/types"
)

// left-linear red-black tree for mapping
type RBnode[K I.Ord, V any] struct {
	Key   K
	Val   V
	L, R  *RBnode[K, V]
	IsRed bool
}

type RBmap[K I.Ord, V any] struct {
	root *RBnode[K, V]
	size int
}

func (tm *RBmap[K, V]) NewTmap() RBmap[K, V] {
	return RBmap[K, V]{
		root: nil,
		size: 0,
	}
}
