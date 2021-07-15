package threadedbinarytree

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jormin/play-with-data-structure/chapter5"
)

// 测试获取值
func TestThreadedNode_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		n       *ThreadedNode
		want    chapter5.ElemType
		wantErr bool
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
func TestThreadedNode_SetValue(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		e        chapter5.ElemType
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && !reflect.DeepEqual(tt.n, tt.wantThreadedNode) {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取双亲结点
func TestThreadedNode_GetParent(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && p.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", p.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取左侧孩子结点
func TestThreadedNode_GetLeftChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && l.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", l.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取右侧孩子结点
func TestThreadedNode_GetRightChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && r.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", r.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取兄弟结点
func TestThreadedNode_GetSibling(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		wantErr  bool
		wantThreadedNode *ThreadedNode
		wantPos  int
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
					if s.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
						t.Errorf("want node error, got %v, want %v", s.PrintInfo(), tt.wantThreadedNode.PrintInfo())
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
func TestThreadedNode_InsertChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		i        int
		e        chapter5.ElemType
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && tt.n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试删除子节点
func TestThreadedNode_DeleteChild(t *testing.T) {
	tests := []struct {
		name     string
		n        *ThreadedNode
		i        int
		wantErr  bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
			n: createThreadedNode(
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
			wantThreadedNode: createThreadedNode(
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
				if err == nil && tt.n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试字符串
func TestThreadedNode_String(t *testing.T) {
	tests := []struct {
		name string
		n    *ThreadedNode
		want string
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
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
			n: createThreadedNode(
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
func TestThreadedNode_PrintInfo(t *testing.T) {
	tests := []struct {
		name string
		n    *ThreadedNode
		want string
	}{
		{
			name: "01",
			n: createThreadedNode(
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
			n: createThreadedNode(
				createParams{
					V:   1,
					P:   nil,
					Pos: NoPos,
					S:   nil,
					L:   nil,
					R:   nil,
				},
			),
			want: `{"l_tag":"0","left_child":"","parent":"","r_tag":"0","right_child":"","self":"1","sibling":"","sibling_pos":"-1"}`,
		},
		{
			name: "03",
			n: createThreadedNode(
				createParams{
					V:   2,
					P:   1,
					Pos: Left,
					S:   3,
					L:   4,
					R:   5,
				},
			),
			want: `{"l_tag":"0","left_child":"4","parent":"1","r_tag":"0","right_child":"5","self":"2","sibling":"3","sibling_pos":"1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.n.PrintInfo()
				if got != tt.want {
					t.Errorf("want print info error, got %v, want %v", got, tt.want)
					return
				}
			},
		)
	}
}

// 空结点
var tn1 *ThreadedNode

// 只有根结点
var tn2 = &ThreadedNode{
	Data: 1,
}

// 根结点只有左子树
var tn3 = &ThreadedNode{
	Data: 1,
	LeftChild: &ThreadedNode{
		Data: 2,
	},
}

// 根结点有左子树和右子树
var tn4 = &ThreadedNode{
	Data: 1,
	LeftChild: &ThreadedNode{
		Data: 2,
	},
	RightChild: &ThreadedNode{
		Data: 3,
	},
}

// 完整结点
var tn5 = &ThreadedNode{
	Data: 1,
	LeftChild: &ThreadedNode{
		Data: 2,
		LeftChild: &ThreadedNode{
			Data: 4,
			LeftChild: &ThreadedNode{
				Data: 8,
			},
			RightChild: &ThreadedNode{
				Data: 9,
			},
		},
		RightChild: &ThreadedNode{
			Data: 5,
			LeftChild: &ThreadedNode{
				Data: 10,
			},
			RightChild: &ThreadedNode{
				Data: 11,
			},
		},
	},
	RightChild: &ThreadedNode{
		Data: 3,
		LeftChild: &ThreadedNode{
			Data: 6,
			LeftChild: &ThreadedNode{
				Data: 12,
			},
			RightChild: &ThreadedNode{
				Data: 13,
			},
		},
		RightChild: &ThreadedNode{
			Data: 7,
			LeftChild: &ThreadedNode{
				Data: 14,
			},
			RightChild: &ThreadedNode{
				Data: 15,
			},
		},
	},
}

// 测试线索前序遍历
func TestThreadedNode_ThreadInfoDLR(t *testing.T) {
	tests := []struct {
		name string
		n    *ThreadedNode
		want string
	}{
		{
			name: "01",
			n:    tn1,
			want: "[]",
		},
		{
			name: "02",
			n:    tn2,
			want: "[1]",
		},
		{
			name: "03",
			n:    tn3,
			want: "[1 2]",
		},
		{
			name: "04",
			n:    tn4,
			want: "[1 2 3]",
		},
		{
			name: "05",
			n:    tn5,
			want: "[1 2 4 8 9 5 10 11 3 6 12 13 7 14 15]",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				DealThreadedNode(tt.n)
				// 处理线索数据
				_ = DLR(tt.n, new(ThreadData))
				ns := tt.n.ThreadInfoDLR()
				got := fmt.Sprintf("%v", ns)
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
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
func createThreadedNode(params createParams) *ThreadedNode {
	var p, n, s, l, r *ThreadedNode
	if params.V == nil {
		return nil
	}
	n = &ThreadedNode{
		Data: chapter5.ElemType(params.V.(int)),
	}
	if params.S != nil {
		s = &ThreadedNode{
			Data: chapter5.ElemType(params.S.(int)),
		}
	}
	if params.P != nil {
		var pl, pr *ThreadedNode
		if params.Pos == Left {
			pl = n
			pr = s
		} else {
			pl = s
			pr = n
		}
		p = &ThreadedNode{
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
		l = &ThreadedNode{
			Data:   chapter5.ElemType(params.L.(int)),
			Parent: n,
		}
	}
	if params.R != nil {
		r = &ThreadedNode{
			Data:   chapter5.ElemType(params.R.(int)),
			Parent: n,
		}
	}
	n.LeftChild = l
	n.RightChild = r
	return n
}