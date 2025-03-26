package sync

import (
	"errors"
	"fmt"
	"sync"
	"time"

	I "algos/types"
)

// ====================
type Future[F any] struct {
	val  I.Result[F]
	ch   chan I.Result[F]
	once *sync.Once
	cpl  bool
}

var errFuture error = errors.New("Future[F any] error")

func FutureError(msg string) error {
	return fmt.Errorf("%w: %s", errFuture, msg)
}

func NewFuture[F any](fn func() I.Result[F]) Future[F] {
	future := Future[F]{
		val:  I.NewResultError[F](FutureError("callback is timed out")),
		ch:   make(chan I.Result[F]),
		once: new(sync.Once),
		cpl:  false,
	}
	go func() {
		future.ch <- fn()
	}()
	return future
}

func (fut *Future[F]) IsCompleted() bool {
	if len(fut.ch) > 0 {
		fut.cpl = true
	}
	return fut.cpl
}

func (fut *Future[F]) Value() I.Result[F] { // Lock until result
	fut.once.Do(func() {
		fut.val = <-fut.ch
		fut.cpl = true
	})
	return fut.val
}

func (fut *Future[F]) OnComplete(fn func(I.Result[F])) {
	go func() {
		fn(fut.Value())
	}()
}

func (fut *Future[F]) ToString() string {
	return fut.val.ToString()
}

// ====================
type Promise[P any] struct {
	Future[P]
	timeout time.Duration
}

func NewPromise[P any](ms int, fn func() I.Result[P]) Promise[P] {
	return Promise[P]{
		Future:  NewFuture(fn),
		timeout: time.Duration(ms) * time.Millisecond,
	}
}

func (pro *Promise[P]) Value() I.Result[P] {
	pro.once.Do(func() {
	label:
		for {
			select {
			case pro.val = <-pro.ch:
				break label
			case <-time.After(pro.timeout):
				break label
			}
		}
		pro.cpl = true
	})
	return pro.val
}

func (pro *Promise[P]) OnComplete(fn func(I.Result[P])) {
	go func() {
		fn(pro.Value())
	}()
}
