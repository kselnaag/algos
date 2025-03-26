package graph

import (
	"algos/amath"
	"algos/array"
)

type UF struct {
	id    []int
	bag   array.Bag[int]
	count int
	size  int
}

func NewUF(n int) UF {
	n = amath.Abs(n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return UF{
		id:    arr,
		bag:   array.NewBag[int](),
		count: n,
		size:  n,
	}
}

func (uf *UF) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[qRoot] = pRoot
	uf.count--
}

func (uf *UF) Find(p int) int {
	p = amath.Abs(p)
	for p != uf.id[p] {
		uf.bag.Add(p)
		p = uf.id[p]
	}
	// path compression
	blen := uf.bag.Size()
	barr := uf.bag.Iterate()
	for i := 0; i < blen; i++ {
		t := barr[i]
		uf.id[t] = p
	}
	uf.bag = array.NewBag[int]()
	return p
}

func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Size() int {
	return uf.size
}
