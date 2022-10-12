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
	case uint:
		jj := any(j).(uint)
		return ii < jj
	case uint32:
		jj := any(j).(uint32)
		return ii < jj
	default:
		s := "algos.(array).equals.lt[T any](i, j T): Type of args is not Ord or Comp interface: "
		s += fmt.Sprintf("arg Type is: %T", i)
		panic(s)
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
	case uint:
		jj := any(j).(uint)
		return ii > jj
	case uint32:
		jj := any(j).(uint32)
		return ii > jj
	default:
		s := "algos.(array).equals.gt[T any](i, j T): Type of args is not Ord or Comp interface: "
		s += fmt.Sprintf("arg Type is: %T", i)
		panic(s)
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
	case uint:
		jj := any(j).(uint)
		return ii == jj
	case uint32:
		jj := any(j).(uint32)
		return ii == jj
	default:
		s := "algos.(array).equals.eq[T any](i, j T): Type of args is not Ord or Comp interface: "
		s += fmt.Sprintf("arg Type is: %T", i)
		panic(s)
	}
}
