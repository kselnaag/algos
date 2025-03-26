package array_test

import (
	"testing"

	"algos/array"

	"github.com/stretchr/testify/assert"
)

func TestSvc(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("RingBuff", func(t *testing.T) {
		rb := array.NewRingBuff[byte](10)
		asrt.Equal(10, rb.Size())
		asrt.Equal(0, rb.DataLen())

		rb.Set(1)
		rb.Set(2)
		rb.Set(3)
		rb.Set(4)
		rb.Set(5)
		rb.Set(6)
		asrt.Equal(6, rb.DataLen())
		asrt.Equal("[1 2 3 4 5 6 0 0 0 0]", rb.ToString())
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		b := rb.Get()
		asrt.Equal(b, byte(6))
		asrt.Equal(0, rb.DataLen())

		rb.Set(1)
		rb.Set(2)
		rb.Set(3)
		rb.Set(4)
		rb.Set(5)
		rb.Set(6)
		asrt.Equal(6, rb.DataLen())
		asrt.Equal("[5 6 3 4 5 6 1 2 3 4]", rb.ToString())
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		_ = rb.Get()
		b = rb.Get()
		asrt.Equal(b, byte(6))
		asrt.Equal(0, rb.DataLen())
	})
}
