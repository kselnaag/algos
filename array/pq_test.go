package array_test

import (
	"fmt"
	"testing"

	"algos/array"

	"github.com/stretchr/testify/assert"
)

func TestPQ(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	EPS := 0.000001

	t.Run("LRU", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			lru := array.NewLRU[int](7)
			asrt.True(lru.IsEmpty())
			asrt.Equal(0, lru.Size())
			lru.Set(134)
			lru.Set(25)
			lru.Set(67)
			lru.Set(43)
			fmt.Println(lru.Iterate())
			asrt.Equal([]int{25, 43, 67, 134}, lru.Iterate())
			asrt.False(lru.IsEmpty())
			asrt.Equal(4, lru.Size())
			lru.Set(29)
			lru.Set(3)
			lru.Set(23)
			lru.Set(86)
			lru.Set(200)
			fmt.Println(lru.Iterate())
			asrt.Equal([]int{3, 23, 25, 29, 43, 67, 86}, lru.Iterate())
			asrt.False(lru.IsEmpty())
			asrt.Equal(7, lru.Size())
		})
	})

	t.Run("MRU", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			mru := array.NewMRU[int](7)
			asrt.True(mru.IsEmpty())
			asrt.Equal(0, mru.Size())
			mru.Set(134)
			mru.Set(3)
			mru.Set(67)
			mru.Set(43)
			fmt.Println(mru.Iterate())
			asrt.Equal([]int{134, 67, 43, 3}, mru.Iterate())
			asrt.False(mru.IsEmpty())
			asrt.Equal(4, mru.Size())
			mru.Set(29)
			mru.Set(25)
			mru.Set(23)
			mru.Set(86)
			mru.Set(1)
			fmt.Println(mru.Iterate())
			asrt.Equal([]int{134, 86, 67, 43, 29, 25, 23}, mru.Iterate())
			asrt.False(mru.IsEmpty())
			asrt.Equal(7, mru.Size())
		})
	})

	t.Run("minPQ", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			minPQ := array.NewMinPQ[int]()
			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())
			minPQ.Add(134)
			minPQ.Add(25)
			minPQ.Add(67)
			minPQ.Add(43)
			minPQ.Add(29)
			minPQ.Add(3)
			minPQ.Add(23)
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			asrt.Equal(3, minPQ.Min())
			asrt.Equal(3, minPQ.GetMin())

			asrt.Equal(23, minPQ.Min())
			asrt.Equal(23, minPQ.GetMin())
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(5, minPQ.Size())

			asrt.Equal(25, minPQ.GetMin())
			asrt.Equal(29, minPQ.GetMin())
			asrt.Equal(43, minPQ.GetMin())
			asrt.Equal(67, minPQ.GetMin())
			asrt.Equal(134, minPQ.GetMin())

			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())

			minPQ.Add(134)
			minPQ.Add(25)
			minPQ.Add(67)
			minPQ.Add(43)
			minPQ.Add(29)
			minPQ.Add(3)
			minPQ.Add(23)
			asrt.Equal([]int{3, 29, 23, 134, 43, 67, 25}, minPQ.Iterate())
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			minPQ = array.NewMinPQ[int]()
			asrt.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			asrt.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
		t.Run("floats", func(t *testing.T) {
			minPQ := array.NewMinPQ[float64]()
			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())
			minPQ.Add(134.0)
			minPQ.Add(25.0)
			minPQ.Add(67.0)
			minPQ.Add(43.0)
			minPQ.Add(29.0)
			minPQ.Add(3.0)
			minPQ.Add(23.0)
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			asrt.InEpsilon(3.0, minPQ.Min(), EPS)
			asrt.InEpsilon(3.0, minPQ.GetMin(), EPS)

			asrt.InEpsilon(23.0, minPQ.Min(), EPS)
			asrt.InEpsilon(23.0, minPQ.GetMin(), EPS)
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(5, minPQ.Size())

			asrt.InEpsilon(25.0, minPQ.GetMin(), EPS)
			asrt.InEpsilon(29.0, minPQ.GetMin(), EPS)
			asrt.InEpsilon(43.0, minPQ.GetMin(), EPS)
			asrt.InEpsilon(67.0, minPQ.GetMin(), EPS)
			asrt.InEpsilon(134.0, minPQ.GetMin(), EPS)

			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())

			minPQ.Add(134.0)
			minPQ.Add(25.0)
			minPQ.Add(67.0)
			minPQ.Add(43.0)
			minPQ.Add(29.0)
			minPQ.Add(3.0)
			minPQ.Add(23.0)
			asrt.Equal([]float64{3.0, 29.0, 23.0, 134.0, 43.0, 67.0, 25.0}, minPQ.Iterate())
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			minPQ = array.NewMinPQ[float64]()
			asrt.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			asrt.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
		t.Run("strings", func(t *testing.T) {
			minPQ := array.NewMinPQ[string]()
			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())
			minPQ.Add("ag")
			minPQ.Add("ac")
			minPQ.Add("af")
			minPQ.Add("ae")
			minPQ.Add("ad")
			minPQ.Add("aa")
			minPQ.Add("ab")
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			asrt.Equal("aa", minPQ.Min())
			asrt.Equal("aa", minPQ.GetMin())

			asrt.Equal("ab", minPQ.Min())
			asrt.Equal("ab", minPQ.GetMin())
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(5, minPQ.Size())

			asrt.Equal("ac", minPQ.GetMin())
			asrt.Equal("ad", minPQ.GetMin())
			asrt.Equal("ae", minPQ.GetMin())
			asrt.Equal("af", minPQ.GetMin())
			asrt.Equal("ag", minPQ.GetMin())

			asrt.True(minPQ.IsEmpty())
			asrt.Equal(0, minPQ.Size())

			minPQ.Add("ag")
			minPQ.Add("ac")
			minPQ.Add("af")
			minPQ.Add("ae")
			minPQ.Add("ad")
			minPQ.Add("aa")
			minPQ.Add("ab")
			asrt.Equal([]string{"aa", "ad", "ab", "ag", "ae", "af", "ac"}, minPQ.Iterate())
			asrt.False(minPQ.IsEmpty())
			asrt.Equal(7, minPQ.Size())

			minPQ = array.NewMinPQ[string]()
			asrt.Panics(func() { minPQ.Min() }, "algos.array.(minPQ).Min():  the code is not panic when structure is empty")
			asrt.Panics(func() { minPQ.GetMin() }, "algos.array.(minPQ).GetMin():  the code is not panic when structure is empty")
		})
	})

	t.Run("maxPQ", func(t *testing.T) {
		t.Run("integers", func(t *testing.T) {
			maxPQ := array.NewMaxPQ[int]()
			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())
			maxPQ.Add(134)
			maxPQ.Add(25)
			maxPQ.Add(67)
			maxPQ.Add(43)
			maxPQ.Add(29)
			maxPQ.Add(3)
			maxPQ.Add(23)
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			asrt.Equal(134, maxPQ.Max())
			asrt.Equal(134, maxPQ.GetMax())

			asrt.Equal(67, maxPQ.Max())
			asrt.Equal(67, maxPQ.GetMax())
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(5, maxPQ.Size())

			asrt.Equal(43, maxPQ.GetMax())
			asrt.Equal(29, maxPQ.GetMax())
			asrt.Equal(25, maxPQ.GetMax())
			asrt.Equal(23, maxPQ.GetMax())
			asrt.Equal(3, maxPQ.GetMax())

			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())

			maxPQ.Add(134)
			maxPQ.Add(25)
			maxPQ.Add(67)
			maxPQ.Add(43)
			maxPQ.Add(29)
			maxPQ.Add(3)
			maxPQ.Add(23)
			asrt.Equal([]int{134, 43, 67, 25, 29, 3, 23}, maxPQ.Iterate())
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			maxPQ = array.NewMaxPQ[int]()
			asrt.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			asrt.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")
		})
		t.Run("floats", func(t *testing.T) {
			maxPQ := array.NewMaxPQ[float64]()
			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())
			maxPQ.Add(134.0)
			maxPQ.Add(25.0)
			maxPQ.Add(67.0)
			maxPQ.Add(43.0)
			maxPQ.Add(29.0)
			maxPQ.Add(3.0)
			maxPQ.Add(23.0)
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			eps := 0.000001
			asrt.InEpsilon(134.0, maxPQ.Max(), eps)
			asrt.InEpsilon(134.0, maxPQ.GetMax(), eps)

			asrt.InEpsilon(67.0, maxPQ.Max(), eps)
			asrt.InEpsilon(67.0, maxPQ.GetMax(), eps)
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(5, maxPQ.Size())

			asrt.InEpsilon(43.0, maxPQ.GetMax(), eps)
			asrt.InEpsilon(29.0, maxPQ.GetMax(), eps)
			asrt.InEpsilon(25.0, maxPQ.GetMax(), eps)
			asrt.InEpsilon(23.0, maxPQ.GetMax(), eps)
			asrt.InEpsilon(3.0, maxPQ.GetMax(), eps)

			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())

			maxPQ.Add(134.0)
			maxPQ.Add(25.0)
			maxPQ.Add(67.0)
			maxPQ.Add(43.0)
			maxPQ.Add(29.0)
			maxPQ.Add(3.0)
			maxPQ.Add(23.0)
			asrt.Equal([]float64{134, 43, 67, 25, 29, 3, 23}, maxPQ.Iterate())
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			maxPQ = array.NewMaxPQ[float64]()
			asrt.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			asrt.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")
		})
		t.Run("strings", func(t *testing.T) {
			maxPQ := array.NewMaxPQ[string]()
			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())
			maxPQ.Add("ag")
			maxPQ.Add("ac")
			maxPQ.Add("af")
			maxPQ.Add("ae")
			maxPQ.Add("ad")
			maxPQ.Add("aa")
			maxPQ.Add("ab")
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			asrt.Equal("ag", maxPQ.Max())
			asrt.Equal("ag", maxPQ.GetMax())

			asrt.Equal("af", maxPQ.Max())
			asrt.Equal("af", maxPQ.GetMax())
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(5, maxPQ.Size())

			asrt.Equal("ae", maxPQ.GetMax())
			asrt.Equal("ad", maxPQ.GetMax())
			asrt.Equal("ac", maxPQ.GetMax())
			asrt.Equal("ab", maxPQ.GetMax())
			asrt.Equal("aa", maxPQ.GetMax())

			asrt.True(maxPQ.IsEmpty())
			asrt.Equal(0, maxPQ.Size())

			maxPQ.Add("ag")
			maxPQ.Add("ac")
			maxPQ.Add("af")
			maxPQ.Add("ae")
			maxPQ.Add("ad")
			maxPQ.Add("aa")
			maxPQ.Add("ab")
			asrt.Equal([]string{"ag", "ae", "af", "ac", "ad", "aa", "ab"}, maxPQ.Iterate())
			asrt.False(maxPQ.IsEmpty())
			asrt.Equal(7, maxPQ.Size())

			maxPQ = array.NewMaxPQ[string]()
			asrt.Panics(func() { maxPQ.Max() }, "algos.array.(maxPQ).Max():  the code is not panic when structure is empty")
			asrt.Panics(func() { maxPQ.GetMax() }, "algos.array.(maxPQ).GetMax():  the code is not panic when structure is empty")
		})
	})
}
