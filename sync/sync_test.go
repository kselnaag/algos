package sync_test

import (
	"errors"
	"testing"
	"time"

	S "github.com/kselnaag/algos/sync"
	I "github.com/kselnaag/algos/types"

	"github.com/stretchr/testify/assert"
)

func TestSync(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("Future", func(t *testing.T) {
		testerror := errors.New("future error")
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
			return I.NewResultError[int](testerror)
		})
		asrt.False(err.IsCompleted())
		err.OnComplete(func(i I.Result[int]) {
			asrt.Panics(func() {
				i.Unbox(func() int { panic("algos.sync.(Future): Result can not be Unboxed") })
			})
		})
		asrt.Equal(testerror, err.Value().Err(func() error { panic("algos.sync.(Future): Result can not be Resed") }))
		asrt.True(err.IsCompleted())
		asrt.Equal("Error(future error)", err.ToString())

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

		testerror := errors.New("algos.types.(Future): callback is timed out")
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
		asrt.Equal(testerror, morp.Value().Err(func() error { panic("algos.sync.(Promise): Result can not be Resed") }))
		asrt.True(morp.IsCompleted())
		asrt.Equal("Error(algos.types.(Future): callback is timed out)", morp.ToString())
	})
}
