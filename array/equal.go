package array

import (
	"fmt"

	I "github.com/kselnaag/algos/types"
)

func lt[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj := any(j).(int)
		return ii < jj
	case float64:
		jj := any(j).(float64)
		return ii < jj
	case string:
		jj := any(j).(string)
		return ii < jj
	case I.Comp:
		jj := any(j).(I.Comp)
		return ii.CompareTo(jj) < 0
	default:
		fmt.Printf("i Type is: %T.  ", i)
		panic("algos.(array).equals.lt[T any](i,j T): Type of args is not Ord or Comp")
	}
}

func gt[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj := any(j).(int)
		return ii > jj
	case float64:
		jj := any(j).(float64)
		return ii > jj
	case string:
		jj := any(j).(string)
		return ii > jj
	case I.Comp:
		jj := any(j).(I.Comp)
		return ii.CompareTo(jj) > 0
	default:
		fmt.Printf("i Type is: %T.  ", i)
		panic("algos.(array).equals.gt[T any](i,j T): Type of args is not Ord or Comp")
	}
}

func eq[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj := any(j).(int)
		return ii == jj
	case float64:
		jj := any(j).(float64)
		return ii == jj
	case string:
		jj := any(j).(string)
		return ii == jj
	case I.Comp:
		jj := any(j).(I.Comp)
		return ii.CompareTo(jj) == 0
	default:
		fmt.Printf("i Type is: %T.  ", i)
		panic("algos.(array).equals.eq[T any](i,j T): Type of args is not Ord or Comp")
	}
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
