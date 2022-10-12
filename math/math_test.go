package math_test

import (
	"testing"

	amath "github.com/kselnaag/algos/math"
	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("ConvToByteArr", func(t *testing.T) {
		assert.Equal(amath.ConvToByteArr(int(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr(int(1)), []byte{0, 0, 0, 0, 0, 0, 0, 1})
		assert.Equal(amath.ConvToByteArr(int(1<<40)), []byte{0, 0, 1, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr(uint(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr(uint(3)), []byte{0, 0, 0, 0, 0, 0, 0, 3})
		assert.Equal(amath.ConvToByteArr(uint(1<<48)), []byte{0, 1, 0, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr(float64(0)), []byte{0, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr(float64(5)), []byte{0, 0, 0, 0, 0, 0, 0, 5})
		assert.Equal(amath.ConvToByteArr(float64(1<<56)), []byte{1, 0, 0, 0, 0, 0, 0, 0})
		assert.Equal(amath.ConvToByteArr("abcdefgh"), []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68})
		assert.Panics(func() { amath.ConvToByteArr(uint8(0)) }, "The code is not panic")
	})
	t.Run("HashPirson32", func(t *testing.T) {
		assert.Equal(amath.HashPirson16([]byte{0, 0, 0, 0, 0, 0, 0, 0}), uint16(0x677c))
		assert.Equal(amath.HashPirson16([]byte{0, 0, 0, 0, 0, 0, 0, 1}), uint16(0xa2f0))
		assert.Equal(amath.HashPirson16([]byte{1, 0, 0, 0, 0, 0, 0, 0}), uint16(0xcb82))
	})
	t.Run("Harmonic", func(t *testing.T) {
		assert.Panics(func() { amath.Harmonic(-1) }, "The code is not panic")
		assert.Panics(func() { amath.Harmonic(-265) }, "The code is not panic")
		assert.Equal(amath.Harmonic(0), 0.0)
		assert.Equal(amath.Harmonic(1), 1.0)
		assert.Equal(amath.Harmonic(4), 2.083333333333333)
		assert.Equal(amath.Harmonic(256), 6.124344962817281)
	})
	t.Run("Sqrt", func(t *testing.T) {
		assert.Panics(func() { amath.Sqrt(-4) }, "The code is not panic")
		assert.Panics(func() { amath.Sqrt(-147) }, "The code is not panic")
		assert.Equal(amath.Sqrt(0), 0.0)
		assert.Equal(amath.Sqrt(4), 2.0)
		assert.Equal(amath.Sqrt(121), 11.0)
		assert.Equal(amath.Sqrt(734449), 857.0)
	})
	t.Run("Abs", func(t *testing.T) {
		assert.Equal(amath.Abs(0), 0)
		assert.Equal(amath.Abs(0.0), 0.0)
		assert.Equal(amath.Abs(-2), 2)
		assert.Equal(amath.Abs(-5.0), 5.0)
		assert.Equal(amath.Abs(467), 467)
		assert.Equal(amath.Abs(975.0), 975.0)
		assert.Equal(amath.Abs(-467), 467)
		assert.Equal(amath.Abs(-975.0), 975.0)

	})
	t.Run("IsPrime", func(t *testing.T) {
		assert.False(amath.IsPrime(-1))
		assert.False(amath.IsPrime(0))
		assert.False(amath.IsPrime(1))
		assert.True(amath.IsPrime(2))
		assert.True(amath.IsPrime(3))
		assert.True(amath.IsPrime(953))
		assert.True(amath.IsPrime(1117))
		assert.False(amath.IsPrime(2000))
		assert.False(amath.IsPrime(4444))
	})
	t.Run("Gcd", func(t *testing.T) {
		assert.Equal(amath.Gcd(123, 0), 123)
		assert.Equal(amath.Gcd(0, 17), 17)
		assert.Equal(amath.Gcd(-16, -32), 16)
		assert.Equal(amath.Gcd(85, 51), 17)
	})
	t.Run("Min", func(t *testing.T) {
		assert.Equal(amath.Min(1, 2), 1)
		assert.Equal(amath.Min(2, 1), 1)
		assert.Equal(amath.Min(2, 0), 0)
		assert.Equal(amath.Min(0, 2), 0)
		assert.Equal(amath.Min(0, 0), 0)
		assert.Equal(amath.Min(-14, 6), -14)
		assert.Equal(amath.Min(-14, -24), -24)
		assert.Equal(amath.Min(4.093, 16.77), 4.093)
		assert.Equal(amath.Min("aac", "aa"), "aa")

	})
	t.Run("Max", func(t *testing.T) {
		assert.Equal(amath.Max(1, 2), 2)
		assert.Equal(amath.Max(2, 1), 2)
		assert.Equal(amath.Max(2, 0), 2)
		assert.Equal(amath.Max(0, 2), 2)
		assert.Equal(amath.Max(0, 0), 0)
		assert.Equal(amath.Max(-14, 6), 6)
		assert.Equal(amath.Max(-14, -24), -14)
		assert.Equal(amath.Max(4.093, 16.77), 16.77)
		assert.Equal(amath.Max("aac", "aa"), "aac")
	})
}
