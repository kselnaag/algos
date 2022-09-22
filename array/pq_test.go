package array_test

import (
	"testing"

	"github.com/kselnaag/algos/array"
	"github.com/stretchr/testify/assert"
)

func TestPQOrd(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("minPQ", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			minPQ := array.NewMinPQOrd[int]()
			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)
			minPQ.Add(134)
			minPQ.Add(25)
			minPQ.Add(67)
			minPQ.Add(43)
			minPQ.Add(29)
			minPQ.Add(3)
			minPQ.Add(23)
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 7)

			assert.Equal(minPQ.Min(), 3)
			assert.Equal(minPQ.GetMin(), 3)

			assert.Equal(minPQ.Min(), 23)
			assert.Equal(minPQ.GetMin(), 23)
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 5)

			assert.Equal(minPQ.GetMin(), 25)
			assert.Equal(minPQ.GetMin(), 29)
			assert.Equal(minPQ.GetMin(), 43)
			assert.Equal(minPQ.GetMin(), 67)
			assert.Equal(minPQ.GetMin(), 134)

			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)

			assert.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			assert.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
		t.Run("floats", func(t *testing.T) {
			minPQ := array.NewMinPQOrd[float64]()
			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)
			minPQ.Add(134.0)
			minPQ.Add(25.0)
			minPQ.Add(67.0)
			minPQ.Add(43.0)
			minPQ.Add(29.0)
			minPQ.Add(3.0)
			minPQ.Add(23.0)
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 7)

			assert.Equal(minPQ.Min(), 3.0)
			assert.Equal(minPQ.GetMin(), 3.0)

			assert.Equal(minPQ.Min(), 23.0)
			assert.Equal(minPQ.GetMin(), 23.0)
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 5)

			assert.Equal(minPQ.GetMin(), 25.0)
			assert.Equal(minPQ.GetMin(), 29.0)
			assert.Equal(minPQ.GetMin(), 43.0)
			assert.Equal(minPQ.GetMin(), 67.0)
			assert.Equal(minPQ.GetMin(), 134.0)

			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)

			assert.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			assert.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
		t.Run("strings", func(t *testing.T) {
			minPQ := array.NewMinPQOrd[string]()
			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)
			minPQ.Add("ag")
			minPQ.Add("ac")
			minPQ.Add("af")
			minPQ.Add("ae")
			minPQ.Add("ad")
			minPQ.Add("aa")
			minPQ.Add("ab")
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 7)

			assert.Equal(minPQ.Min(), "aa")
			assert.Equal(minPQ.GetMin(), "aa")

			assert.Equal(minPQ.Min(), "ab")
			assert.Equal(minPQ.GetMin(), "ab")
			assert.False(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 5)

			assert.Equal(minPQ.GetMin(), "ac")
			assert.Equal(minPQ.GetMin(), "ad")
			assert.Equal(minPQ.GetMin(), "ae")
			assert.Equal(minPQ.GetMin(), "af")
			assert.Equal(minPQ.GetMin(), "ag")

			assert.True(minPQ.IsEmpty())
			assert.Equal(minPQ.Size(), 0)

			assert.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			assert.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
	})

	t.Run("maxPQ", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			maxPQ := array.NewMaxPQOrd[int]()
			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)
			maxPQ.Add(134)
			maxPQ.Add(25)
			maxPQ.Add(67)
			maxPQ.Add(43)
			maxPQ.Add(29)
			maxPQ.Add(3)
			maxPQ.Add(23)
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 7)

			assert.Equal(maxPQ.Max(), 134)
			assert.Equal(maxPQ.GetMax(), 134)

			assert.Equal(maxPQ.Max(), 67)
			assert.Equal(maxPQ.GetMax(), 67)
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 5)

			assert.Equal(maxPQ.GetMax(), 43)
			assert.Equal(maxPQ.GetMax(), 29)
			assert.Equal(maxPQ.GetMax(), 25)
			assert.Equal(maxPQ.GetMax(), 23)
			assert.Equal(maxPQ.GetMax(), 3)

			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)

			assert.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			assert.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")

		})
		t.Run("floats", func(t *testing.T) {
			maxPQ := array.NewMaxPQOrd[float64]()
			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)
			maxPQ.Add(134.0)
			maxPQ.Add(25.0)
			maxPQ.Add(67.0)
			maxPQ.Add(43.0)
			maxPQ.Add(29.0)
			maxPQ.Add(3.0)
			maxPQ.Add(23.0)
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 7)

			assert.Equal(maxPQ.Max(), 134.0)
			assert.Equal(maxPQ.GetMax(), 134.0)

			assert.Equal(maxPQ.Max(), 67.0)
			assert.Equal(maxPQ.GetMax(), 67.0)
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 5)

			assert.Equal(maxPQ.GetMax(), 43.0)
			assert.Equal(maxPQ.GetMax(), 29.0)
			assert.Equal(maxPQ.GetMax(), 25.0)
			assert.Equal(maxPQ.GetMax(), 23.0)
			assert.Equal(maxPQ.GetMax(), 3.0)

			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)

			assert.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			assert.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")
		})
		t.Run("strings", func(t *testing.T) {
			maxPQ := array.NewMaxPQOrd[string]()
			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)
			maxPQ.Add("ag")
			maxPQ.Add("ac")
			maxPQ.Add("af")
			maxPQ.Add("ae")
			maxPQ.Add("ad")
			maxPQ.Add("aa")
			maxPQ.Add("ab")
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 7)

			assert.Equal(maxPQ.Max(), "ag")
			assert.Equal(maxPQ.GetMax(), "ag")

			assert.Equal(maxPQ.Max(), "af")
			assert.Equal(maxPQ.GetMax(), "af")
			assert.False(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 5)

			assert.Equal(maxPQ.GetMax(), "ae")
			assert.Equal(maxPQ.GetMax(), "ad")
			assert.Equal(maxPQ.GetMax(), "ac")
			assert.Equal(maxPQ.GetMax(), "ab")
			assert.Equal(maxPQ.GetMax(), "aa")

			assert.True(maxPQ.IsEmpty())
			assert.Equal(maxPQ.Size(), 0)

			assert.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			assert.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")
		})
	})
}
