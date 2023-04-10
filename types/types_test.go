package types_test

import (
	"errors"
	"testing"

	I "algos/types"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		err := recover()
		assert.Nil(err)
	}()

	t.Run("Option", func(t *testing.T) {
		some := I.NewOptionSome(42)
		none := I.NewOptionNone[int]()
		assert.Equal(some.IsDefined(), true)
		assert.Equal(none.IsDefined(), false)
		assert.Equal(some.Unbox(func() int {
			panic("algos.types.(Option): Option can not be Unboxed")
		}), 42)
		assert.Panics(func() {
			none.Unbox(func() int {
				panic("algos.types.(Option): Option can not be Unboxed")
			})
		})
		assert.Equal(some.ToString(), "Some(42)")
		assert.Equal(none.ToString(), "None")

	})
	t.Run("Either", func(t *testing.T) {
		left := I.NewEitherLeft[int, string](42)
		right := I.NewEitherRight[int]("string")
		wrong := I.Either[int, string]{}

		assert.Equal(left.IsLeft(), true)
		assert.Equal(left.IsRight(), false)
		assert.Equal(right.IsLeft(), false)
		assert.Equal(right.IsRight(), true)
		assert.Panics(func() { wrong.IsLeft() })
		assert.Panics(func() { wrong.IsRight() })

		assert.Equal(left.ValLeft(), I.NewOptionSome(42))
		assert.Equal(left.ValRight(), I.NewOptionNone[string]())
		assert.Equal(right.ValLeft(), I.NewOptionNone[int]())
		assert.Equal(right.ValRight(), I.NewOptionSome("string"))
		assert.Panics(func() { wrong.ValLeft() })
		assert.Panics(func() { wrong.ValRight() })

		assert.Equal(left.Left(func() int {
			panic("algos.types.(Either): Either can not be Lefted")
		}), 42)
		assert.Panics(func() {
			left.Right(func() string {
				panic("algos.types.(Either): Either can not be Righted")
			})
		})
		assert.Panics(func() {
			right.Left(func() int {
				panic("algos.types.(Either): Either can not be Lefted")
			})
		})
		assert.Equal(right.Right(func() string {
			panic("algos.types.(Either): Either can not be Righted")
		}), "string")
		assert.Panics(func() {
			wrong.Left(func() int {
				panic("algos.types.(Either): Either can not be Lefted")
			})
		})
		assert.Panics(func() {
			wrong.Right(func() string {
				panic("algos.types.(Either): Either can not be Righted")
			})
		})

		assert.Equal(left.ToString(), "Left(42)")
		assert.Equal(right.ToString(), "Right(string)")
		assert.Panics(func() { wrong.ToString() })
	})
	t.Run("Result", func(t *testing.T) {
		testerror := errors.New("test error")
		result := I.NewResult(42)
		err := I.NewResultError[int](testerror)
		wrong := I.Result[int]{}

		assert.Equal(result.IsErr(), false)
		assert.Equal(result.IsRes(), true)
		assert.Equal(err.IsErr(), true)
		assert.Equal(err.IsRes(), false)
		assert.Panics(func() { wrong.IsErr() })
		assert.Panics(func() { wrong.IsRes() })

		assert.Equal(result.ValErr(), I.NewOptionNone[error]())
		assert.Equal(result.ValRes(), I.NewOptionSome(42))
		assert.Equal(err.ValErr(), I.NewOptionSome(testerror))
		assert.Equal(err.ValRes(), I.NewOptionNone[int]())
		assert.Panics(func() { wrong.ValErr() })
		assert.Panics(func() { wrong.ValRes() })

		assert.Equal(result.Res(func() int { panic("algos.types.(Result): Result can not be Resed") }), 42)
		assert.Panics(func() {
			e := result.Err(func() error { panic("algos.types.(Result): Result can not be Erred") })
			if e != nil {
				panic("algos.types.(Result): Result can not be Erred")
			}
		})
		assert.Panics(func() {
			err.Res(func() int { panic("algos.types.(Result): Result can not be Resed") })
		})
		assert.Equal(err.Err(func() error { panic("algos.types.(Result): Result can not be Erred") }), testerror)
		assert.Panics(func() {
			wrong.Res(func() int { panic("algos.types.(Result): Result can not be Resed") })
		})
		assert.Panics(func() {
			e := wrong.Err(func() error { panic("algos.types.(Result): Result can not be Erred") })
			if e != nil {
				panic("algos.types.(Result): Result can not be Erred")
			}
		})

		assert.Equal(result.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") }), 42)
		assert.Panics(func() {
			err.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") })
		})
		assert.Panics(func() {
			wrong.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") })
		})

		assert.Equal(result.ToString(), "Result(42)")
		assert.Equal(err.ToString(), "Error(test error)")
		assert.Panics(func() { wrong.ToString() })
	})
}
