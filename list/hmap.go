package list

import (
	"github.com/kselnaag/algos/amath"
	I "github.com/kselnaag/algos/types"
)

const maxUint8 int = (1 << 8)

type Mnode[K I.Ord, V any] struct {
	Key  K
	Val  V
	Next *Mnode[K, V]
}

type Hmap[K I.Ord, V any] struct {
	hmarr   [maxUint8]*Mnode[K, V]
	keysnum int
}

func NewHmap[K I.Ord, V any]() Hmap[K, V] {
	return Hmap[K, V]{
		hmarr:   [maxUint8]*Mnode[K, V]{},
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
	hashIDX := hashFromKey(key)
	setnode := &Mnode[K, V]{Key: key, Val: val, Next: nil}
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			node.Val = val
			return
		}
	}
	setnode.Next = hm.hmarr[hashIDX]
	hm.hmarr[hashIDX] = setnode
	hm.keysnum++
}

func (hm *Hmap[K, V]) IsKey(key K) bool {
	hashIDX := hashFromKey(key)
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			return true
		}
	}
	return false
}

func (hm *Hmap[K, V]) Get(key K) V {
	hashIDX := hashFromKey(key)
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			return node.Val
		}
	}
	panic("algos.list.(Hmap).Get(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) Del(key K) {
	hashIDX := hashFromKey(key)
	prev := hm.hmarr[hashIDX]
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			if prev == hm.hmarr[hashIDX] {
				hm.hmarr[hashIDX] = node.Next
			} else {
				prev.Next = node.Next
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
		for node := ptr; node != nil; node = node.Next {
			res = append(res, node.Key)
		}
	}
	return res
}
