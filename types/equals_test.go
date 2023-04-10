package types_test

import (
	"testing"

	I "algos/types"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	s1 := &I.TestStruct{A: 1, B: 2}
	s2 := &I.TestStruct{A: 3, B: 1}
	s3 := &I.TestStruct{A: 2, B: 3}

	t.Run("ConvToByteArr", func(t *testing.T) {
		assert.Equal(I.ConvToByteArr(int(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr(int(1)), []byte{0, 0, 0, 0, 0, 0, 0, 1})
		assert.Equal(I.ConvToByteArr(int(1<<40)), []byte{0, 0, 1, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr(uint(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr(uint(3)), []byte{0, 0, 0, 0, 0, 0, 0, 3})
		assert.Equal(I.ConvToByteArr(uint(1<<48)), []byte{0, 1, 0, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr(float64(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr(float64(5)), []byte{0, 0, 0, 0, 0, 0, 0, 5})
		assert.Equal(I.ConvToByteArr(float64(1<<56)), []byte{1, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(I.ConvToByteArr("abcdefgh"), []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68})
		assert.Panics(func() { I.ConvToByteArr(uint8(0)) }, "algos.types.ConvToByteArr():  Is not panics when args Type is not processed")
	})
	t.Run("LT", func(t *testing.T) {
		assert.True(I.LT(int(-1), int(16)))
		assert.False(I.LT(int(14), int(0)))

		assert.True(I.LT(uint(0), uint(26)))
		assert.False(I.LT(uint(5), uint(0)))

		assert.True(I.LT(uint32(0), uint32(36)))
		assert.False(I.LT(uint32(5), uint32(0)))

		assert.True(I.LT(float64(15.0), float64(16.0)))
		assert.False(I.LT(float64(5.0), float64(-1.0)))

		assert.True(I.LT("a", "aa"))
		assert.False(I.LT("bb", "b"))

		assert.True(I.LT(s1, s2))
		assert.False(I.LT(s3, s2))

		assert.Panics(func() { I.LT(float32(5.0), float32(-1.0)) }, "algos.types.LT():  Is not panics when args Type is not processed")
	})
	t.Run("GT", func(t *testing.T) {
		assert.True(I.GT(int(16), int(-1)))
		assert.False(I.GT(int(0), int(14)))

		assert.True(I.GT(uint(26), uint(0)))
		assert.False(I.GT(uint(0), uint(5)))

		assert.True(I.GT(uint32(36), uint32(0)))
		assert.False(I.GT(uint32(0), uint32(5)))

		assert.True(I.GT(float64(16.0), float64(14.0)))
		assert.False(I.GT(float64(-1.0), float64(5.0)))

		assert.True(I.GT("aa", "a"))
		assert.False(I.GT("b", "cc"))

		assert.True(I.GT(s2, s1))
		assert.False(I.GT(s2, s3))

		assert.Panics(func() { I.GT(float32(5.0), float32(-1.0)) }, "algos.types.LT():  Is not panics when args Type is not processed")
	})
	t.Run("EQ", func(t *testing.T) {
		assert.True(I.EQ(int(16), int(16)))
		assert.False(I.EQ(int(0), int(14)))

		assert.True(I.EQ(uint(26), uint(26)))
		assert.False(I.EQ(uint(0), uint(5)))

		assert.True(I.EQ(uint32(0), uint32(0)))
		assert.False(I.EQ(uint32(0), uint32(5)))

		assert.True(I.EQ(float64(16.0), float64(16.0)))
		assert.False(I.EQ(float64(-1.0), float64(5.0)))

		assert.True(I.EQ("aa", "aa"))
		assert.False(I.EQ("b", "cc"))

		assert.True(I.EQ(s1, s1))
		assert.False(I.EQ(s2, s3))

		assert.Panics(func() { I.EQ(float32(5.0), float32(-1.0)) }, "algos.types.LT():  Is not panics when args Type is not processed")
	})
}
