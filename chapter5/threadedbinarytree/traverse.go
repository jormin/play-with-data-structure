package threadedbinarytree

// 线索化数据
type ThreadData struct {
	pre *ThreadedNode
}

// 前序遍历（Data、LChild、RChild）
func DLR(n *ThreadedNode, td *ThreadData) []*ThreadedNode {
	var ns []*ThreadedNode
	if n == nil {
		return ns
	}
	threadAble(n, td)
	ns = append(ns, n)
	if n.LTag == TagNode {
		ns = append(ns, DLR(n.LeftChild, td)...)
	}
	if n.RTag == TagNode {
		ns = append(ns, DLR(n.RightChild, td)...)
	}
	return ns
}

// 中序遍历（LChild、Data、RChild）
func LDR(n *ThreadedNode, td *ThreadData) []*ThreadedNode {
	var ns []*ThreadedNode
	if n == nil {
		return ns
	}
	ns = append(ns, LDR(n.LeftChild, td)...)
	threadAble(n, td)
	ns = append(ns, n)
	ns = append(ns, LDR(n.RightChild, td)...)
	return ns
}

// 后续遍历（LChild、RChild、Data）
func LRD(n *ThreadedNode, td *ThreadData) []*ThreadedNode {
	var ns []*ThreadedNode
	if n == nil {
		return ns
	}
	ns = append(ns, LRD(n.LeftChild, td)...)
	ns = append(ns, LRD(n.RightChild, td)...)
	threadAble(n, td)
	ns = append(ns, n)
	return ns
}

// 线索化
func threadAble(n *ThreadedNode, td *ThreadData) {
	if td.pre == nil {
		td.pre = n
		return
	}
	// 如果当前结点左子树为空
	if n.LeftChild == nil {
		// 左标记设置为线索
		n.LTag = TagThread
		// 左子树设置为前驱结点
		n.LeftChild = td.pre
	} else {
		n.LTag = TagNode
	}
	// 如果前结点的右子树为空
	if td.pre.RightChild == nil {
		// 右标记设置为线索
		td.pre.RTag = TagThread
		// 右子树设置为后续结点
		td.pre.RightChild = n
	} else {
		n.RTag = TagNode
	}
	td.pre = n
	return
}
