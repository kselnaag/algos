package array

func Map[T1, T2 any](arr []T1, fnc func(T1) T2) []T2 {
	res := make([]T2, 0, len(arr))
	for _, el := range arr {
		res = append(res, fnc(el))
	}
	return res
}

func MapA[T1, T2 any](arr []T1, fnc func(T1) T2) []T2 {
	res := make([]T2, 0, len(arr))
	chans := make([]chan T2, len(arr))
	for i, el := range arr {
		elemChan := make(chan T2)
		chans[i] = elemChan
		go func(elemChan chan<- T2, el T1) {
			elemChan <- fnc(el)
		}(elemChan, el)
	}
	for _, elemChan := range chans {
		res = append(res, <-elemChan)
	}
	return res
}

func Reduce[T1, T2 any](arr []T1, fnc func(T2, T1) T2, acc T2) T2 {
	for _, el := range arr {
		acc = fnc(acc, el)
	}
	return acc
}

func ReduceR[T1, T2 any](arr []T1, fnc func(T2, T1) T2, acc T2) T2 {
	alen := len(arr)
	for i := alen - 1; i >= 0; i-- {
		acc = fnc(acc, arr[i])
	}
	return acc
}

func Filter[T any](arr []T, fnc func(T) bool) []T {
	res := make([]T, 0)
	for _, el := range arr {
		if fnc(el) {
			res = append(res, el)
		}
	}
	return res
}

func Perm[T any](a []T, f func([]T)) {
	perm(a, f, 0)
}

func perm[T any](a []T, f func([]T), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
