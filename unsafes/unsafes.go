package unsafes

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintAsBinary(a any) string {
	type iface struct {
		t, v unsafe.Pointer
	}
	p := uintptr((*iface)(unsafe.Pointer(&a)).v)
	t := reflect.TypeOf(a)
	resStr := ""
	for i := 0; i < int(t.Size()); i++ {
		n := *((*byte)(unsafe.Pointer(p)))
		resStr += fmt.Sprintf("%08b ", n)
		p += unsafe.Sizeof(n)
	}
	return resStr
}
