package list

import (
	"sync/atomic"
	"time"

	I "algos/types"
)

const maxUint8 int = (1 << 8)

type LMnode[K I.Ord, V any] struct {
	LongKey   K
	LongExpr  time.Time
	ShortKey  K
	ShortExpr time.Time
	Val       V
	Next      *LMnode[K, V]
}

type SMnode[K I.Ord, V any] struct {
	ShortKey  K
	ShortExpr time.Time
	Link      *LMnode[K, V]
	Next      *SMnode[K, V]
}

type Cache3A[K I.Ord, V any] struct {
	lmarr    [maxUint8]*LMnode[K, V]
	smarr    [maxUint8]*SMnode[K, V]
	lkeysnum atomic.Uint32
	// mu       sync.RWMutex
}

func NewCache3A[K I.Ord, V any]() *Cache3A[K, V] {
	return &Cache3A[K, V]{
		lmarr: [maxUint8]*LMnode[K, V]{},
		smarr: [maxUint8]*SMnode[K, V]{},
	}
}

func (c *Cache3A[K, V]) Size() int {
	return int(c.lkeysnum.Load())
}

func (c *Cache3A[K, V]) IsEmpty() bool {
	return (c.Size() == 0)
}

/* func hashFromKeys[K I.Ord](key K) int {
	bytesarr := I.ConvToByteArr(key)
	hash := A.HashDJB2a[uint32](bytesarr)
	return int(hash & 0x000000FF)
} */

func (c *Cache3A[K, V]) SetLong(long K, val V) {
	//
}

func (c *Cache3A[K, V]) SetShort(short, long K) {

}

func (c *Cache3A[K, V]) GetLong(long K) *V {
	return nil
}

func (c *Cache3A[K, V]) GetShort(short K) *V {
	return nil
}

func (c *Cache3A[K, V]) DelLong(long K) {

}

func (c *Cache3A[K, V]) DelShort(short K) {

}

func (c *Cache3A[K, V]) IterateLongKeys() []K {
	return []K{}
}
