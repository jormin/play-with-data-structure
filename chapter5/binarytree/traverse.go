package binarytree

// 前序遍历（Data、LChild、RChild）
func DLR(n *Node) []*Node {
	var ns []*Node
	if n == nil {
		return ns
	}
	ns = append(ns, n)
	ns = append(ns, DLR(n.LeftChild)...)
	ns = append(ns, DLR(n.RightChild)...)
	return ns
}

// 中序遍历（LChild、Data、RChild）
func LDR(n *Node) []*Node {
	var ns []*Node
	if n == nil {
		return ns
	}
	ns = append(ns, LDR(n.LeftChild)...)
	ns = append(ns, n)
	ns = append(ns, LDR(n.RightChild)...)
	return ns
}

// 后续遍历（LChild、RChild、Data）
func LRD(n *Node) []*Node {
	var ns []*Node
	if n == nil {
		return ns
	}
	ns = append(ns, LRD(n.LeftChild)...)
	ns = append(ns, LRD(n.RightChild)...)
	ns = append(ns, n)
	return ns
}
