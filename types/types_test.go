package types_test

import (
	"errors"
	"testing"

	I "github.com/kselnaag/algos/types"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	asrt := assert.New(t)
	defer func() {
		err := recover()
		asrt.Nil(err)
	}()

	t.Run("Option", func(t *testing.T) {
		some := I.NewOptionSome(42)
		none := I.NewOptionNone[int]()
		asrt.True(some.IsDefined())
		asrt.False(none.IsDefined())
		asrt.Equal(42, some.Unbox(func() int {
			panic("algos.types.(Option): Option can not be Unboxed")
		}))
		asrt.Panics(func() {
			none.Unbox(func() int {
				panic("algos.types.(Option): Option can not be Unboxed")
			})
		})
		asrt.Equal("Some(42)", some.ToString())
		asrt.Equal("None", none.ToString())

	})
	t.Run("Either", func(t *testing.T) {
		left := I.NewEitherLeft[int, string](42)
		right := I.NewEitherRight[int]("string")
		wrong := I.Either[int, string]{}

		asrt.True(left.IsLeft())
		asrt.False(left.IsRight())
		asrt.False(right.IsLeft())
		asrt.True(right.IsRight())
		asrt.Panics(func() { wrong.IsLeft() })
		asrt.Panics(func() { wrong.IsRight() })

		asrt.Equal(I.NewOptionSome(42), left.ValLeft())
		asrt.Equal(I.NewOptionNone[string](), left.ValRight())
		asrt.Equal(I.NewOptionNone[int](), right.ValLeft())
		asrt.Equal(I.NewOptionSome("string"), right.ValRight())
		asrt.Panics(func() { wrong.ValLeft() })
		asrt.Panics(func() { wrong.ValRight() })

		asrt.Equal(42, left.Left(func() int {
			panic("algos.types.(Either): Either can not be Left")
		}))
		asrt.Panics(func() {
			left.Right(func() string {
				panic("algos.types.(Either): Either can not be Right")
			})
		})
		asrt.Panics(func() {
			right.Left(func() int {
				panic("algos.types.(Either): Either can not be Left")
			})
		})
		asrt.Equal("string", right.Right(func() string {
			panic("algos.types.(Either): Either can not be Right")
		}))
		asrt.Panics(func() {
			wrong.Left(func() int {
				panic("algos.types.(Either): Either can not be Left")
			})
		})
		asrt.Panics(func() {
			wrong.Right(func() string {
				panic("algos.types.(Either): Either can not be Right")
			})
		})

		asrt.Equal("Left(42)", left.ToString())
		asrt.Equal("Right(string)", right.ToString())
		asrt.Panics(func() { wrong.ToString() })
	})
	t.Run("Result", func(t *testing.T) {
		testerror := errors.New("test error")
		result := I.NewResult(42)
		err := I.NewResultError[int](testerror)
		wrong := I.Result[int]{}

		asrt.False(result.IsErr())
		asrt.True(result.IsRes())
		asrt.True(err.IsErr())
		asrt.False(err.IsRes())
		asrt.Panics(func() { wrong.IsErr() })
		asrt.Panics(func() { wrong.IsRes() })

		asrt.Equal(I.NewOptionNone[error](), result.ValErr())
		asrt.Equal(I.NewOptionSome(42), result.ValRes())
		asrt.Equal(I.NewOptionSome(testerror), err.ValErr())
		asrt.Equal(I.NewOptionNone[int](), err.ValRes())
		asrt.Panics(func() { wrong.ValErr() })
		asrt.Panics(func() { wrong.ValRes() })

		asrt.Equal(42, result.Res(func() int { panic("algos.types.(Result): Result can not be Resed") }))
		asrt.Panics(func() {
			e := result.Err(func() error { panic("algos.types.(Result): Result can not be Erred") })
			if e != nil {
				panic("algos.types.(Result): Result can not be Erred")
			}
		})
		asrt.Panics(func() {
			err.Res(func() int { panic("algos.types.(Result): Result can not be Resed") })
		})
		asrt.Equal(testerror, err.Err(func() error { panic("algos.types.(Result): Result can not be Erred") }))
		asrt.Panics(func() {
			wrong.Res(func() int { panic("algos.types.(Result): Result can not be Resed") })
		})
		asrt.Panics(func() {
			e := wrong.Err(func() error { panic("algos.types.(Result): Result can not be Erred") })
			if e != nil {
				panic("algos.types.(Result): Result can not be Erred")
			}
		})

		asrt.Equal(42, result.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") }))
		asrt.Panics(func() {
			err.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") })
		})
		asrt.Panics(func() {
			wrong.Unbox(func() int { panic("algos.types.(Result): Result can not be Unboxed") })
		})

		asrt.Equal("Result(42)", result.ToString())
		asrt.Equal("Error(test error)", err.ToString())
		asrt.Panics(func() { wrong.ToString() })
	})
}
