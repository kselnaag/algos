package types

import "fmt"

type TestStruct struct {
	A int
	B int
}

func (s TestStruct) CompareTo(obj STOrd) int {
	objComp, ok := obj.(*TestStruct)
	if !ok {
		panic(fmt.Sprintf("algos.types.TestStruct.CompareTo(obj STOrd): "+
			"Type of arg is unknown, expected *types.TestStruct, actual %T", obj))
	}
	compFactor := func(st TestStruct) int {
		return st.A + st.B
	}
	this := compFactor(s)
	that := compFactor(*objComp)
	switch {
	case this < that:
		return -1
	case this > that:
		return +1
	default:
		return 0
	}
}

func ConvToByteArr(key any) []byte {
	return []byte(fmt.Sprintf("%v", key))
}

func LT[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case STOrd:
		jj, ok := any(j).(STOrd)
		if ok {
			return ii.CompareTo(jj) < 0
		}
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii < jj
		}
	case int8:
		jj, ok := any(j).(int8)
		if ok {
			return ii < jj
		}
	case int16:
		jj, ok := any(j).(int16)
		if ok {
			return ii < jj
		}
	case int32:
		jj, ok := any(j).(int32)
		if ok {
			return ii < jj
		}
	case int64:
		jj, ok := any(j).(int64)
		if ok {
			return ii < jj
		}
	case float32:
		jj, ok := any(j).(float32)
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
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii < jj
		}
	case uint8:
		jj, ok := any(j).(uint8)
		if ok {
			return ii < jj
		}
	case uint16:
		jj, ok := any(j).(uint16)
		if ok {
			return ii < jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii < jj
		}
	case uint64:
		jj, ok := any(j).(uint64)
		if ok {
			return ii < jj
		}
	}
	panic(fmt.Sprintf("algos.types.LT[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}

func GT[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case STOrd:
		jj, ok := any(j).(STOrd)
		if ok {
			return ii.CompareTo(jj) > 0
		}
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii > jj
		}
	case int8:
		jj, ok := any(j).(int8)
		if ok {
			return ii > jj
		}
	case int16:
		jj, ok := any(j).(int16)
		if ok {
			return ii > jj
		}
	case int32:
		jj, ok := any(j).(int32)
		if ok {
			return ii > jj
		}
	case int64:
		jj, ok := any(j).(int64)
		if ok {
			return ii > jj
		}
	case float32:
		jj, ok := any(j).(float32)
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
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii > jj
		}
	case uint8:
		jj, ok := any(j).(uint8)
		if ok {
			return ii > jj
		}
	case uint16:
		jj, ok := any(j).(uint16)
		if ok {
			return ii > jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii > jj
		}
	case uint64:
		jj, ok := any(j).(uint64)
		if ok {
			return ii > jj
		}
	}
	panic(fmt.Sprintf("algos.types.GT[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}

func EQ[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case STOrd:
		jj, ok := any(j).(STOrd)
		if ok {
			return ii.CompareTo(jj) == 0
		}
	case int:
		jj, ok := any(j).(int)
		if ok {
			return ii == jj
		}
	case int8:
		jj, ok := any(j).(int8)
		if ok {
			return ii == jj
		}
	case int16:
		jj, ok := any(j).(int16)
		if ok {
			return ii == jj
		}
	case int32:
		jj, ok := any(j).(int32)
		if ok {
			return ii == jj
		}
	case int64:
		jj, ok := any(j).(int64)
		if ok {
			return ii == jj
		}
	case float32:
		jj, ok := any(j).(float32)
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
	case uint:
		jj, ok := any(j).(uint)
		if ok {
			return ii == jj
		}
	case uint8:
		jj, ok := any(j).(uint8)
		if ok {
			return ii == jj
		}
	case uint16:
		jj, ok := any(j).(uint16)
		if ok {
			return ii == jj
		}
	case uint32:
		jj, ok := any(j).(uint32)
		if ok {
			return ii == jj
		}
	case uint64:
		jj, ok := any(j).(uint64)
		if ok {
			return ii == jj
		}
	}
	panic(fmt.Sprintf("algos.types.EQ[T any](i, j T): "+
		"Type of args is not processed: arg Type is %T", i))
}
