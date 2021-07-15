package threadedbinarytree

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jormin/play-with-data-structure/chapter5"
	"github.com/jormin/play-with-data-structure/chapter5/binarytree"
)

const (
	TagNode   = 0 // 代表指针域指向的是左右结点
	TagThread = 1 // 代表指针域指向的是前驱或者后继
)

const (
	Left  = 0  // 左侧
	Right = 1  // 右侧
	NoPos = -1 // 不存在
)

// 二叉树结点
type ThreadedNode struct {
	Data       chapter5.ElemType `json:"data" remark:"数据"`
	Parent     *ThreadedNode     `json:"parent" remark:"双亲结点"`
	LeftChild  *ThreadedNode     `json:"left_child" remark:"左侧孩子"`
	RightChild *ThreadedNode     `json:"right_child" remark:"右侧孩子"`
	LTag       int               `json:"l_tag" remark:"左标记"`
	RTag       int               `json:"r_tag" remark:"右标记"`
}

// 获取值
func (n *ThreadedNode) GetValue() (chapter5.ElemType, error) {
	if n == nil {
		return 0, errors.New("node not exists")
	}
	return n.Data, nil
}

// 设置值
func (n *ThreadedNode) SetValue(e chapter5.ElemType) error {
	if n == nil {
		return errors.New("node not exists")
	}
	n.Data = e
	return nil
}

// 获取双亲结点
func (n *ThreadedNode) GetParent() (*ThreadedNode, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.Parent != nil {
		return n.Parent, nil
	}
	return nil, errors.New("parent not exists")
}

// 获取最左孩子
func (n *ThreadedNode) GetLeftChild() (*ThreadedNode, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.LeftChild == nil {
		return nil, errors.New("left child not exists")
	}
	return n.LeftChild, nil
}

// 获取最右孩子
func (n *ThreadedNode) GetRightChild() (*ThreadedNode, error) {
	if n == nil {
		return nil, errors.New("node not exists")
	}
	if n.RightChild == nil {
		return nil, errors.New("right child not exists")
	}
	return n.RightChild, nil
}

// 获取左侧兄弟
func (n *ThreadedNode) GetSibling() (*ThreadedNode, int, error) {
	if n == nil {
		return nil, binarytree.NoPos, errors.New("node not exists")
	}
	// 判断是否有双亲结点
	if n.Parent == nil {
		return nil, binarytree.NoPos, errors.New("parent not exists")
	}
	var s *ThreadedNode
	var pos int
	if n.Parent.LeftChild == n {
		s = n.Parent.RightChild
		pos = binarytree.Right
	} else {
		s = n.Parent.LeftChild
		pos = binarytree.Left
	}
	if s == nil {
		return nil, binarytree.NoPos, errors.New("sibling node not exists")
	}
	return s, pos, nil
}

// 给结点插入第i个子树，值为e
func (n *ThreadedNode) InsertChild(i int, e chapter5.ElemType) error {
	if n == nil {
		return errors.New("node not exists")
	}
	if i > 1 {
		return errors.New("a binary tree node has at most two child nodes")
	}
	tn := &ThreadedNode{
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
func (n *ThreadedNode) DeleteChild(i int) error {
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
func (n *ThreadedNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v", n.Data)
}

// 打印
func (n *ThreadedNode) PrintInfo() string {
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
		"l_tag":       fmt.Sprintf("%d", n.LTag),
		"r_tag":       fmt.Sprintf("%d", n.RTag),
	}
	b, _ := json.Marshal(d)
	return fmt.Sprintf("%s", b)
}

// 前序遍历线索信息
func (n *ThreadedNode) ThreadInfoDLR() []*ThreadedNode {
	var ts []*ThreadedNode
	if n == nil {
		return ts
	}
	ts = append(ts, n)
	if n.LeftChild == nil {
		return ts
	}
	tmp := n.LeftChild
	for tmp != nil && tmp != n {
		ts = append(ts, tmp)
		for tmp.LTag == TagNode && tmp.LeftChild != nil {
			ts = append(ts, tmp.LeftChild)
			tmp = tmp.LeftChild
		}
		tmp = tmp.RightChild
	}
	return ts
}
