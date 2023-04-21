package array_test

import (
	"testing"

	"github.com/kselnaag/algos/array"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("bag", func(t *testing.T) {
		bag := array.NewBag[int]()
		asrt.True(bag.IsEmpty())
		asrt.Equal(0, bag.Size())
		bag.Add(134)
		bag.Add(25)
		bag.Add(67)
		bag.Add(43)
		bag.Add(29)
		bag.Add(3)
		bag.Add(23)
		asrt.False(bag.IsEmpty())
		asrt.Equal(7, bag.Size())
		asrt.Equal([]int{134, 25, 67, 43, 29, 3, 23}, bag.Iterate())
		bag.Reverse()
		asrt.Equal([]int{23, 3, 29, 43, 67, 25, 134}, bag.Iterate())
		asrt.False(bag.IsEmpty())
		asrt.Equal(7, bag.Size())
	})

	t.Run("stack", func(t *testing.T) {
		stack := array.NewStack[int]()
		asrt.True(stack.IsEmpty())
		asrt.Equal(0, stack.Size())
		stack.Push(134)
		stack.Push(25)
		stack.Push(67)
		stack.Push(43)
		stack.Push(29)
		stack.Push(3)
		stack.Push(23)
		asrt.False(stack.IsEmpty())
		asrt.Equal(7, stack.Size())
		asrt.Equal([]int{134, 25, 67, 43, 29, 3, 23}, stack.Iterate())

		stack.Reverse()
		slen := stack.Size()
		arr := make([]int, 7)
		for i := 0; i < slen; i++ {
			arr[i] = stack.Pop()
		}
		asrt.Equal([]int{134, 25, 67, 43, 29, 3, 23}, arr)
		asrt.True(stack.IsEmpty())
		asrt.Equal(0, stack.Size())

		stack.Push(134)
		stack.Push(25)
		stack.Push(67)
		stack.Push(43)
		stack.Push(29)
		stack.Push(3)
		stack.Push(23)
		asrt.False(stack.IsEmpty())
		asrt.Equal(7, stack.Size())

		stack = array.NewStack[int]()
		asrt.Panics(func() { stack.Pop() }, "algos.array.(Stack).Pop():  the code is not panic when structure is empty")
	})

	t.Run("queue", func(t *testing.T) {
		queue := array.NewQueue[int]()
		asrt.True(queue.IsEmpty())
		asrt.Equal(0, queue.Size())
		queue.Enq(134)
		queue.Enq(25)
		queue.Enq(67)
		queue.Enq(43)
		queue.Enq(29)
		queue.Enq(3)
		queue.Enq(23)
		asrt.False(queue.IsEmpty())
		asrt.Equal(7, queue.Size())
		asrt.Equal([]int{134, 25, 67, 43, 29, 3, 23}, queue.Iterate())
		queue.Reverse()
		qlen := queue.Size()
		arr := make([]int, 7)
		for i := 0; i < qlen; i++ {
			arr[i] = queue.Deq()
		}
		asrt.Equal([]int{23, 3, 29, 43, 67, 25, 134}, arr)
		asrt.True(queue.IsEmpty())
		asrt.Equal(0, queue.Size())

		queue.Enq(134)
		queue.Enq(25)
		queue.Enq(67)
		queue.Enq(43)
		queue.Enq(29)
		queue.Enq(3)
		queue.Enq(23)
		asrt.False(queue.IsEmpty())
		asrt.Equal(7, queue.Size())

		queue = array.NewQueue[int]()
		asrt.Panics(func() { queue.Deq() }, "algos.array.(Queue).Deq():  the code is not panic when structure is empty")
	})
}
