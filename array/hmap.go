package array

import (
	"math"

	amath "github.com/kselnaag/algos/math"
	I "github.com/kselnaag/algos/types"
)

type kvnode[K I.Ord, V any] struct {
	key K
	val V
}

type Hmap[K I.Ord, V any] struct {
	hmarr   [math.MaxUint16 + 1]*[]kvnode[K, V]
	keysnum int
}

func NewHmap[K I.Ord, V any]() Hmap[K, V] {
	return Hmap[K, V]{
		hmarr:   [math.MaxUint16 + 1]*[]kvnode[K, V]{},
		keysnum: 0,
	}
}

func (hm *Hmap[K, V]) Drop() {
	hm.hmarr = [math.MaxUint16 + 1]*[]kvnode[K, V]{}
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
	node := kvnode[K, V]{key: key, val: val}
	if hm.hmarr[int(hash)] == nil {
		hm.hmarr[int(hash)] = &[]kvnode[K, V]{node}
		hm.keysnum++
		return
	}
	for i, hnode := range *hm.hmarr[int(hash)] {
		if hnode.key == key {
			(*hm.hmarr[int(hash)])[i] = node
			return
		}
	}
	*hm.hmarr[int(hash)] = append(*hm.hmarr[int(hash)], node)
	hm.keysnum++
}

func (hm *Hmap[K, V]) IsKey(key K) bool {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	if hm.hmarr[int(hash)] == nil {
		return false
	}
	for _, hnode := range *hm.hmarr[int(hash)] {
		if hnode.key == key {
			return true
		}
	}
	return false
}

func (hm *Hmap[K, V]) Get(key K) V {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	if hm.hmarr[int(hash)] == nil {
		panic("algos.array.(Hmap).Get(key K): No any key found, check first")
	}
	for _, hnode := range *hm.hmarr[int(hash)] {
		if hnode.key == key {
			return hnode.val
		}
	}
	panic("algos.array.(Hmap).Get(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) Del(key K) {
	abytes := amath.ConvToByteArr(key)
	hash := amath.HashPirson16(abytes)
	if hm.hmarr[int(hash)] == nil {
		panic("algos.array.(Hmap).Del(key K): No any key found, check first")
	}
	for i, hnode := range *hm.hmarr[int(hash)] {
		if hnode.key == key {
			*hm.hmarr[int(hash)] = append((*hm.hmarr[int(hash)])[:i], (*hm.hmarr[int(hash)])[i+1:]...)
			hm.keysnum--
			return
		}
	}
	panic("algos.array.(Hmap).Del(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) IterateKeys() []K {
	res := []K{}
	fwd := uint16(0)
	for {
		if hm.hmarr[int(fwd)] != nil {
			for _, hnode := range *hm.hmarr[int(fwd)] {
				res = append(res, hnode.key)
			}
		}
		fwd++
		if fwd == 0 {
			break
		}
	}
	QuickSort(res)
	return res
}
