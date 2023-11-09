package array

import (
	"github.com/kselnaag/algos/amath"
	I "github.com/kselnaag/algos/types"
)

const maxUint8 int = (1 << 8)

type Anode[K I.Ord, V any] struct {
	Key K
	Val V
}

type Hmap[K I.Ord, V any] struct {
	hmarr   [maxUint8]*[]Anode[K, V]
	keysnum int
}

func NewHmap[K I.Ord, V any]() Hmap[K, V] {
	return Hmap[K, V]{
		hmarr:   [maxUint8]*[]Anode[K, V]{},
		keysnum: 0,
	}
}

func (hm *Hmap[K, V]) Size() int {
	return hm.keysnum
}

func (hm *Hmap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func hashFromKey[K I.Ord](key K) int {
	bytesarr := I.ConvToByteArr(key)
	hash := amath.HashDJB2a[uint32](bytesarr)
	return int(hash & 0x000000FF)
}

func (hm *Hmap[K, V]) Set(key K, val V) {
	hashIdx := hashFromKey(key)
	node := Anode[K, V]{Key: key, Val: val}
	if hm.hmarr[hashIdx] == nil {
		hm.hmarr[hashIdx] = &[]Anode[K, V]{node}
		hm.keysnum++
		return
	}
	for i, hnode := range *hm.hmarr[hashIdx] {
		if hnode.Key == key {
			(*hm.hmarr[hashIdx])[i] = node
			return
		}
	}
	*hm.hmarr[hashIdx] = append(*hm.hmarr[hashIdx], node)
	hm.keysnum++
}

func (hm *Hmap[K, V]) IsKey(key K) bool {
	hashIdx := hashFromKey(key)
	if hm.hmarr[hashIdx] == nil {
		return false
	}
	for _, hnode := range *hm.hmarr[hashIdx] {
		if hnode.Key == key {
			return true
		}
	}
	return false
}

func (hm *Hmap[K, V]) Get(key K) V {
	hashIdx := hashFromKey(key)
	if hm.hmarr[hashIdx] == nil {
		panic("algos.array.(Hmap).Get(key K): No any key found, check first")
	}
	for _, hnode := range *hm.hmarr[hashIdx] {
		if hnode.Key == key {
			return hnode.Val
		}
	}
	panic("algos.array.(Hmap).Get(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) Del(key K) {
	hashIdx := hashFromKey(key)
	if hm.hmarr[hashIdx] == nil {
		panic("algos.array.(Hmap).Del(key K): No any key found, check first")
	}
	for i, hnode := range *hm.hmarr[hashIdx] {
		if hnode.Key == key {
			*hm.hmarr[hashIdx] = append((*hm.hmarr[hashIdx])[:i], (*hm.hmarr[hashIdx])[i+1:]...)
			hm.keysnum--
			return
		}
	}
	panic("algos.array.(Hmap).Del(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) IterateKeys() []K {
	res := []K{}
	for _, ptr := range hm.hmarr {
		if ptr != nil {
			for _, kvnode := range *ptr {
				res = append(res, kvnode.Key)
			}
		}
	}
	return res
}
