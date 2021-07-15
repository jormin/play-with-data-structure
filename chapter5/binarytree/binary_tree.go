package binarytree

import (
	"encoding/json"
	"fmt"

	"github.com/jormin/play-with-data-structure/chapter5"
)

// 二叉树
type BinaryTree struct {
	Root   *Node `json:"root" remark:"根结点"`
	Degree int   `json:"degree" remark:"度"`
	Depth  int   `json:"depth" remark:"深度"`
}

// 初始化
func (b *BinaryTree) InitTree() error {
	b.Root = nil
	b.Degree = 0
	b.Depth = 0
	return nil
}

// 销毁
func (b *BinaryTree) DestroyTree() error {
	return b.InitTree()
}

// 清空
func (b *BinaryTree) ClearTree() error {
	return b.InitTree()
}

// 是否为空
func (b *BinaryTree) TreeEmpty() bool {
	return b.Root == nil
}

// 树的深度
func (b *BinaryTree) TreeDepth() int {
	return b.Depth
}

// 树的度
func (b *BinaryTree) TreeDegree() int {
	return b.Degree
}

// 获取根结点
func (b *BinaryTree) GetRoot() (*Node, error) {
	return b.Root, nil
}

// 获取结点值
func (b *BinaryTree) Value(n *Node) (chapter5.ElemType, error) {
	return n.GetValue()
}

// 设置结点值
func (b *BinaryTree) Assign(n *Node, e chapter5.ElemType) error {
	return n.SetValue(e)
}

// 获取指定结点的双亲结点
func (b *BinaryTree) Parent(n *Node) (*Node, error) {
	return n.GetParent()
}

// 获取指定结点的最左孩子
func (b *BinaryTree) LeftChild(n *Node) (*Node, error) {
	return n.GetLeftChild()
}

// 获取指定结点的最右孩子
func (b *BinaryTree) RightChild(n *Node) (*Node, error) {
	return n.GetRightChild()
}

// 获取指定结点的左侧兄弟
func (b *BinaryTree) Sibling(n *Node) (*Node, int, error) {
	return n.GetSibling()
}

// 给指定结点插入第i个子树，值为e
func (b *BinaryTree) InsertChild(n *Node, i int, e chapter5.ElemType) error {
	return n.InsertChild(i, e)
}

// 删除指定结点的第i个子树
func (b *BinaryTree) DeleteChild(n *Node, i int) error {
	return n.DeleteChild(i)
}

// 统计空指针数
func (b *BinaryTree) CountEmptyPoints(n *Node) int {
	i := 0
	if n == nil {
		return 1
	}
	if n.LeftChild == nil {
		i++
	} else {
		i += b.CountEmptyPoints(n.LeftChild)
	}
	if n.RightChild == nil {
		i++
	} else {
		i += b.CountEmptyPoints(n.RightChild)
	}
	fmt.Println(n.PrintInfo(), i)
	return i
}

// 字符串
func (b *BinaryTree) String() string {
	bt, _ := json.Marshal(b)
	return fmt.Sprintf("%s", bt)
}

// Json序列化
func (b *BinaryTree) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"degree": b.Degree,
		"depth":  b.Depth,
	}
	var nodes []*Node
	if !b.TreeEmpty() {
		nodes = DLR(b.Root)
	}
	d["nodes"] = fmt.Sprintf("%v", nodes)
	return json.Marshal(d)
}
