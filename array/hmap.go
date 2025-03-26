package array

import (
	"fmt"
	"sync"
)

// [array/hmap] Open-addressed.
// If the number of K-V pairs is more than 2^15 (32 768) it seems you should use some other tool.

const (
	MAXUINT16  int     = (1 << 16)
	STARTBCKTS int     = 200
	GROWCOND   float64 = 0.5
	GROWCOEFF  int     = 2
)

type OAnode[K comparable, V any] struct {
	Key K
	Val V
}

type HMap[K comparable, V any] struct {
	keysnum int
	bktsnum int
	rwm     sync.RWMutex
	hmarr   []*OAnode[K, V]
}

func NewHMap[K comparable, V any](buckets ...uint16) *HMap[K, V] {
	bktsnum := STARTBCKTS
	if (len(buckets) > 0) && (int(buckets[0]) > STARTBCKTS) {
		bktsnum = int(buckets[0])
	}
	return &HMap[K, V]{
		keysnum: 0,
		bktsnum: bktsnum,
		hmarr:   make([]*OAnode[K, V], bktsnum),
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
	return (hm.Size() == 0)
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

	newhmarr := make([]*OAnode[K, V], hm.bktsnum)
	for i, ptr := range hm.hmarr {
		if ptr == nil {
			continue
		}
		hashIdx := hm.hashFromKey(ptr.Key)
		cycl := 0
		for newhmarr[hashIdx] != nil {
			hashIdx++
			if hashIdx >= hm.bktsnum {
				hashIdx = 0
			}
			cycl++
			if cycl >= hm.bktsnum*2 {
				return
			}
		}
		newhmarr[hashIdx] = ptr
		hm.hmarr[i] = nil
	}
	hm.hmarr = newhmarr
}

func (hm *HMap[K, V]) Set(key K, val V) {
	hm.rwm.Lock()
	defer hm.rwm.Unlock()

	newnode := &OAnode[K, V]{Key: key, Val: val}
	hashIdx := hm.hashFromKey(key)
	if (hm.hmarr[hashIdx] != nil) && (hm.hmarr[hashIdx].Key == key) {
		hm.hmarr[hashIdx].Val = val
		return
	}

	isNeedToGrow := (hm.keysnum >= (int(float64(hm.bktsnum) * GROWCOND)))
	if isNeedToGrow {
		hm.evacuate()
		hashIdx = hm.hashFromKey(key)
	}

	cycl := 0
	for hm.hmarr[hashIdx] != nil {
		hashIdx++
		if hashIdx >= hm.bktsnum {
			hashIdx = 0
		}
		cycl++
		if cycl >= hm.bktsnum*2 {
			return
		}
	}
	hm.hmarr[hashIdx] = newnode
	hm.keysnum++
}

func (hm *HMap[K, V]) Del(key K) *V {
	hm.rwm.Lock()
	defer hm.rwm.Unlock()

	var (
		hashIdx = hm.hashFromKey(key)
		cycl    int
		val     V
		keyIdx  int
		found   bool
	)
	for hm.hmarr[hashIdx] != nil {
		if hm.hmarr[hashIdx].Key == key {
			keyIdx = hashIdx
			val = hm.hmarr[hashIdx].Val
			hm.keysnum--
			found = true
		}
		hashIdx++
		if hashIdx >= hm.bktsnum {
			hashIdx = 0
		}
		cycl++
		if cycl >= hm.bktsnum*2 {
			break
		}
	}
	if found {
		if hashIdx == 0 {
			hashIdx = hm.bktsnum
		}
		hashIdx--
		hm.hmarr[keyIdx] = hm.hmarr[hashIdx]
		hm.hmarr[hashIdx] = nil
	}
	return &val
}

func (hm *HMap[K, V]) Get(key K) *V {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	cycl := 0
	hashIdx := hm.hashFromKey(key)
	for hm.hmarr[hashIdx] != nil {
		if hm.hmarr[hashIdx].Key == key {
			return &hm.hmarr[hashIdx].Val
		}
		hashIdx++
		if hashIdx >= hm.bktsnum {
			hashIdx = 0
		}
		cycl++
		if cycl >= hm.bktsnum*2 {
			break
		}
	}
	return nil
}

func (hm *HMap[K, V]) IterateKeys() []K {
	hm.rwm.RLock()
	defer hm.rwm.RUnlock()

	keysarr := make([]K, 0, hm.keysnum)
	for _, node := range hm.hmarr {
		if node != nil {
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
		if ptr != nil {
			fmt.Printf(" -> [key:%v, val:%v]", ptr.Key, ptr.Val)
		}
		fmt.Printf("\n")
	}
}
