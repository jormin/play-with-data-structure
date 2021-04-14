package binarytree

import (
	"reflect"
	"testing"

	"gitlab.wcxst.com/jormin/play-with-data-structure/chapter5"
)

// 测试获取值
func TestNode_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		n       *Node
		want    chapter5.ElemType
		wantErr bool
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want:    1,
			wantErr: false,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := tt.n.GetValue()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && got != tt.want {
					t.Errorf("GetValue() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试设置值
func TestNode_SetValue(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		e        chapter5.ElemType
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			e:       2,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			e:       3,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   3,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.n.SetValue(tt.e)
				if (err != nil) != tt.wantErr {
					t.Errorf("SetValue() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && !reflect.DeepEqual(tt.n, tt.wantNode) {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取双亲结点
func TestNode_GetParent(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   2,
					R:   3,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p, err := tt.n.GetParent()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetParent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && p.PrintInfo() != tt.wantNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", p.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取左侧孩子结点
func TestNode_GetLeftChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   4,
					P:   2,
					Pos: Left,
					S:   5,
					L:   nil,
					R:   nil,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l, err := tt.n.GetLeftChild()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetLeftChild() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && l.PrintInfo() != tt.wantNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", l.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取右侧孩子结点
func TestNode_GetRightChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   5,
					P:   2,
					Pos: Right,
					S:   4,
					L:   nil,
					R:   nil,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r, err := tt.n.GetRightChild()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetRightChild() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && r.PrintInfo() != tt.wantNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", r.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取兄弟结点
func TestNode_GetSibling(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		wantErr  bool
		wantNode *Node
		wantPos  int
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   3,
					P:   1,
					Pos: Right,
					S:   2,
					L:   nil,
					R:   nil,
				},
			),
			wantPos: Right,
		},
		{
			name: "04",
			n: createNode(
				createParams{
					V:   3,
					P:   1,
					Pos: Right,
					S:   2,
					L:   4,
					R:   5,
				},
			),
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   nil,
					R:   nil,
				},
			),
			wantPos: Left,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s, pos, err := tt.n.GetSibling()
				if (err != nil) != tt.wantErr {
					t.Errorf("GetSibling() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil {
					if s.PrintInfo() != tt.wantNode.PrintInfo() {
						t.Errorf("want node error, got %v, want %v", s.PrintInfo(), tt.wantNode.PrintInfo())
					}
					if pos != tt.wantPos {
						t.Errorf("want pos error, got %v, want %v", pos, tt.wantPos)
					}
				}
			},
		)
	}
}

// 测试添加子节点
func TestNode_InsertChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		i        int
		e        chapter5.ElemType
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),

			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			i:       2,
			e:       1,
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			i:       Left,
			e:       1,
			wantErr: true,
		},
		{
			name: "04",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			i:       Right,
			e:       1,
			wantErr: true,
		},
		{
			name: "05",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   nil,
					R:   5,
				},
			),
			i:       Left,
			e:       4,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
		},
		{
			name: "06",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   nil,
				},
			),
			i:       Right,
			e:       5,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.n.InsertChild(tt.i, tt.e)
				if (err != nil) != tt.wantErr {
					t.Errorf("InsertChild() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && tt.n.PrintInfo() != tt.wantNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试删除子节点
func TestNode_DeleteChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *Node
		i        int
		wantErr  bool
		wantNode *Node
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),

			wantErr: true,
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			i:       2,
			wantErr: true,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   nil,
					R:   5,
				},
			),
			i:       Left,
			wantErr: true,
		},
		{
			name: "04",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   nil,
				},
			),
			i:       Right,
			wantErr: true,
		},
		{
			name: "05",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			i:       Left,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   nil,
					R:   5,
				},
			),
		},
		{
			name: "06",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			i:       Right,
			wantErr: false,
			wantNode: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   nil,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.n.DeleteChild(tt.i)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeleteChild() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && tt.n.PrintInfo() != tt.wantNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantNode.PrintInfo())
				}
			},
		)
	}
}

// 测试字符串
func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want: "",
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want: "1",
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.n.String()
				if got != tt.want {
					t.Errorf("want string error, got %v, wantErr %v", got, tt.want)
					return
				}
			},
		)
	}
}

// 测试打印信息
func TestNode_PrintInfo(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{
			name: "01",
			n: createNode(
				createParams{
					V:   nil,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want: "",
		},
		{
			name: "02",
			n: createNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want: `{"left_child":"","parent":"","right_child":"","self":"1","sibling":"","sibling_pos":"-1"}`,
		},
		{
			name: "03",
			n: createNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			want: `{"left_child":"4","parent":"1","right_child":"5","self":"2","sibling":"3","sibling_pos":"1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.n.PrintInfo()
				if got != tt.want {
					t.Errorf("want print info error, got %v, wantErr %v", got, tt.want)
					return
				}
			},
		)
	}
}

// 创建参数
type createParams struct {
	V   interface{} `json:"v" remark:"值"`
	P   interface{} `json:"p" remark:"双亲的值"`
	Pos int         `json:"pos" remark:"双亲的左子树(0)还是右子树(1)"`
	S   interface{} `json:"s" remark:"兄弟结点的值"`
	L   interface{} `json:"l" remark:"左侧孩子的值"`
	R   interface{} `json:"r" remark:"右侧孩子的值"`
}

// 生成结点
func createNode(params createParams) *Node {
	var p, n, s, l, r *Node
	if params.V == nil {
		return nil
	}
	n = &Node{
		Data: chapter5.ElemType(params.V.(int)),
	}
	if params.S != nil {
		s = &Node{
			Data: chapter5.ElemType(params.S.(int)),
		}
	}
	if params.P != nil {
		var pl, pr *Node
		if params.Pos == Left {
			pl = n
			pr = s
		} else {
			pl = s
			pr = n
		}
		p = &Node{
			Data:       chapter5.ElemType(params.P.(int)),
			Parent:     nil,
			LeftChild:  pl,
			RightChild: pr,
		}
		n.Parent = p
		if s != nil {
			s.Parent = p
		}
	}
	if params.L != nil {
		l = &Node{
			Data:   chapter5.ElemType(params.L.(int)),
			Parent: n,
		}
	}
	if params.R != nil {
		r = &Node{
			Data:   chapter5.ElemType(params.R.(int)),
			Parent: n,
		}
	}
	n.LeftChild = l
	n.RightChild = r
	return n
}
