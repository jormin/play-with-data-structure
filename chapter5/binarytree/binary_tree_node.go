package binarytree

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jormin/play-with-data-structure/chapter5"
)

const (
	Left  = 0  // 左侧
	Right = 1  // 右侧
	NoPos = -1 // 不存在
)

// 二叉树结点
type Node struct {
	Data       chapter5.ElemType `json:"data" remark:"数据"`
	Parent     *Node             `json:"parent" remark:"双亲结点"`
	LeftChild  *Node             `json:"left_child" remark:"左侧孩子"`
	RightChild *Node             `json:"right_child" remark:"右侧孩子"`
}

// 获取值
func (n *Node) GetValue() (chapter5.ElemType, error) {
	if n == nil {
		return 0, errors.New("node not exists")
	}
	return n.Data, nil
}

// 设置值
func (n *Node) SetValue(e chapter5.ElemType) error {
	if n == nil {
		return errors.New("node not exists")
	}
	n.Data = e
	return nil
}

// 获取双亲结点
func (n *Node) GetParent() (*Node, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.Parent != nil {
		return n.Parent, nil
	}
	return nil, errors.New("parent not exists")
}

// 获取最左孩子
func (n *Node) GetLeftChild() (*Node, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.LeftChild == nil {
		return nil, errors.New("left child not exists")
	}
	return n.LeftChild, nil
}

// 获取最右孩子
func (n *Node) GetRightChild() (*Node, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.RightChild == nil {
		return nil, errors.New("right child not exists")
	}
	return n.RightChild, nil
}

// 获取左侧兄弟
func (n *Node) GetSibling() (*Node, int, error) {
	if n == nil {
		return nil, NoPos, errors.New("node not exists")
	}
	// 判断是否有双亲结点
	if n.Parent == nil {
		return nil, NoPos, errors.New("parent not exists")
	}
	var s *Node
	var pos int
	if n.Parent.LeftChild == n {
		s = n.Parent.RightChild
		pos = Right
	} else {
		s = n.Parent.LeftChild
		pos = Left
	}
	if s == nil {
		return nil, NoPos, errors.New("sibling node not exists")
	}
	return s, pos, nil
}

// 给结点插入第i个子树，值为e
func (n *Node) InsertChild(i int, e chapter5.ElemType) error {
	if n == nil {
		return errors.New("node not exists")
	}
	if i > 1 {
		return errors.New("a binary tree node has at most two child nodes")
	}
	tn := &Node{
		Data:       e,
		Parent:     n,
		LeftChild:  nil,
		RightChild: nil,
	}
	switch i {
	case 0:
		if n.LeftChild != nil {
			return errors.New("left child node exists")
		}
		n.LeftChild = tn
	case 1:
		if n.RightChild != nil {
			return errors.New("right child node exists")
		}
		n.RightChild = tn
	}
	return nil
}

// 删除结点的第i个子树
func (n *Node) DeleteChild(i int) error {
	if n == nil {
		return errors.New("node not exists")
	}
	if i > 1 {
		return errors.New("a binary tree node has at most two child nodes")
	}
	switch i {
	case 0:
		if n.LeftChild == nil {
			return errors.New("left child not exists")
		}
		n.LeftChild = nil
	case 1:
		if n.RightChild == nil {
			return errors.New("right child not exists")
		}
		n.RightChild = nil
	}
	return nil
}

// 字符串
func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v", n.Data)
}

// 打印
func (n *Node) PrintInfo() string {
	if n == nil {
		return ""
	}
	// 双亲结点
	p, _ := n.GetParent()
	// 兄弟结点
	s, pos, _ := n.GetSibling()
	d := map[string]string{
		"parent":      p.String(),
		"self":        n.String(),
		"sibling":     s.String(),
		"sibling_pos": fmt.Sprintf("%v", pos),
		"left_child":  n.LeftChild.String(),
		"right_child": n.RightChild.String(),
	}
	b, _ := json.Marshal(d)
	return fmt.Sprintf("%s", b)
}
