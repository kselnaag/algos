package list

func Map[T1, T2 any](root *Snode[T1], fnc func(T1) T2) *Snode[T2] {
	var res, ptr *Snode[T2]
	for onode := root; onode != nil; onode = onode.Next {
		nnode := &Snode[T2]{Val: fnc(onode.Val), Next: nil}
		if onode == root {
			res = nnode
		} else {
			ptr.Next = nnode
		}
		ptr = nnode
	}
	return res
}

func MapA[T1, T2 any](root *Snode[T1], fnc func(T1) T2) *Snode[T2] {
	llen := ListSize(root)
	chans := make([]chan T2, llen)
	i := 0
	for node := root; node != nil; node = node.Next {
		elemChan := make(chan T2)
		chans[i] = elemChan
		i++
		go func(elemChan chan<- T2, el T1) {
			elemChan <- fnc(el)
		}(elemChan, node.Val)
	}
	var res, ptr *Snode[T2]
	for i, elemChan := range chans {
		node := &Snode[T2]{Val: <-elemChan, Next: nil}
		if i == 0 {
			res = node
		} else {
			ptr.Next = node
		}
		ptr = node
	}
	return res
}

func Reduce[T1, T2 any](root *Snode[T1], fnc func(T2, T1) T2, acc T2) T2 {
	for node := root; node != nil; node = node.Next {
		acc = fnc(acc, node.Val)
	}
	return acc
}

func ReduceR[T1, T2 any](root *Snode[T1], fnc func(T2, T1) T2, acc T2) T2 {
	if root == nil {
		return acc
	}
	return fnc(ReduceR(root.Next, fnc, acc), root.Val)
}

func Filter[T any](root *Snode[T], fnc func(T) bool) *Snode[T] {
	var res, ptr *Snode[T]
	for onode := root; onode != nil; onode = onode.Next {
		node := &Snode[T]{Val: onode.Val, Next: nil}
		if fnc(node.Val) {
			if res == nil {
				res = node
			} else {
				ptr.Next = node
			}
			ptr = node
		}
	}
	return res
}
