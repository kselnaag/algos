package types

import "fmt"

type Comp interface {
	CompareTo(Comp) int
}

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

type TestStruct struct {
	A int
	B int
}

func (s TestStruct) CompareTo(obj Comp) int {
	this := s.A + s.B
	that := (obj.(*TestStruct)).A + (obj.(*TestStruct)).B
	if this < that {
		return -1
	}
	if this > that {
		return +1
	}
	return 0
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
		s := "algos.(math).ConvToByteArr(mess any): Type of arg is Ord interface, but not processed: "
		s += fmt.Sprintf("arg Type is: %T", mess)
		panic(s)
	}
	return res
}

func LT[T any](i, j T) bool {
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
	case Comp:
		jj := any(j).(Comp)
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

func GT[T any](i, j T) bool {
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
	case Comp:
		jj := any(j).(Comp)
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

func EQ[T any](i, j T) bool {
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
	case Comp:
		jj := any(j).(Comp)
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
