package math_test

import (
	"testing"

	"github.com/kselnaag/algos/math"
	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("Harmonic", func(t *testing.T) {
		assert.Panics(func() { math.Harmonic(-1) }, "The code is not panic")
		assert.Panics(func() { math.Harmonic(-265) }, "The code is not panic")
		assert.Equal(math.Harmonic(0), 0.0)
		assert.Equal(math.Harmonic(1), 1.0)
		assert.Equal(math.Harmonic(4), 2.083333333333333)
		assert.Equal(math.Harmonic(256), 6.124344962817281)
	})
	t.Run("Sqrt", func(t *testing.T) {
		assert.Panics(func() { math.Sqrt(-4) }, "The code is not panic")
		assert.Panics(func() { math.Sqrt(-147) }, "The code is not panic")
		assert.Equal(math.Sqrt(0), 0.0)
		assert.Equal(math.Sqrt(4), 2.0)
		assert.Equal(math.Sqrt(121), 11.0)
		assert.Equal(math.Sqrt(734449), 857.0)
	})
	t.Run("Abs", func(t *testing.T) {
		assert.Equal(math.Abs(0), 0)
		assert.Equal(math.Abs(0.0), 0.0)
		assert.Equal(math.Abs(-2), 2)
		assert.Equal(math.Abs(-5.0), 5.0)
		assert.Equal(math.Abs(467), 467)
		assert.Equal(math.Abs(975.0), 975.0)
		assert.Equal(math.Abs(-467), 467)
		assert.Equal(math.Abs(-975.0), 975.0)

	})
	t.Run("IsPrime", func(t *testing.T) {
		assert.False(math.IsPrime(-1))
		assert.False(math.IsPrime(0))
		assert.False(math.IsPrime(1))
		assert.True(math.IsPrime(2))
		assert.True(math.IsPrime(3))
		assert.True(math.IsPrime(953))
		assert.True(math.IsPrime(1117))
		assert.False(math.IsPrime(2000))
		assert.False(math.IsPrime(4444))
	})
	t.Run("Gcd", func(t *testing.T) {
		assert.Equal(math.Gcd(123, 0), 123)
		assert.Equal(math.Gcd(0, 17), 17)
		assert.Equal(math.Gcd(-16, -32), 16)
		assert.Equal(math.Gcd(85, 51), 17)
	})
	t.Run("Min", func(t *testing.T) {
		assert.Equal(math.Min(1, 2), 1)
		assert.Equal(math.Min(2, 1), 1)
		assert.Equal(math.Min(2, 0), 0)
		assert.Equal(math.Min(0, 2), 0)
		assert.Equal(math.Min(0, 0), 0)
		assert.Equal(math.Min(-14, 6), -14)
		assert.Equal(math.Min(-14, -24), -24)
		assert.Equal(math.Min(4.093, 16.77), 4.093)
		assert.Equal(math.Min("aac", "aa"), "aa")

	})
	t.Run("Max", func(t *testing.T) {
		assert.Equal(math.Max(1, 2), 2)
		assert.Equal(math.Max(2, 1), 2)
		assert.Equal(math.Max(2, 0), 2)
		assert.Equal(math.Max(0, 2), 2)
		assert.Equal(math.Max(0, 0), 0)
		assert.Equal(math.Max(-14, 6), 6)
		assert.Equal(math.Max(-14, -24), -14)
		assert.Equal(math.Max(4.093, 16.77), 16.77)
		assert.Equal(math.Max("aac", "aa"), "aac")
	})
}
