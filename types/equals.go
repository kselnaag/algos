package types

import "fmt"

type Comp interface {
	CompareTo(Comp) int
}

type TestStruct struct {
	A int
	B int
}

func (s TestStruct) CompareTo(obj Comp) int {
	var this, that int
	compFactor := func(st TestStruct) int {
		return st.A + st.B
	}
	switch objComp := obj.(type) {
	case *TestStruct:
		this = compFactor(s)
		that = compFactor(*objComp)
	default:
		panic(fmt.Sprintf("algos.types.TestStruct.CompareTo(obj Comp): "+
			"Type of arg is unknown, expected *types.TestStruct, actual %T", obj))
	}
	switch {
	case this < that:
		return -1
	case this > that:
		return +1
	default:
		return 0
	}
}

func ConvToByteArr[T Ord](mess T) []byte {
	res := make([]byte, 8)
	switch m := any(mess).(type) {
	case int:
		mi := uint(m)
		for i := 0; i < 8; i++ {
			res[i] = byte(mi >> (8 * (7 - i)))
		}
	case uint:
		for i := 0; i < 8; i++ {
			res[i] = byte(m >> (8 * (7 - i)))
		}
	case float64:
		mi := uint(m)
		for i := 0; i < 8; i++ {
			res[i] = byte(mi >> (8 * (7 - i)))
		}
	case string:
		return []byte(m)
	default:
		panic(fmt.Sprintf("algos.types.ConvToByteArr(mess any): "+
			"Type of arg is Ord interface, but not processed: arg Type is %T", mess))
	}
	return res
}

func LT[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii < jj
		}
	case float64:
		jj, ok := any(j).(float64)
		if ok {
			return ii < jj
		}
	case string:
		jj, ok := any(j).(string)
		if ok {
			return ii < jj
		}
	case Comp:
		jj, ok := any(j).(Comp)
		if ok {
			return ii.CompareTo(jj) < 0
		}
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii < jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii < jj
		}
	}
	panic(fmt.Sprintf("algos.types.LT[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}

func GT[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii > jj
		}
	case float64:
		jj, ok := any(j).(float64)
		if ok {
			return ii > jj
		}
	case string:
		jj, ok := any(j).(string)
		if ok {
			return ii > jj
		}
	case Comp:
		jj, ok := any(j).(Comp)
		if ok {
			return ii.CompareTo(jj) > 0
		}
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii > jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii > jj
		}
	}
	panic(fmt.Sprintf("algos.types.GT[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}

func EQ[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii == jj
		}
	case float64:
		jj, ok := any(j).(float64)
		if ok {
			return ii == jj
		}
	case string:
		jj, ok := any(j).(string)
		if ok {
			return ii == jj
		}
	case Comp:
		jj, ok := any(j).(Comp)
		if ok {
			return ii.CompareTo(jj) == 0
		}
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii == jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii == jj
		}
	}
	panic(fmt.Sprintf("algos.types.EQ[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}
