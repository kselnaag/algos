package graph_test

import (
	"testing"

	"github.com/kselnaag/algos/graph"
	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	assert := assert.New(t)
	t.Run("UF", func(t *testing.T) {
		uf := graph.NewUF(10)
		assert.Equal(uf.Count(), 10)
		assert.Equal(uf.Size(), 10)
		uf.Union(0, 2)
		uf.Union(4, 6)
		uf.Union(1, 3)
		uf.Union(5, 7)
		uf.Union(8, 9)
		assert.Equal(uf.Count(), 5)
		assert.Equal(uf.Size(), 10)
		arr := make([]int, 10)
		for i := 0; i < 10; i++ {
			arr[i] = uf.Find(i)
		}
		assert.Equal(arr, []int{0, 1, 0, 1, 4, 5, 4, 5, 8, 8})
		uf.Union(2, 4)
		uf.Union(3, 5)
		uf.Union(7, 8)
		assert.Equal(uf.Count(), 2)
		assert.Equal(uf.Size(), 10)
		for i := 0; i < 10; i++ {
			arr[i] = uf.Find(i)
		}
		assert.Equal(arr, []int{0, 1, 0, 1, 0, 1, 0, 1, 1, 1})
		assert.False(uf.Connected(0, 1))
		assert.False(uf.Connected(2, 3))
		assert.True(uf.Connected(1, 3))
		assert.True(uf.Connected(5, 7))
		assert.True(uf.Connected(7, 8))
		assert.True(uf.Connected(8, 9))
		assert.False(uf.Connected(9, 0))
	})
}
