package list_test

import (
	"testing"

	"github.com/kselnaag/algos/list"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("bag", func(t *testing.T) {
		bag := list.NewBag[int]()
		assert.True(bag.IsEmpty())
		assert.Equal(bag.Size(), 0)
		bag.Add(134)
		bag.Add(25)
		bag.Add(67)
		bag.Add(43)
		bag.Add(29)
		bag.Add(3)
		bag.Add(23)
		assert.False(bag.IsEmpty())
		assert.Equal(bag.Size(), 7)
		blen := bag.Size()
		arr := make([]int, 7)
		for i := 0; i < blen; i++ {
			arr[i] = bag.Next()
		}
		assert.Equal(arr, []int{23, 3, 29, 43, 67, 25, 134})
		bag.Reverse()
		for i := 0; i < blen; i++ {
			arr[i] = bag.Next()
		}
		assert.Equal(arr, []int{134, 25, 67, 43, 29, 3, 23})
		assert.False(bag.IsEmpty())
		assert.Equal(bag.Size(), 7)
		bag.Drop()
		assert.True(bag.IsEmpty())
		assert.Equal(bag.Size(), 0)
		assert.Panics(func() { bag.Next() }, "The code is not panic")
	})

	t.Run("stack", func(t *testing.T) {
		stack := list.NewStack[int]()
		assert.True(stack.IsEmpty())
		assert.Equal(stack.Size(), 0)
		stack.Push(134)
		stack.Push(25)
		stack.Push(67)
		stack.Push(43)
		stack.Push(29)
		stack.Push(3)
		stack.Push(23)
		assert.False(stack.IsEmpty())
		assert.Equal(stack.Size(), 7)
		slen := stack.Size()
		arr := make([]int, 7)
		for i := 0; i < slen; i++ {
			arr[i] = stack.Next()
		}
		assert.Equal(arr, []int{23, 3, 29, 43, 67, 25, 134})
		stack.Reverse()
		for i := 0; i < slen; i++ {
			arr[i] = stack.Pop()
		}
		assert.Equal(arr, []int{134, 25, 67, 43, 29, 3, 23})
		assert.True(stack.IsEmpty())
		assert.Equal(stack.Size(), 0)

		stack.Push(134)
		stack.Push(25)
		stack.Push(67)
		stack.Push(43)
		stack.Push(29)
		stack.Push(3)
		stack.Push(23)
		assert.False(stack.IsEmpty())
		assert.Equal(stack.Size(), 7)
		stack.Drop()
		assert.True(stack.IsEmpty())
		assert.Equal(stack.Size(), 0)
		assert.Panics(func() { stack.Next() }, "The code is not panic")
		assert.Panics(func() { stack.Pop() }, "The code is not panics")
	})

	t.Run("queue", func(t *testing.T) {
		queue := list.NewQueue[int]()
		assert.True(queue.IsEmpty())
		assert.Equal(queue.Size(), 0)
		queue.Enq(134)
		queue.Enq(25)
		queue.Enq(67)
		queue.Enq(43)
		queue.Enq(29)
		queue.Enq(3)
		queue.Enq(23)
		assert.False(queue.IsEmpty())
		assert.Equal(queue.Size(), 7)
		qlen := queue.Size()
		arr := make([]int, 7)
		for i := 0; i < qlen; i++ {
			arr[i] = queue.Next()
		}
		assert.Equal(arr, []int{134, 25, 67, 43, 29, 3, 23})
		queue.Reverse()
		for i := 0; i < qlen; i++ {
			arr[i] = queue.Deq()
		}
		assert.Equal(arr, []int{23, 3, 29, 43, 67, 25, 134})
		assert.True(queue.IsEmpty())
		assert.Equal(queue.Size(), 0)

		queue.Enq(134)
		queue.Enq(25)
		queue.Enq(67)
		queue.Enq(43)
		queue.Enq(29)
		queue.Enq(3)
		queue.Enq(23)
		assert.False(queue.IsEmpty())
		assert.Equal(queue.Size(), 7)
		queue.Drop()
		assert.True(queue.IsEmpty())
		assert.Equal(queue.Size(), 0)
		assert.Panics(func() { queue.Next() }, "The code is not panic")
		assert.Panics(func() { queue.Deq() }, "The code is not panic")
	})
}
