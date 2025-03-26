package sync_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	S "algos/sync"
	I "algos/types"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test error")

func testError(msg string) error {
	return fmt.Errorf("%w: %s", errTest, msg)
}

func TestSync(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("Future", func(t *testing.T) {
		fut := S.NewFuture(func() I.Result[int] {
			time.Sleep(10 * time.Millisecond)
			return I.NewResult(42)
		})
		asrt.False(fut.IsCompleted())
		r := make(chan int)
		fut.OnComplete(func(i I.Result[int]) {
			r <- i.Unbox(func() int { panic("algos.sync.(Future): Result can not be Unboxed") })
		})
		asrt.Equal(42, <-r)
		asrt.Equal(42, fut.Value().Res(func() int { panic("algos.sync.(Future): Result can not be Resed") }))
		asrt.True(fut.IsCompleted())
		asrt.Equal("Result(42)", fut.ToString())

		err := S.NewFuture(func() I.Result[int] {
			time.Sleep(10 * time.Millisecond)
			return I.NewResultError[int](testError(""))
		})
		asrt.False(err.IsCompleted())
		err.OnComplete(func(i I.Result[int]) {
			asrt.Panics(func() {
				i.Unbox(func() int { panic("algos.sync.(Future): Result can not be Unboxed") })
			})
		})
		asrt.Equal(testError(""), err.Value().Err(func() error { panic("algos.sync.(Future): Result can not be Resed") }))
		asrt.True(err.IsCompleted())
		asrt.Equal("Error(test error: )", err.ToString())

		wrong := S.NewFuture(func() I.Result[int] { return I.Result[int]{} })
		asrt.False(wrong.IsCompleted())
		wrong.OnComplete(func(i I.Result[int]) {
			asrt.Panics(func() {
				i.Unbox(func() int { panic("algos.sync.(Future): Result can not be Unboxed") })
			})
		})
		asrt.Panics(func() {
			e := wrong.Value().Err(func() error { panic("algos.sync.(Future): Result can not be Resed") })
			if e != nil {
				panic("algos.sync.(Future): Result can not be Resed")
			}
		})
		asrt.True(wrong.IsCompleted())
		asrt.Panics(func() { wrong.ToString() })
	})
	t.Run("Promise", func(t *testing.T) {
		d := make(chan int)
		prom := S.NewPromise(200, func() I.Result[int] {
			time.Sleep(10 * time.Millisecond)
			return I.NewResult(42)
		})
		asrt.False(prom.IsCompleted())
		prom.OnComplete(func(i I.Result[int]) {
			d <- i.Unbox(func() int { panic("algos.sync.(Promise): Result can not be Unboxed") })
		})
		asrt.Equal(42, <-d)
		asrt.Equal(42, prom.Value().Res(func() int { panic("algos.sync.(Promise): Result can not be Resed") }))
		asrt.True(prom.IsCompleted())
		asrt.Equal("Result(42)", prom.ToString())

		morp := S.NewPromise(50, func() I.Result[int] {
			time.Sleep(100 * time.Millisecond)
			return I.NewResult(42)
		})
		asrt.False(morp.IsCompleted())
		morp.OnComplete(func(i I.Result[int]) {
			asrt.Panics(func() {
				i.Unbox(func() int { panic("algos.sync.(Promise): Result can not be Unboxed") })
			})
		})
		asrt.Equal(S.FutureError("callback is timed out"), morp.Value().Err(func() error { panic("algos.sync.(Promise): Result can not be Resed") }))
		asrt.True(morp.IsCompleted())
		asrt.Equal("Error("+S.FutureError("callback is timed out").Error()+")", morp.ToString())
	})
}
