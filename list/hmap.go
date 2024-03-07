package list

import (
	"fmt"
	"sync"
)

// [list/hmap] This struct is size-optimized.
// If the number of K-V pairs is more than 10^6 it seems you should use some other tool.

const (
	MAXUINT16  int = (1 << 16)
	STARTBCKTS int = 10
	GROWCOEFF  int = 5
	GROWCOND   int = 10
)

type HMnode[K comparable, V any] struct {
	Key  K
	Val  V
	Next *HMnode[K, V]
}

type HMap[K comparable, V any] struct {
	keysnum int
	bktsnum int
	rwm     sync.RWMutex
	hmarr   []*HMnode[K, V]
}

func NewHMap[K comparable, V any](buckets ...uint16) *HMap[K, V] {
	bktsnum := STARTBCKTS
	if len(buckets) > 0 && (int(buckets[0]) > STARTBCKTS) {
		bktsnum = int(buckets[0])
	}
	return &HMap[K, V]{
		keysnum: 0,
		bktsnum: bktsnum,
		hmarr:   make([]*HMnode[K, V], bktsnum),
	}
}

func (hm *HMap[K, V]) Buckets() int {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	return hm.bktsnum
}

func (hm *HMap[K, V]) Size() int {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	return hm.keysnum
}

func (hm *HMap[K, V]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm *HMap[K, V]) convToBytes(key K) []byte {
	return []byte(fmt.Sprintf("%v", key))
}

func (hm *HMap[K, V]) hashDJB2a(data []byte) uint32 {
	var hash uint32 = 5381
	mlen := len(data)
	for i := 0; i < mlen; i++ {
		hash = ((hash << 5) + hash) ^ uint32(data[i])
	}
	return hash
}

func (hm *HMap[K, V]) hashFromKey(key K) int {
	bytesarr := hm.convToBytes(key)
	hash := hm.hashDJB2a(bytesarr)
	return int(hash&0x0000ffff) % hm.bktsnum
}

func (hm *HMap[K, V]) evacuate() {
	newbkts := hm.bktsnum * GROWCOEFF
	switch { // isSpaceToGrow
	case newbkts < MAXUINT16:
		hm.bktsnum = newbkts
	case newbkts > MAXUINT16:
		hm.bktsnum = MAXUINT16
	default:
		return
	}

	newhmarr := make([]*HMnode[K, V], hm.bktsnum)
	for i, ptr := range hm.hmarr {
		if ptr != nil {
			for node, next := ptr, ptr.Next; node != nil; {
				idx := hm.hashFromKey(node.Key)
				node.Next = newhmarr[idx]
				newhmarr[idx] = node

				node = next
				if next != nil {
					next = next.Next
				}
			}
			hm.hmarr[i] = nil
		}
	}
	hm.hmarr = newhmarr
}

func (hm *HMap[K, V]) Set(key K, val V) {
	hm.rwm.Lock()
	defer hm.rwm.Unlock()

	newnode := &HMnode[K, V]{Key: key, Val: val, Next: nil}
	hashIDX := hm.hashFromKey(key)
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			node.Val = val
			return
		}
	}

	isNeedToGrow := (hm.keysnum >= (hm.bktsnum * GROWCOND))
	if isNeedToGrow {
		hm.evacuate()
	}

	newnode.Next = hm.hmarr[hashIDX]
	hm.hmarr[hashIDX] = newnode
	hm.keysnum++
}

func (hm *HMap[K, V]) Del(key K) *V {
	hm.rwm.Lock()
	defer hm.rwm.Unlock()

	hashIDX := hm.hashFromKey(key)
	prev := hm.hmarr[hashIDX]
	for node := prev; node != nil; node = node.Next {
		if node.Key == key {
			val := &node.Val
			if prev == hm.hmarr[hashIDX] {
				hm.hmarr[hashIDX] = node.Next
			} else {
				prev.Next = node.Next
			}
			hm.keysnum--
			return val
		}
		prev = node
	}
	return nil
}

func (hm *HMap[K, V]) Get(key K) *V {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	hashIDX := hm.hashFromKey(key)
	for node := hm.hmarr[hashIDX]; node != nil; node = node.Next {
		if node.Key == key {
			return &node.Val
		}
	}
	return nil
}

func (hm *HMap[K, V]) IterateKeys() []K {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	keysarr := make([]K, 0, hm.keysnum)
	for _, ptr := range hm.hmarr {
		for node := ptr; node != nil; node = node.Next {
			keysarr = append(keysarr, node.Key)
		}
	}
	return keysarr
}

func (hm *HMap[K, V]) PrintAll() {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	for idx, ptr := range hm.hmarr {
		fmt.Printf("| %d | %p |", idx, ptr)
		for node := ptr; node != nil; node = node.Next {
			fmt.Printf(" -> [key:%v, val:%v, next:%p]", node.Key, node.Val, node.Next)
		}
		fmt.Printf("\n")
	}
}
