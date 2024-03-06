package types_test

import (
	"testing"

	I "github.com/kselnaag/algos/types"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	s1 := &I.TestStruct{A: 1, B: 2}
	s2 := &I.TestStruct{A: 3, B: 1}
	s3 := &I.TestStruct{A: 2, B: 3}

	t.Run("ConvToBytes", func(t *testing.T) {
		type STtest struct {
			A int
			B float64
			C string
		}
		asrt.Equal([]byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}, I.ConvToBytes(123456789))
		asrt.Equal([]byte{0x39, 0x38, 0x37, 0x36, 0x35, 0x2e, 0x34, 0x33, 0x32, 0x31}, I.ConvToBytes(98765.4321))
		asrt.Equal([]byte{0x7b, 0x32, 0x33, 0x34, 0x35, 0x20, 0x31, 0x36, 0x37, 0x38, 0x2e, 0x32, 0x33, 0x34, 0x35, 0x20, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21, 0x7d},
			I.ConvToBytes(STtest{2345, 1678.2345, "Hello, world!"}))
		asrt.Equal([]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21}, I.ConvToBytes("Hello, world!"))
	})
	t.Run("LT", func(t *testing.T) {
		asrt.True(I.LT(int(-1), int(16)))
		asrt.False(I.LT(int(14), int(0)))

		asrt.True(I.LT(uint(0), uint(26)))
		asrt.False(I.LT(uint(5), uint(0)))

		asrt.True(I.LT(uint32(0), uint32(36)))
		asrt.False(I.LT(uint32(5), uint32(0)))

		asrt.True(I.LT(float64(15.0), float64(16.0)))
		asrt.False(I.LT(float64(5.0), float64(-1.0)))

		asrt.True(I.LT("a", "aa"))
		asrt.False(I.LT("bb", "b"))

		asrt.True(I.LT(s1, s2))
		asrt.False(I.LT(s3, s2))
	})
	t.Run("GT", func(t *testing.T) {
		asrt.True(I.GT(int(16), int(-1)))
		asrt.False(I.GT(int(0), int(14)))

		asrt.True(I.GT(uint(26), uint(0)))
		asrt.False(I.GT(uint(0), uint(5)))

		asrt.True(I.GT(uint32(36), uint32(0)))
		asrt.False(I.GT(uint32(0), uint32(5)))

		asrt.True(I.GT(float64(16.0), float64(14.0)))
		asrt.False(I.GT(float64(-1.0), float64(5.0)))

		asrt.True(I.GT("aa", "a"))
		asrt.False(I.GT("b", "cc"))

		asrt.True(I.GT(s2, s1))
		asrt.False(I.GT(s2, s3))
	})
	t.Run("EQ", func(t *testing.T) {
		asrt.True(I.EQ(int(16), int(16)))
		asrt.False(I.EQ(int(0), int(14)))

		asrt.True(I.EQ(uint(26), uint(26)))
		asrt.False(I.EQ(uint(0), uint(5)))

		asrt.True(I.EQ(uint32(0), uint32(0)))
		asrt.False(I.EQ(uint32(0), uint32(5)))

		asrt.True(I.EQ(float64(16.0), float64(16.0)))
		asrt.False(I.EQ(float64(-1.0), float64(5.0)))

		asrt.True(I.EQ("aa", "aa"))
		asrt.False(I.EQ("b", "cc"))

		asrt.True(I.EQ(s1, s1))
		asrt.False(I.EQ(s2, s3))
	})
}
