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

	t.Run("ConvToByteArr", func(t *testing.T) {
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0}, I.ConvToByteArr(int(0)))
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 1}, I.ConvToByteArr(int(1)))
		asrt.Equal([]byte{0, 0, 1, 0, 0, 0, 0, 0}, I.ConvToByteArr(int(1<<40)))
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0}, I.ConvToByteArr(uint(0)))
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 3}, I.ConvToByteArr(uint(3)))
		asrt.Equal([]byte{0, 1, 0, 0, 0, 0, 0, 0}, I.ConvToByteArr(uint(1<<48)))
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0}, I.ConvToByteArr(float64(0)))
		asrt.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 5}, I.ConvToByteArr(float64(5)))
		asrt.Equal([]byte{1, 0, 0, 0, 0, 0, 0, 0}, I.ConvToByteArr(float64(1<<56)))
		asrt.Equal([]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68}, I.ConvToByteArr("abcdefgh"))
		asrt.Panics(func() { I.ConvToByteArr(uint8(0)) }, "algos.types.ConvToByteArr():  Is not panics when args Type is not processed")
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
