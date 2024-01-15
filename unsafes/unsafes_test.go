package unsafes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsafes(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("PrintAsBinary", func(t *testing.T) {
		/*
			var x uint16 = 32768
			asrt.Equal("00000000 10000000 ", U.PrintAsBinary(x))
			var y uint64 = math.MaxUint64
			asrt.Equal("11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111 ", U.PrintAsBinary(y))
			var z uint8 = 1
			asrt.Equal("00000001 ", U.PrintAsBinary(z))
		*/
	})
}
