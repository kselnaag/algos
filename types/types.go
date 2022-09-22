package types

type Iter[T any] interface {
	Next() T
}

type Comp interface {
	CompareTo(Comp) int
}

type Ord interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
