package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	//
}
