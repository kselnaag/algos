package amath_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"algos/amath"
)

func TestMath(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	eps := 0.000001

	t.Run("HashELF", func(t *testing.T) {
		asrt.Equal(uint32(0x00000048), amath.HashELF([]byte("H")))
		asrt.Equal(uint32(0x04a85151), amath.HashELF([]byte("HelloWorld!")))
		asrt.Equal(uint32(0x034a8514), amath.HashELF([]byte("HelloWorld")))
	})
	t.Run("HashDJB2a", func(t *testing.T) {
		asrt.Equal(uint32(0x0002b5ed), amath.HashDJB2[uint32]([]byte("H")))
		asrt.Equal(uint32(0xfe9502c4), amath.HashDJB2a[uint32]([]byte("HelloWorld!")))
		asrt.Equal(uint32(0xa2ddba45), amath.HashDJB2a[uint32]([]byte("HelloWorld")))

		asrt.Equal(uint64(0x000000000002b5ed), amath.HashDJB2[uint64]([]byte("H")))
		asrt.Equal(uint64(0xbfe37049fe9502c4), amath.HashDJB2a[uint64]([]byte("HelloWorld!")))
		asrt.Equal(uint64(0x726bbd95a2ddba45), amath.HashDJB2a[uint64]([]byte("HelloWorld")))
	})
	t.Run("HashDJB2", func(t *testing.T) {
		asrt.Equal(uint32(0x0002b5ed), amath.HashDJB2[uint32]([]byte("H")))
		asrt.Equal(uint32(0x0977a182), amath.HashDJB2[uint32]([]byte("HelloWorld!")))
		asrt.Equal(uint32(0x7c687941), amath.HashDJB2[uint32]([]byte("HelloWorld")))

		asrt.Equal(uint64(0x000000000002b5ed), amath.HashDJB2[uint64]([]byte("H")))
		asrt.Equal(uint64(0xbfe621040977a182), amath.HashDJB2[uint64]([]byte("HelloWorld!")))
		asrt.Equal(uint64(0x726bd2747c687941), amath.HashDJB2[uint64]([]byte("HelloWorld")))
	})
	t.Run("HashPirson", func(t *testing.T) {
		asrt.Equal(uint8(0x7c), amath.HashPirson[uint8]([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
		asrt.Equal(uint16(0xa2f0), amath.HashPirson[uint16]([]byte{0, 0, 0, 0, 0, 0, 0, 1}))
		asrt.Equal(uint32(0x9657cb82), amath.HashPirson[uint32]([]byte{1, 0, 0, 0, 0, 0, 0, 0}))
		asrt.Equal(uint64(0x1be6c14bc5b612bd), amath.HashPirson[uint64]([]byte{0, 0, 0, 1, 0, 0, 0, 0}))
	})
	t.Run("Harmonic", func(t *testing.T) {
		asrt.True(math.IsNaN(amath.Harmonic(-1)))
		asrt.True(math.IsNaN(amath.Harmonic(-265)))
		asrt.Zero(amath.Harmonic(0))
		asrt.InEpsilon(1.0, amath.Harmonic(1), eps)
		asrt.InEpsilon(2.083333333333333, amath.Harmonic(4), eps)
		asrt.InEpsilon(6.124344962817281, amath.Harmonic(256), eps)
	})
	t.Run("Sqrt", func(t *testing.T) {
		asrt.True(math.IsNaN(amath.Sqrt(-4)))
		asrt.True(math.IsNaN(amath.Sqrt(-147)))
		asrt.Zero(amath.Sqrt(0))
		asrt.InEpsilon(2.0, amath.Sqrt(4), eps)
		asrt.InEpsilon(11.0, amath.Sqrt(121), eps)
		asrt.InEpsilon(857.0, amath.Sqrt(734449), eps)
	})
	t.Run("Abs", func(t *testing.T) {
		asrt.Zero(amath.Abs(0))
		asrt.Zero(amath.Abs(0.0))
		asrt.Equal(2, amath.Abs(-2))
		asrt.InEpsilon(5.0, amath.Abs(-5.0), eps)
		asrt.Equal(467, amath.Abs(467))
		asrt.InEpsilon(975.0, amath.Abs(975.0), eps)
		asrt.Equal(467, amath.Abs(-467))
		asrt.InEpsilon(975.0, amath.Abs(-975.0), eps)
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
		asrt.InEpsilon(4.093, amath.Min(4.093, 16.77), eps)
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
		asrt.InEpsilon(16.77, amath.Max(4.093, 16.77), eps)
		asrt.Equal("aac", amath.Max("aac", "aa"))
	})
	t.Run("Ternar", func(t *testing.T) {
		asrt.False(amath.Ternar(true, false, true))
		asrt.True(amath.Ternar(false, false, true))

		asrt.Equal(1, amath.Ternar(true, 1, 2))
		asrt.Equal(2, amath.Ternar(false, 1, 2))

		asrt.InEpsilon(2.0, amath.Ternar(true, 2.0, -1.1), eps)
		asrt.InEpsilon(-1.1, amath.Ternar(false, 2.0, -1.1), eps)

		asrt.Equal("aaa", amath.Ternar(true, "aaa", "bbb"))
		asrt.Equal("bbb", amath.Ternar(false, "aaa", "bbb"))
	})
}
