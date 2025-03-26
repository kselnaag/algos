package list_test

import (
	"testing"

	"algos/list"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("bag", func(t *testing.T) {
		bag := list.NewBag[int]()
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
		asrt.Equal([]int{23, 3, 29, 43, 67, 25, 134}, bag.Iterate())
		bag.Reverse()
		asrt.Equal([]int{134, 25, 67, 43, 29, 3, 23}, bag.Iterate())
		asrt.False(bag.IsEmpty())
		asrt.Equal(7, bag.Size())
	})

	t.Run("stack", func(t *testing.T) {
		stack := list.NewStack[int]()
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
		asrt.Equal([]int{23, 3, 29, 43, 67, 25, 134}, stack.Iterate())
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
		stack = list.NewStack[int]()
		asrt.Panics(func() { stack.Pop() }, "algos.list.(Queue).Deq():  the code is not panic when structure is empty")
	})

	t.Run("queue", func(t *testing.T) {
		queue := list.NewQueue[int]()
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
		queue = list.NewQueue[int]()
		asrt.Panics(func() { queue.Deq() }, "algos.list.(Queue).Deq():  the code is not panic when structure is empty")
	})

	t.Run("deque", func(t *testing.T) {
		deq := list.NewDeque[int]()
		asrt.True(deq.IsEmpty())
		asrt.Equal(0, deq.Size())
		deq.PushFront(134)
		deq.PushFront(25)
		deq.PushFront(67)
		deq.PushBack(43)
		deq.PushBack(29)
		deq.PushBack(3)
		asrt.False(deq.IsEmpty())
		asrt.Equal(6, deq.Size())
		asrt.Equal([]int{67, 25, 134, 43, 29, 3}, deq.Iterate())

		deq.Reverse()
		asrt.Equal([]int{3, 29, 43, 134, 25, 67}, deq.Iterate())
		asrt.Equal(3, deq.Front())
		asrt.Equal(67, deq.Back())

		asrt.Equal(3, deq.PopFront())
		asrt.Equal(29, deq.PopFront())
		asrt.Equal(43, deq.PopFront())
		asrt.Equal(67, deq.PopBack())
		asrt.Equal(25, deq.PopBack())
		asrt.Equal(134, deq.PopBack())
		asrt.True(deq.IsEmpty())
		asrt.Equal(0, deq.Size())

		asrt.Panics(func() { deq.PopFront() }, "algos.list.(Deque).PopFront():  the code is not panic when structure is empty")
		asrt.Panics(func() { deq.PopBack() }, "algos.list.(Deque).PopBack():  the code is not panic when structure is empty")
	})
}
