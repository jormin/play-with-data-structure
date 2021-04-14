package binarytree

// 前序遍历
func PreOrderTraverse(n *Node,) []*Node {
	var ns []*Node
	if n == nil {
		return ns
	}
	ns = append(ns, n)
	ns = append(ns, PreOrderTraverse(n.LeftChild)...)
	ns = append(ns, PreOrderTraverse(n.RightChild)...)
	return ns
}
