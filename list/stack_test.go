package list_test

import (
	"testing"

	"algos/list"

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
		assert.Equal(bag.Iterate(), []int{23, 3, 29, 43, 67, 25, 134})
		bag.Reverse()
		assert.Equal(bag.Iterate(), []int{134, 25, 67, 43, 29, 3, 23})
		assert.False(bag.IsEmpty())
		assert.Equal(bag.Size(), 7)
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
		assert.Equal(stack.Iterate(), []int{23, 3, 29, 43, 67, 25, 134})
		stack.Reverse()
		slen := stack.Size()
		arr := make([]int, 7)
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
		stack = list.NewStack[int]()
		assert.Panics(func() { stack.Pop() }, "algos.list.(Queue).Deq():  the code is not panic when structure is empty")
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
		assert.Equal(queue.Iterate(), []int{134, 25, 67, 43, 29, 3, 23})
		queue.Reverse()
		qlen := queue.Size()
		arr := make([]int, 7)
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
		queue = list.NewQueue[int]()
		assert.Panics(func() { queue.Deq() }, "algos.list.(Queue).Deq():  the code is not panic when structure is empty")
	})
}
