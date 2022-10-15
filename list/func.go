package list

func Map[T1, T2 any](root *Node[T1], fnc func(T1) T2) *Node[T2] {
	if root == nil {
		return nil
	}
	var res, ptr *Node[T2]
	for onode := root; onode != nil; onode = onode.Next {
		nnode := &Node[T2]{Val: fnc(onode.Val), Next: nil}
		if onode == root {
			res = nnode
		} else {
			ptr.Next = nnode
		}
		ptr = nnode
	}
	return res
}

func MapA[T1, T2 any](root *Node[T1], fnc func(T1) T2) *Node[T2] {
	if root == nil {
		return nil
	}
	llen := ListSize(root)
	chans := make(chan chan T2, llen)
	for node := root; node != nil; node = node.Next {
		elemChan := make(chan T2)
		chans <- elemChan
		go func(elemChan chan<- T2, el T1) {
			elemChan <- fnc(el)
		}(elemChan, node.Val)
	}
	close(chans)
	var res, ptr *Node[T2]
	i := 0
	for elemChan := range chans {
		node := &Node[T2]{Val: <-elemChan, Next: nil}
		if i == 0 {
			res = node
		} else {
			ptr.Next = node
		}
		ptr = node
		i++
	}
	return res
}

func Reduce[T1, T2 any](root *Node[T1], fnc func(T2, T1) T2, acc T2) T2 {
	for node := root; node != nil; node = node.Next {
		acc = fnc(acc, node.Val)
	}
	return acc
}

func ReduceR[T1, T2 any](root *Node[T1], fnc func(T2, T1) T2, acc T2) T2 {
	if root == nil {
		return acc
	}
	return fnc(ReduceR(root.Next, fnc, acc), root.Val)
}

func Filter[T any](root *Node[T], fnc func(T) bool) *Node[T] {
	var res, ptr *Node[T]
	for onode := root; onode != nil; onode = onode.Next {
		node := &Node[T]{Val: onode.Val, Next: nil}
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
