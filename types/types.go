package types

import (
	"fmt"
)

type Ord interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Ptr interface {
	~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// ====================
type Option[O any] struct {
	val O
	def bool
}

func NewOptionSome[O any](val O) Option[O] {
	return Option[O]{
		val: val,
		def: true,
	}
}

func NewOptionNone[O any]() Option[O] {
	return Option[O]{}
}

func (opt Option[O]) IsDefined() bool {
	return opt.def
}

func (opt Option[O]) Unbox(catch func() O) O {
	if opt.IsDefined() {
		return opt.val
	}
	return catch()
}

func (opt Option[O]) ToString() string {
	if opt.IsDefined() {
		return fmt.Sprintf("Some(%v)", opt.val)
	}
	return "None"
}

// ====================
type Either[L, R any] struct { // Either[L,R]{} - WRONG CONSTRUCTOR !  APPROPRIATE INIT REQUIRED
	l Option[L]
	r Option[R]
}

func NewEitherLeft[L, R any](left L) Either[L, R] {
	return Either[L, R]{
		l: NewOptionSome(left),
		r: NewOptionNone[R](),
	}
}

func NewEitherRight[L, R any](right R) Either[L, R] {
	return Either[L, R]{
		l: NewOptionNone[L](),
		r: NewOptionSome(right),
	}
}

func (eit Either[L, R]) isvalid() {
	bothNone := !eit.l.IsDefined() && !eit.r.IsDefined()
	bothSome := eit.l.IsDefined() && eit.r.IsDefined()
	if bothNone || bothSome {
		panic("algos.types.(Either[L,R]).isvalid() fails: There is a bad structure inside Either type")
	}
}

func (eit Either[L, R]) IsLeft() bool {
	eit.isvalid()
	return eit.l.IsDefined()
}

func (eit Either[L, R]) IsRight() bool {
	eit.isvalid()
	return eit.r.IsDefined()
}

func (eit Either[L, R]) ValLeft() Option[L] {
	eit.isvalid()
	return eit.l
}

func (eit Either[L, R]) ValRight() Option[R] {
	eit.isvalid()
	return eit.r
}

func (eit Either[L, R]) Left(catch func() L) L {
	eit.isvalid()
	if eit.IsLeft() {
		return eit.l.Unbox(catch)
	}
	return catch()
}

func (eit Either[L, R]) Right(catch func() R) R {
	eit.isvalid()
	if eit.IsRight() {
		return eit.r.Unbox(catch)
	}
	return catch()
}

func (eit Either[L, R]) ToString() string {
	eit.isvalid()
	if eit.IsLeft() {
		return fmt.Sprintf("Left(%v)", eit.Left(func() L {
			panic("algos.types.(Either[L,R]): cannot get L value")
		}))
	}
	return fmt.Sprintf("Right(%v)", eit.Right(func() R {
		panic("algos.types.(Either[L,R]): cannot get R value")
	}))
}

// ====================
type Result[T any] struct { // Result[T]]{} - WRONG CONSTRUCTOR ! APPROPRIATE INIT REQUIRED
	Either[error, T]
}

func NewResultError[T any](err error) Result[T] {
	return Result[T]{
		Either: NewEitherLeft[error, T](err),
	}
}

func NewResult[T any](res T) Result[T] {
	return Result[T]{
		Either: NewEitherRight[error](res),
	}
}

func (res Result[T]) IsErr() bool {
	return res.IsLeft()
}

func (res Result[T]) IsRes() bool {
	return res.IsRight()
}

func (res Result[T]) ValErr() Option[error] {
	return res.ValLeft()
}

func (res Result[T]) ValRes() Option[T] {
	return res.ValRight()
}

func (res Result[T]) Err(catch func() error) error {
	if res.IsErr() {
		return res.ValErr().Unbox(catch)
	}
	return catch()
}

func (res Result[T]) Res(catch func() T) T {
	if res.IsRes() {
		return res.ValRes().Unbox(catch)
	}
	return catch()
}

func (res Result[T]) Unbox(catch func() T) T {
	if res.IsRes() {
		return res.Res(catch)
	}
	return catch()
}

func (res Result[T]) ToString() string {
	if res.IsErr() {
		return fmt.Sprintf("Error(%v)", res.Err(func() error {
			panic("algos.types.(Result): cannot get Error value")
		}).Error())
	}
	return fmt.Sprintf("Result(%v)", res.Res(func() T {
		panic("algos.types.(Result): cannot get Result value")
	}))
}
