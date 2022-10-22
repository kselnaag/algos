<p align="left">
	<img src="https://img.shields.io/github/languages/top/kselnaag/algos?style=plastic" title="main language" alt="main language">
	<img src="https://img.shields.io/github/go-mod/go-version/kselnaag/algos?style=plastic" title="language version" alt="language version">
	<img src="https://img.shields.io/github/languages/code-size/kselnaag/algos?style=plastic" title="src files size" alt="src files size">
	<img src="https://img.shields.io/github/repo-size/kselnaag/algos?style=plastic" title="repo size" alt="repo size">
	<a href="https://github.com/kselnaag/algos/blob/master/LICENSE" title="LICENSE"><img src="https://img.shields.io/github/license/kselnaag/algos?style=plastic" alt="license"></a>
	<a href="https://github.com/kselnaag/algos/actions" title="Workflows"><img src="https://img.shields.io/github/workflow/status/kselnaag/algos/algoscommit/master?style=plastic" alt="tests checks"></a>
	<a href="https://kselnaag.github.io/algos" title="coverage"><img src="https://img.shields.io/badge/GHpages-coverage-blueviolet?style=plastic" alt="coverage"></a>
</p>

### **ALGOS**
----

This is the hand-made algorithms and data-structures module with go generics. It has main abstractions based on Robert Sedgewick "Algorithms" book and his Prinestone video course. There is implementation of basic versions of arrays, lists, trees, tries, graphs and some algos to work with them.

### **Types and func approximate scheme**
<p align="center">
  <img src="https://raw.githubusercontent.com/kselnaag/algos/master/pics/algos.png" alt="Types and func scheme"/>
</p>

The main idea of this module is to suggest more convenient way for sorting slices of different types. We have Pre-defined Data Types (PDTs: ints, floats, strings) and Abstract Data Types (ADTs: self-made structures) in our code. Now in std go lib we have to wrap PDTs in structs and bind 3 methods (LEN, LESS, SWAP) to call a sort function, same with ADTs. This module dedicates `Ord` interface for PDTs with `>`, `==`, `<` operators and `Comp` interface for ADTs with `CompareTo` method for sorting internal and abstract types more easy way.

We can build ADTs:
```
type myType struct {
	a int
	b int
}

func (s myType) CompareTo(st Comp) int {
	this := s.a + s.b
	that := (st.(*myType)).a + (st.(*myType)).b
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
func lt[T any](i, j T) bool {
	switch ii := any(i).(type) {
	case I.Comp:
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
		s := "algos.(array).equals.lt[T any](i, j T): Type of args is not Ord or Comp interface: "
		s += fmt.Sprintf("arg Type is: %T", i)
		panic(s)
	}
}
```

We can sort everything:
```
func InsertSort[T any](arr []T) {
	alen := len(arr)
	for i := 1; i < alen; i++ {
		for j := i; j > 0; j-- {
			if lt(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}
```

For PDTs, value slice:
```
arr := []int{3, 2, 1}
array.InsertSort(arr)
// {1, 2, 3}
```

For ADTs, pointer slice:
```
s1 := &myType{1, 2}
s2 := &myType{3, 1}
s3 := &myType{2, 3}
arr := []*myType{s3, s2, s1}
array.InsertSort(arr)
// {s1, s2, s3}
```

This approach used in arrays. Other data types are under construction and use `Ord` or `any` interfaces.

----

### **Links**: 
| [samber/lo](https://github.com/samber/lo "Lodash-style Go library") | [samber/do](https://github.com/samber/do "Dependency injection toolkit based on Go 1.18+ generics") | [samber/mo](https://github.com/samber/mo "Monads based on Go 1.18+ generics") | [ialekseev/go4fun](https://github.com/ialekseev/go4fun "Functional primitives and patterns in go") |




