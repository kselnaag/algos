package array_test

import (
	"testing"

	"github.com/kselnaag/algos/array"
	I "github.com/kselnaag/algos/types"
	"github.com/stretchr/testify/assert"
)

type myType struct {
	a int
	b int
}

func (s myType) CompareTo(st I.Comp) int {
	if (s.a + s.b) < ((st.(myType)).a + (st.(myType)).b) {
		return -1
	}
	if (s.a + s.b) > ((st.(myType)).a + (st.(myType)).b) {
		return +1
	}
	return 0
}

func TestSortComp(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("structs", func(t *testing.T) {
		s0 := &myType{}
		s1 := &myType{1, 2}
		s2 := &myType{3, 1}
		s3 := &myType{2, 3}
		s4 := &myType{3, 4}
		s5 := &myType{6, 7}
		s6 := &myType{39, 17}
		s7 := &myType{45, 61}
		s8 := &myType{-39, 17}

		tests := []struct {
			args []*myType
			rets []*myType
		}{
			{[]*myType{}, []*myType{}},
			{[]*myType{s2}, []*myType{s2}},
			{[]*myType{s2, s1}, []*myType{s1, s2}},
			{[]*myType{s1, s2, s3}, []*myType{s1, s2, s3}},
			{[]*myType{s3, s2, s1}, []*myType{s1, s2, s3}},
			{[]*myType{s3, s1, s2}, []*myType{s1, s2, s3}},
			{[]*myType{s2, s3, s1}, []*myType{s1, s2, s3}},
			{[]*myType{s0, s8, s3}, []*myType{s8, s0, s3}},
			{[]*myType{s6, s2, s1, s4, s5, s3}, []*myType{s1, s2, s3, s4, s5, s6}},
			{[]*myType{s7, s3, s6, s5, s4, s1, s2}, []*myType{s1, s2, s3, s4, s5, s6, s7}},
		}
		t.Run("InsertSortComp", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSortComp(test.args)
				assert.Equal(test.args, test.rets)
			}
		})
		t.Run("SelectSortComp", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSortComp(test.args)
				assert.Equal(test.args, test.rets)
			}
		})
		t.Run("ReverseSortComp", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSortComp(test.args)
				assert.Equal(test.args, test.rets)
			}
		})

	})
}

func TestSortOrd(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("utils", func(t *testing.T) {
		assert.Equal(array.ContainedOrd([]int{32, 45, 26}, 45), 1)
		assert.Equal(array.ContainedOrd([]int{32, 45, 26}, 66), -1)
		assert.Equal(array.ContainedOrd([]int{}, 66), -1)

		assert.Equal(array.BinarySearchOrd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8), 7)
		assert.Equal(array.BinarySearchOrd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55), -1)
		assert.Equal(array.BinarySearchOrd([]int{}, 55), -1)
	})

	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			args []int
			rets []int
		}{
			{[]int{}, []int{}},
			{[]int{437}, []int{437}},
			{[]int{2, 1}, []int{1, 2}},
			{[]int{1, 2, 3}, []int{1, 2, 3}},
			{[]int{3, 2, 1}, []int{1, 2, 3}},
			{[]int{3, 1, 2}, []int{1, 2, 3}},
			{[]int{2, 3, 1}, []int{1, 2, 3}},
			{[]int{0, -2, 3}, []int{-2, 0, 3}},
			{[]int{24, 67, 54, 32, 87, 65}, []int{24, 32, 54, 65, 67, 87}},
			{[]int{134, 25, 67, 43, 29, 03, 23}, []int{3, 23, 25, 29, 43, 67, 134}},
			{[]int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
		}
		t.Run("InsertSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ShellSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("SelectSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ReverseSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("MergeSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("QuickSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("HeapSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			args []float64
			rets []float64
		}{
			{[]float64{}, []float64{}},
			{[]float64{437}, []float64{437}},
			{[]float64{2, 1}, []float64{1, 2}},
			{[]float64{1, 2, 3}, []float64{1, 2, 3}},
			{[]float64{3, 2, 1}, []float64{1, 2, 3}},
			{[]float64{3, 1, 2}, []float64{1, 2, 3}},
			{[]float64{2, 3, 1}, []float64{1, 2, 3}},
			{[]float64{0, -2, 3}, []float64{-2, 0, 3}},
			{[]float64{24, 67, 54, 32, 87, 65}, []float64{24, 32, 54, 65, 67, 87}},
			{[]float64{134, 25, 67, 43, 29, 03, 23}, []float64{3, 23, 25, 29, 43, 67, 134}},
			{[]float64{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
		}
		t.Run("InsertSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ShellSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("SelectSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ReverseSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("MergeSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("QuickSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("HeapSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
	})

	t.Run("strings", func(t *testing.T) {
		tests := []struct {
			args []string
			rets []string
		}{
			{[]string{""}, []string{""}},
			{[]string{"hello"}, []string{"hello"}},
			{[]string{"ab", "aa"}, []string{"aa", "ab"}},
			{[]string{"aa", "ab", "ac"}, []string{"aa", "ab", "ac"}},
			{[]string{"ac", "ab", "aa"}, []string{"aa", "ab", "ac"}},
			{[]string{"ac", "aa", "ab"}, []string{"aa", "ab", "ac"}},
			{[]string{"ab", "ac", "aa"}, []string{"aa", "ab", "ac"}},
			{[]string{"a", "", "ac"}, []string{"", "a", "ac"}},
			{[]string{"ac", "ae", "ab", "ad", "aa", "af"}, []string{"aa", "ab", "ac", "ad", "ae", "af"}},
			{[]string{"ac", "ae", "ab", "aj", "ad", "aa", "af"}, []string{"aa", "ab", "ac", "ad", "ae", "af", "aj"}},
			{[]string{"az", "ay", "ax", "aw", "av", "au", "at", "as", "ar", "aq", "ap", "ao", "an", "am", "al", "ak", "aj", "ai", "ah", "ag", "af", "ae", "ad", "ac", "ab", "aa"},
				[]string{"aa", "ab", "ac", "ad", "ae", "af", "ag", "ah", "ai", "aj", "ak", "al", "am", "an", "ao", "ap", "aq", "ar", "as", "at", "au", "av", "aw", "ax", "ay", "az"}},
		}
		t.Run("InsertSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ShellSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("SelectSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("ReverseSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("MergeSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("QuickSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
		t.Run("HeapSortOrd", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSortOrd(test.args)
				assert.Equal(test.args, test.rets)
				assert.True(array.IsSortedOrd(test.args))
			}
		})
	})
}
