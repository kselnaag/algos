package graph_test

import (
	"testing"

	"github.com/kselnaag/algos/graph"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()
	t.Run("UF", func(t *testing.T) {
		uf := graph.NewUF(10)
		asrt.Equal(10, uf.Count())
		asrt.Equal(10, uf.Size())
		uf.Union(0, 2)
		uf.Union(4, 6)
		uf.Union(1, 3)
		uf.Union(5, 7)
		uf.Union(8, 9)
		asrt.Equal(5, uf.Count())
		asrt.Equal(10, uf.Size())
		arr := make([]int, 10)
		for i := 0; i < 10; i++ {
			arr[i] = uf.Find(i)
		}
		asrt.Equal([]int{0, 1, 0, 1, 4, 5, 4, 5, 8, 8}, arr)
		uf.Union(2, 4)
		uf.Union(3, 5)
		uf.Union(7, 8)
		asrt.Equal(2, uf.Count())
		asrt.Equal(10, uf.Size())
		for i := 0; i < 10; i++ {
			arr[i] = uf.Find(i)
		}
		asrt.Equal([]int{0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, arr)
		asrt.False(uf.Connected(0, 1))
		asrt.False(uf.Connected(2, 3))
		asrt.True(uf.Connected(1, 3))
		asrt.True(uf.Connected(5, 7))
		asrt.True(uf.Connected(7, 8))
		asrt.True(uf.Connected(8, 9))
		asrt.False(uf.Connected(9, 0))
	})
}
