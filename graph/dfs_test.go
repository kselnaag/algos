package graph_test

import (
	"testing"

	// "github.com/kselnaag/algos/graph"

	"github.com/stretchr/testify/assert"
)

func TestGraphSearch(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()
	t.Run("DFS", func(t *testing.T) {

	})
	t.Run("BFS", func(t *testing.T) {

	})
}
