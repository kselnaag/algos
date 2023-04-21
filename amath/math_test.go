package amath_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kselnaag/algos/amath"
)

func TestMath(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("HashPirson", func(t *testing.T) {
		asrt.Equal(uint8(0x7c), amath.HashPirson[uint8]([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
		asrt.Equal(uint16(0xa2f0), amath.HashPirson[uint16]([]byte{0, 0, 0, 0, 0, 0, 0, 1}))
		asrt.Equal(uint32(0x9657cb82), amath.HashPirson[uint32]([]byte{1, 0, 0, 0, 0, 0, 0, 0}))
		asrt.Equal(uint64(0x1be6c14bc5b612bd), amath.HashPirson[uint64]([]byte{0, 0, 0, 1, 0, 0, 0, 0}))
	})
	t.Run("Harmonic", func(t *testing.T) {
		asrt.Panics(func() { amath.Harmonic(-1) }, "The code is not panic")
		asrt.Panics(func() { amath.Harmonic(-265) }, "The code is not panic")
		asrt.Equal(0.0, amath.Harmonic(0))
		asrt.Equal(1.0, amath.Harmonic(1))
		asrt.Equal(2.083333333333333, amath.Harmonic(4))
		asrt.Equal(6.124344962817281, amath.Harmonic(256))
	})
	t.Run("Sqrt", func(t *testing.T) {
		asrt.Panics(func() { amath.Sqrt(-4) }, "The code is not panic")
		asrt.Panics(func() { amath.Sqrt(-147) }, "The code is not panic")
		asrt.Equal(0.0, amath.Sqrt(0))
		asrt.Equal(2.0, amath.Sqrt(4))
		asrt.Equal(11.0, amath.Sqrt(121))
		asrt.Equal(857.0, amath.Sqrt(734449))
	})
	t.Run("Abs", func(t *testing.T) {
		asrt.Equal(0, amath.Abs(0))
		asrt.Equal(0.0, amath.Abs(0.0))
		asrt.Equal(2, amath.Abs(-2))
		asrt.Equal(5.0, amath.Abs(-5.0))
		asrt.Equal(467, amath.Abs(467))
		asrt.Equal(975.0, amath.Abs(975.0))
		asrt.Equal(467, amath.Abs(-467))
		asrt.Equal(975.0, amath.Abs(-975.0))

	})
	t.Run("IsPrime", func(t *testing.T) {
		asrt.False(amath.IsPrime(-1))
		asrt.False(amath.IsPrime(0))
		asrt.False(amath.IsPrime(1))
		asrt.True(amath.IsPrime(2))
		asrt.True(amath.IsPrime(3))
		asrt.True(amath.IsPrime(953))
		asrt.True(amath.IsPrime(1117))
		asrt.False(amath.IsPrime(2000))
		asrt.False(amath.IsPrime(4444))
	})
	t.Run("Gcd", func(t *testing.T) {
		asrt.Equal(123, amath.Gcd(123, 0))
		asrt.Equal(17, amath.Gcd(0, 17))
		asrt.Equal(16, amath.Gcd(-16, -32))
		asrt.Equal(17, amath.Gcd(85, 51))
	})
	t.Run("Min", func(t *testing.T) {
		asrt.Equal(1, amath.Min(1, 2))
		asrt.Equal(1, amath.Min(2, 1))
		asrt.Equal(0, amath.Min(2, 0))
		asrt.Equal(0, amath.Min(0, 2))
		asrt.Equal(0, amath.Min(0, 0))
		asrt.Equal(-14, amath.Min(-14, 6))
		asrt.Equal(-24, amath.Min(-14, -24))
		asrt.Equal(4.093, amath.Min(4.093, 16.77))
		asrt.Equal("aa", amath.Min("aac", "aa"))
	})
	t.Run("Max", func(t *testing.T) {
		asrt.Equal(2, amath.Max(1, 2))
		asrt.Equal(2, amath.Max(2, 1))
		asrt.Equal(2, amath.Max(2, 0))
		asrt.Equal(2, amath.Max(0, 2))
		asrt.Equal(0, amath.Max(0, 0))
		asrt.Equal(6, amath.Max(-14, 6))
		asrt.Equal(-14, amath.Max(-14, -24))
		asrt.Equal(16.77, amath.Max(4.093, 16.77))
		asrt.Equal("aac", amath.Max("aac", "aa"))
	})
}
