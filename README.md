<p align="left">
	<img src="https://img.shields.io/github/languages/code-size/kselnaag/algos?style=plastic" title="src files size" alt="src files size">
	<img src="https://img.shields.io/github/repo-size/kselnaag/algos?style=plastic" title="repo size" alt="repo size">
	<a href="https://github.com/kselnaag/algos/blob/master/LICENSE" title="LICENSE"><img src="https://img.shields.io/github/license/kselnaag/algos?style=plastic" alt="license"></a>
	<a href="https://github.com/kselnaag/algos/actions" title="Workflows"><img src="https://img.shields.io/github/actions/workflow/status/kselnaag/algos/go.yml?branch=master&style=plastic" alt="tests checks"></a>
	<a href="https://kselnaag.github.io/algos" title="coverage"><img src="https://img.shields.io/badge/GHpages-coverage-blueviolet?style=plastic" alt="coverage"></a>
</p>
<p align="left">
	<img src="https://img.shields.io/static/v1?label=build%20by&message=Go&color=ffa757&style=plastic" alt="build by Go">
	<img src="https://img.shields.io/static/v1?label=make%20with&message=%F0%9F%8E%89%E2%9C%A8&color=ffa757&style=plastic" alt="make with fun">
</p>

### **ALGOS**
----

<p align="center">
  <img src="https://raw.githubusercontent.com/kselnaag/algos/master/pics/myGophers.jpg" title="#DIY #GENERICS #ERRORS" alt="#DIY #GENERICS #ERRORS"/>
</p>

This is the hand-made algorithms and data-structures module with go generics. It has main abstractions based on Robert Sedgewick "Algorithms" book and his Prinestone video course. There is implementation of basic versions of arrays, lists, trees, tries, graphs and some algos to work with them.

### **Types and func scheme:**
<p align="center">
  <img src="https://raw.githubusercontent.com/kselnaag/algos/master/pics/algos.png" title="Types and func scheme" alt="Types and func scheme"/>
</p>

**Motivation:**
The main idea of this module is to suggest more convenient way for sorting slices of different types. We have Pre-defined Data Types (PDTs: ints, floats, strings) and Abstract Data Types (ADTs: self-made structures). Now in GOLANG stdlib we have to wrap PDTs in structs and bind 3 methods (`len`, `less`, `swap`) to call a sort function, same with ADTs. This module dedicates `Ord` interface for PDTs with `>`, `==`, `<` operators and `Comp` interface for ADTs with `CompareTo` method and `(+1, 0, -1)` values as the result for sorting internal and abstract types more easily.

**The Main Idea:**
We can build ADTs:
```
type Comp interface {
	CompareTo(Comp) int
}

type TestStruct struct {
	A int
	B int
}

func (s TestStruct) CompareTo(obj Comp) int {
	switch obj.(type) {
	case *TestStruct:
		break
	default:
		panic(fmt.Sprintf("algos.types.TestStruct.CompareTo(obj Comp): Type of arg is unknown, expected *types.TestStruct, actual %T", obj))
	}
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
```

We can compare data types:
```
func LT[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case Comp:
		jj := any(j).(I.Comp)
		return ii.CompareTo(jj) < 0
	case int:
		jj := any(j).(int)
		return ii < jj
	case uint:
		jj := any(j).(uint)
		return ii < jj
	case float64:
		jj := any(j).(float64)
		return ii < jj
	case string:
		jj := any(j).(string)
		return ii < jj
	default:
		panic(fmt.Sprintf("algos.types.LT[T any](i, j T): Type of args is not processed: arg Type is: %T", i))
	}
}
```

We can sort everything:
```
func InsertSort[T any](arr []T) {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		for j := i; j > 0; j-- {
			if LT(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}
```

**Result:**
For PDTs, value slice:
```
arr := []int{3, 2, 1}
array.InsertSort(arr)
// {1, 2, 3}
```

For ADTs, pointer slice:
```
s1 := &TestStruct{1, 2}
s2 := &TestStruct{3, 1}
s3 := &TestStruct{2, 3}
arr := []*TestStruct{s3, s2, s1}
array.InsertSort(arr)
// {s1, s2, s3}
```

This approach used in arrays. Other data types are under construction and use `Ord` and `any` interfaces for Keys and Values.

----

### **Links**: 
| [samber/lo](https://github.com/samber/lo "Lodash-style Go library") | [samber/do](https://github.com/samber/do "Dependency injection toolkit based on Go 1.18+ generics") | [samber/mo](https://github.com/samber/mo "Monads based on Go 1.18+ generics") | [ialekseev/go4fun](https://github.com/ialekseev/go4fun "Functional primitives and patterns in go") |
