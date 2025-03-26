package array_test

import (
	"testing"

	"algos/array"
	I "algos/types"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) { //nolint:gocognit // everything normal
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("utils", func(t *testing.T) {
		t.Run("Contained", func(t *testing.T) {
			asrt.Equal(2, array.Contained([]int{32, 45, 26}, 26))
			asrt.Equal(-1, array.Contained([]int{32, 45, 26}, 66))
			asrt.Equal(-1, array.Contained([]int{}, 48))
		})
		t.Run("BinarySearch", func(t *testing.T) {
			asrt.Equal(7, array.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))
			asrt.Equal(-1, array.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55))
			asrt.Equal(-1, array.BinarySearch([]int{}, 55))
		})
		t.Run("IsSorted", func(t *testing.T) {
			asrt.True(array.IsSorted([]int{24, 32, 54, 65, 67, 87}))
			asrt.True(array.IsSorted([]int{3, 23, 25, 29, 43, 67, 134}))
			asrt.False(array.IsSorted([]int{24, 67, 54, 32, 87, 65}))
			asrt.False(array.IsSorted([]int{134, 25, 67, 43, 29, 3, 23}))
		})
		t.Run("Reverse", func(t *testing.T) {
			arr := []int{24, 67, 54, 32, 87, 65}
			array.Reverse(arr)
			asrt.Equal([]int{65, 87, 32, 54, 67, 24}, arr)

			arr = []int{134, 25, 67, 43, 29, 3, 23}
			array.Reverse(arr)
			asrt.Equal([]int{23, 3, 29, 43, 67, 25, 134}, arr)

			arr = []int{24, 67, 54, 32, 87, 65}
			array.Reverse(arr)
			asrt.NotEqual([]int{24, 67, 54, 32, 87, 65}, arr)

			arr = []int{24, 67, 54, 32, 87, 65}
			brr := []int{24, 67, 54, 32, 87, 65}
			array.Reverse(arr)
			array.Shuffle(brr)
			asrt.NotEqual(arr, brr)

			arr = []int{134, 25, 67, 43, 29, 3, 23}
			array.Reverse(arr)
			asrt.NotEqual([]int{134, 25, 67, 43, 29, 3, 23}, arr)

			arr = []int{134, 25, 67, 43, 29, 3, 23}
			brr = []int{134, 25, 67, 43, 29, 3, 23}
			array.Reverse(arr)
			array.Shuffle(brr)
			asrt.NotEqual(arr, brr)
		})
	})

	t.Run("structs", func(t *testing.T) {
		s0 := &I.TestStruct{}
		s1 := &I.TestStruct{A: 1, B: 2}
		s2 := &I.TestStruct{A: 3, B: 1}
		s3 := &I.TestStruct{A: 2, B: 3}
		s4 := &I.TestStruct{A: 3, B: 4}
		s5 := &I.TestStruct{A: 6, B: 7}
		s6 := &I.TestStruct{A: 39, B: 17}
		s7 := &I.TestStruct{A: 45, B: 61}
		s8 := &I.TestStruct{A: -39, B: 17}

		tests := []struct {
			args []*I.TestStruct
			rets []*I.TestStruct
		}{
			{[]*I.TestStruct{}, []*I.TestStruct{}},
			{[]*I.TestStruct{s2}, []*I.TestStruct{s2}},
			{[]*I.TestStruct{s2, s1}, []*I.TestStruct{s1, s2}},
			{[]*I.TestStruct{s1, s2, s3}, []*I.TestStruct{s1, s2, s3}},
			{[]*I.TestStruct{s3, s2, s1}, []*I.TestStruct{s1, s2, s3}},
			{[]*I.TestStruct{s3, s1, s2}, []*I.TestStruct{s1, s2, s3}},
			{[]*I.TestStruct{s2, s3, s1}, []*I.TestStruct{s1, s2, s3}},
			{[]*I.TestStruct{s0, s8, s3}, []*I.TestStruct{s8, s0, s3}},
			{[]*I.TestStruct{s6, s2, s1, s4, s5, s3}, []*I.TestStruct{s1, s2, s3, s4, s5, s6}},
			{[]*I.TestStruct{s7, s3, s6, s5, s4, s1, s2}, []*I.TestStruct{s1, s2, s3, s4, s5, s6, s7}},
		}
		t.Run("InsertSort", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ShellSort", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("SelectSort", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ReverseSort", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("MergeSort", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("QuickSort", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("HeapSort", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("Quick3Sort", func(t *testing.T) {
			for _, test := range tests {
				array.Quick3Sort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
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
			{[]int{134, 25, 67, 43, 29, 3, 23}, []int{3, 23, 25, 29, 43, 67, 134}},
			{[]int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
		}
		t.Run("InsertSort", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ShellSort", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("SelectSort", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ReverseSort", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("MergeSort", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("QuickSort", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("HeapSort", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("Quick3Sort", func(t *testing.T) {
			for _, test := range tests {
				array.Quick3Sort(test.args)
				asrt.Equal(test.rets, test.args)
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
			{[]float64{134, 25, 67, 43, 29, 3, 23}, []float64{3, 23, 25, 29, 43, 67, 134}},
			{[]float64{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
		}
		t.Run("InsertSort", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ShellSort", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("SelectSort", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ReverseSort", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("MergeSort", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("QuickSort", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("HeapSort", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("Quick3Sort", func(t *testing.T) {
			for _, test := range tests {
				array.Quick3Sort(test.args)
				asrt.Equal(test.rets, test.args)
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
		t.Run("InsertSort", func(t *testing.T) {
			for _, test := range tests {
				array.InsertSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ShellSort", func(t *testing.T) {
			for _, test := range tests {
				array.ShellSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("SelectSort", func(t *testing.T) {
			for _, test := range tests {
				array.SelectSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("ReverseSort", func(t *testing.T) {
			for _, test := range tests {
				array.ReverseSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("MergeSort", func(t *testing.T) {
			for _, test := range tests {
				array.MergeSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("QuickSort", func(t *testing.T) {
			for _, test := range tests {
				array.QuickSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("HeapSort", func(t *testing.T) {
			for _, test := range tests {
				array.HeapSort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
		t.Run("Quick3Sort", func(t *testing.T) {
			for _, test := range tests {
				array.Quick3Sort(test.args)
				asrt.Equal(test.rets, test.args)
			}
		})
	})
}
