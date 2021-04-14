package chapter5

// 线性表元素类型，当前设定为int
type ElemType int

// 结点
type Node interface {
	// 获取值
	GetValue() (ElemType, error)
	// 设置值
	SetValue(e ElemType) error
	// 获取双亲结点
	GetParent() (Node, error)
	// 获取最左孩子
	GetLeftChild() (Node, error)
	// 获取右侧兄弟
	GetRightSibling() (Node, error)
	// 给结点插入第i个子树，值为e
	InsertChild(i int, e ElemType) error
	// 删除结点的第i个子树
	DeleteChild(i int) error
}

// 树
type Tree interface {
	// 初始化
	InitTree() error
	// 销毁
	DestroyTree() error
	// 清空
	ClearTree() error
	// 是否为空
	TreeEmpty() bool
	// 树的深度
	TreeDepth() int
	// 树的度
	TreeDegree() int
	// 获取根结点
	GetRoot() (Node, error)
	// 获取结点值
	Value(n Node) (ElemType, error)
	// 设置结点值
	Assign(n Node, e ElemType) error
	// 获取指定结点的双亲结点
	Parent(n Node) (Node, error)
	// 获取指定结点的最左孩子
	LeftChild(n Node) (Node, error)
	// 获取指定结点的右侧兄弟
	RightSibling(n Node) (Node, error)
	// 给指定结点插入第i个子树，值为e
	InsertChild(n Node, i int, e ElemType) error
	// 删除指定结点的第i个子树
	DeleteChild(n Node, i int) error
}
