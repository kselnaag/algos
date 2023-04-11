package list

import (
	"github.com/kselnaag/algos/amath"
	I "github.com/kselnaag/algos/types"
)

const maxUint16 int = 1<<16 - 1

type Mnode[K I.Ord, V any] struct {
	Key  K
	Val  V
	Next *Mnode[K, V]
}

type Hmap[K I.Ord, V any] struct {
	hmarr   [maxUint16]*Mnode[K, V]
	keysnum int
}

func NewHmap[K I.Ord, V any]() Hmap[K, V] {
	return Hmap[K, V]{
		hmarr:   [maxUint16]*Mnode[K, V]{},
		keysnum: 0,
	}
}

func (hm *Hmap[K, V]) Size() int {
	return hm.keysnum
}

func (hm *Hmap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm *Hmap[K, V]) Set(key K, val V) {
	abytes := I.ConvToByteArr(key)
	hash := amath.HashPirson[uint16](abytes)
	setnode := &Mnode[K, V]{Key: key, Val: val, Next: nil}
	for node := hm.hmarr[int(hash)]; node != nil; node = node.Next {
		if node.Key == key {
			node.Val = val
			return
		}
	}
	setnode.Next = hm.hmarr[int(hash)]
	hm.hmarr[int(hash)] = setnode
	hm.keysnum++
}

func (hm *Hmap[K, V]) IsKey(key K) bool {
	abytes := I.ConvToByteArr(key)
	hash := amath.HashPirson[uint16](abytes)
	for node := hm.hmarr[int(hash)]; node != nil; node = node.Next {
		if node.Key == key {
			return true
		}
	}
	return false
}

func (hm *Hmap[K, V]) Get(key K) V {
	abytes := I.ConvToByteArr(key)
	hash := amath.HashPirson[uint16](abytes)
	for node := hm.hmarr[int(hash)]; node != nil; node = node.Next {
		if node.Key == key {
			return node.Val
		}
	}
	panic("algos.list.(Hmap).Get(key K): No any key found, check first")
}

func (hm *Hmap[K, V]) Del(key K) {
	abytes := I.ConvToByteArr(key)
	hash := amath.HashPirson[uint16](abytes)
	prev := hm.hmarr[int(hash)]
	for node := hm.hmarr[int(hash)]; node != nil; node = node.Next {
		if node.Key == key {
			if prev == hm.hmarr[int(hash)] {
				hm.hmarr[int(hash)] = node.Next
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
