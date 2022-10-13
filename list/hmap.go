package list

import (
	"math"

	"github.com/kselnaag/algos/array"
	amath "github.com/kselnaag/algos/math"
	I "github.com/kselnaag/algos/types"
)

type kvnode[K I.Ord, V any] struct {
	key  K
	val  V
	next *kvnode[K, V]
}

type Hmap[K I.Ord, V any] struct {
	hmarr   [math.MaxUint16]*kvnode[K, V]
	keysnum int
}

func NewHmap[K I.Ord, V any]() Hmap[K, V] {
	return Hmap[K, V]{
		hmarr:   [math.MaxUint16]*kvnode[K, V]{},
		keysnum: 0,
	}
}

func (hm *Hmap[K, V]) Drop() {
	hm.hmarr = [math.MaxUint16]*kvnode[K, V]{}
	hm.keysnum = 0
}

func (hm *Hmap[K, V]) Size() int {
	return hm.keysnum
}

func (hm *Hmap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm *Hmap[K, V]) Set(key K, val V) {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	setnode := kvnode[K, V]{key: key, val: val, next: nil}
	for node := hm.hmarr[int(hash)]; node != nil; node = node.next {
		if node.key == key {
			node.val = val
			return
		}
	}
	setnode.next = hm.hmarr[int(hash)]
	hm.hmarr[int(hash)] = &setnode
	hm.keysnum++
}

func (hm *Hmap[K, V]) IsKey(key K) bool {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	for node := hm.hmarr[int(hash)]; node != nil; node = node.next {
		if node.key == key {
			return true
		}
	}
	return false
}

func (hm *Hmap[K, V]) Get(key K) V {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	for node := hm.hmarr[int(hash)]; node != nil; node = node.next {
		if node.key == key {
			return node.val
		}
	}
	panic("algos.list.(Hmap).Get(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) Del(key K) {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	prev := hm.hmarr[int(hash)]
	for node := hm.hmarr[int(hash)]; node != nil; node = node.next {
		if node.key == key {
			if prev == hm.hmarr[int(hash)] {
				hm.hmarr[int(hash)] = node.next
			} else {
				prev.next = node.next
			}
			hm.keysnum--
			return
		}
		prev = node
	}
	panic("algos.list.(Hmap).Del(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) IterateKeys() []K {
	res := []K{}
	for _, ptr := range hm.hmarr {
		for node := ptr; node != nil; node = node.next {
			res = append(res, node.key)
		}
	}
	array.QuickSort(res)
	return res
}
