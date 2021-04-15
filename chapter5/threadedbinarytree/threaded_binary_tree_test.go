package threadedbinarytree

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.wcxst.com/jormin/play-with-data-structure/chapter5"
)

// 测试初始化
func TestThreadedBinaryTree_InitTree(t *testing.T) {
	tests := []struct {
		name    string
		b       *ThreadedBinaryTree
		wantErr bool
	}{
		{
			name:    "01",
			b:       new(ThreadedBinaryTree),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.InitTree()
				if (err != nil) != tt.wantErr {
					t.Errorf("InitTree() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

// 测试销毁
func TestThreadedBinaryTree_DestroyTree(t *testing.T) {
	tests := []struct {
		name                   string
		b                      *ThreadedBinaryTree
		wantErr                bool
		wantThreadedBinaryTree *ThreadedBinaryTree
	}{
		{
			name: "01",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			wantErr:                false,
			wantThreadedBinaryTree: createThreadedBinaryTree(0, 0, nil),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.DestroyTree()
				if (err != nil) != tt.wantErr {
					t.Errorf("DestroyTree() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(tt.b, tt.wantThreadedBinaryTree) {
					t.Errorf("want binary tree error, got %v, wantErr %v", tt.b, tt.wantThreadedBinaryTree)
				}
			},
		)
	}
}

// 测试清空
func TestThreadedBinaryTree_ClearTree(t *testing.T) {
	tests := []struct {
		name                   string
		b                      *ThreadedBinaryTree
		wantErr                bool
		wantThreadedBinaryTree *ThreadedBinaryTree
	}{
		{
			name: "01",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			wantErr:                false,
			wantThreadedBinaryTree: createThreadedBinaryTree(0, 0, nil),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.ClearTree()
				if (err != nil) != tt.wantErr {
					t.Errorf("ClearTree() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(tt.b, tt.wantThreadedBinaryTree) {
					t.Errorf("want binary tree error, got %v, wantErr %v", tt.b, tt.wantThreadedBinaryTree)
				}
			},
		)
	}
}

// 测试判断是否为空
func TestThreadedBinaryTree_TreeEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want bool
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.b.TreeEmpty()
				if got != tt.want {
					t.Errorf("want res error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试树深度
func TestThreadedBinaryTree_TreeDepth(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want int
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: 0,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				0, 1, &ThreadedNode{
					Data: 1,
				},
			),
			want: 1,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.b.TreeDepth()
				if got != tt.want {
					t.Errorf("want res error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试树的度
func TestThreadedBinaryTree_TreeDegree(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want int
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: 0,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				0, 1, &ThreadedNode{
					Data: 1,
				},
			),
			want: 0,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
				},
			),
			want: 1,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.b.TreeDegree()
				if got != tt.want {
					t.Errorf("want res error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试获取根结点
func TestThreadedBinaryTree_GetRoot(t *testing.T) {
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:             "01",
			b:                createThreadedBinaryTree(0, 0, nil),
			wantErr:          false,
			wantThreadedNode: nil,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
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
				n, err := tt.b.GetRoot()
				if (err != nil) != tt.wantErr {
					t.Errorf("ClearTree() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf(
						"want node print info error, got %v, want %v", n.PrintInfo(), tt.wantThreadedNode.PrintInfo(),
					)
				}
			},
		)
	}
}

// 测试获取结点的值
func TestThreadedBinaryTree_Value(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V: 2,
		},
	)
	tests := []struct {
		name    string
		b       *ThreadedBinaryTree
		n       *ThreadedNode
		wantErr bool
		want    chapter5.ElemType
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data:      1,
					LeftChild: n,
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			n:       n,
			wantErr: false,
			want:    2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := tt.b.Value(tt.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && got != tt.want {
					t.Errorf("want res error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试设置结点的值
func TestThreadedBinaryTree_Assign(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V: 2,
		},
	)
	tests := []struct {
		name                   string
		b                      *ThreadedBinaryTree
		n                      *ThreadedNode
		e                      chapter5.ElemType
		wantErr                bool
		wantThreadedBinaryTree *ThreadedBinaryTree
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data:      1,
					LeftChild: n,
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			n:       n,
			e:       20,
			wantErr: false,
			wantThreadedBinaryTree: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 20,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.Assign(tt.n, tt.e)
				if (err != nil) != tt.wantErr {
					t.Errorf("Assign() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && !reflect.DeepEqual(tt.b, tt.wantThreadedBinaryTree) {
					t.Errorf("want binary tree error, got %v, want %v", tt.b, tt.wantThreadedBinaryTree)
				}
			},
		)
	}
}

// 测试获取指定结点的双亲结点
func TestThreadedBinaryTree_Parent(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V: 2,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data:      1,
					LeftChild: n,
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			n:       n,
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
				n, err := tt.b.Parent(tt.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("Assign() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want parent node error, got %v, want %v", n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试获取指定节点的左侧孩子结点
func TestThreadedBinaryTree_LeftChild(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V:   1,
			P:   nil,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   3,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, n,
			),
			n:       n,
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
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				n, err := tt.b.LeftChild(tt.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("LeftChild() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf(
						"want left child node error, got %v, want %v", n.PrintInfo(), tt.wantThreadedNode.PrintInfo(),
					)
				}
			},
		)
	}
}

// 测试获取指定结点的右侧孩子结点
func TestThreadedBinaryTree_RightChild(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V:   1,
			P:   nil,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   3,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				2, 2, n,
			),
			n:       n,
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
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				n, err := tt.b.RightChild(tt.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("RightChild() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf(
						"want right child node error, got %v, want %v", n.PrintInfo(), tt.wantThreadedNode.PrintInfo(),
					)
				}
			},
		)
	}
}

// 测试获取指定结点的兄弟结点
func TestThreadedBinaryTree_Sibling(t *testing.T) {
	n1 := createThreadedNode(
		createParams{
			V:   2,
			P:   1,
			Pos: Left,
			S:   nil,
			L:   nil,
			R:   nil,
		},
	)
	n2 := createThreadedNode(
		createParams{
			V:   2,
			P:   1,
			Pos: Left,
			S:   3,
			L:   nil,
			R:   nil,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				1, 2,
				&ThreadedNode{
					Data:       1,
					LeftChild:  n1,
					RightChild: nil,
				},
			),
			n:       n1,
			wantErr: true,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				2, 2,
				&ThreadedNode{
					Data:      1,
					LeftChild: n2,
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			n:       n2,
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
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				n, _, err := tt.b.Sibling(tt.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("Sibling() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf(
						"want right child node error, got %v, want %v", n.PrintInfo(), tt.wantThreadedNode.PrintInfo(),
					)
				}
			},
		)
	}
}

// 测试给指定结点添加结点
func TestThreadedBinaryTree_InsertChild(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   3,
		},
	)
	n1 := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   nil,
		},
	)
	n2 := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   nil,
			R:   3,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		i                int
		e                chapter5.ElemType
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				1, 2, n,
			),
			n:       n,
			i:       2,
			wantErr: true,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, n,
			),
			n:       n,
			i:       Left,
			wantErr: true,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, n,
			),
			n:       n,
			i:       Right,
			wantErr: true,
		},
		{
			name: "05",
			b: createThreadedBinaryTree(
				2, 2, n1,
			),
			n:                n1,
			i:                Right,
			e:                3,
			wantErr:          false,
			wantThreadedNode: n,
		},
		{
			name: "06",
			b: createThreadedBinaryTree(
				2, 2, n2,
			),
			n:                n2,
			i:                Left,
			e:                2,
			wantErr:          false,
			wantThreadedNode: n,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.InsertChild(tt.n, tt.i, tt.e)
				if (err != nil) != tt.wantErr {
					t.Errorf("InsertChild() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && tt.n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试指定结点删除结点
func TestThreadedBinaryTree_DeleteChild(t *testing.T) {
	n := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   3,
		},
	)
	n1 := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   2,
			R:   nil,
		},
	)
	n2 := createThreadedNode(
		createParams{
			V:   1,
			P:   0,
			Pos: NoPos,
			S:   nil,
			L:   nil,
			R:   3,
		},
	)
	tests := []struct {
		name             string
		b                *ThreadedBinaryTree
		n                *ThreadedNode
		i                int
		wantErr          bool
		wantThreadedNode *ThreadedNode
	}{
		{
			name:    "01",
			b:       createThreadedBinaryTree(0, 0, nil),
			wantErr: true,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				1, 2, n,
			),
			n:       n,
			i:       2,
			wantErr: true,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, n2,
			),
			n:       n2,
			i:       Left,
			wantErr: true,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, n1,
			),
			n:       n1,
			i:       Right,
			wantErr: true,
		},
		{
			name: "05",
			b: createThreadedBinaryTree(
				2, 2, n,
			),
			n:                n,
			i:                Left,
			wantErr:          false,
			wantThreadedNode: n2,
		},
		{
			name: "06",
			b: createThreadedBinaryTree(
				2, 2, n,
			),
			n:                n,
			i:                Right,
			wantErr:          false,
			wantThreadedNode: n,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.b.DeleteChild(tt.n, tt.i)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeleteChild() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && tt.n.PrintInfo() != tt.wantThreadedNode.PrintInfo() {
					t.Errorf("want node error, got %v, want %v", tt.n.PrintInfo(), tt.wantThreadedNode.PrintInfo())
				}
			},
		)
	}
}

// 测试生成字符串
func TestThreadedBinaryTree_String(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want string
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: `{"degree":0,"depth":0,"nodes":"[]"}`,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				0, 1, &ThreadedNode{
					Data: 1,
				},
			),
			want: `{"degree":0,"depth":1,"nodes":"[1]"}`,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
				},
			),
			want: `{"degree":1,"depth":2,"nodes":"[1 2]"}`,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: `{"degree":2,"depth":2,"nodes":"[1 2 3]"}`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.b.String()
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试Json序列化
func TestThreadedBinaryTree_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want string
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: `{"degree":0,"depth":0,"nodes":"[]"}`,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				0, 1, &ThreadedNode{
					Data: 1,
				},
			),
			want: `{"degree":0,"depth":1,"nodes":"[1]"}`,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
				},
			),
			want: `{"degree":1,"depth":2,"nodes":"[1 2]"}`,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: `{"degree":2,"depth":2,"nodes":"[1 2 3]"}`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				b, _ := tt.b.MarshalJSON()
				got := fmt.Sprintf("%s", b)
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试统计空指针数量
func TestThreadedBinaryTree_CountEmptyPoints(t *testing.T) {
	tests := []struct {
		name string
		b    *ThreadedBinaryTree
		want int
	}{
		{
			name: "01",
			b:    createThreadedBinaryTree(0, 0, nil),
			want: 1,
		},
		{
			name: "02",
			b: createThreadedBinaryTree(
				0, 1, &ThreadedNode{
					Data: 1,
				},
			),
			want: 2,
		},
		{
			name: "03",
			b: createThreadedBinaryTree(
				1, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
				},
			),
			want: 3,
		},
		{
			name: "04",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
					},
					RightChild: &ThreadedNode{
						Data: 3,
					},
				},
			),
			want: 4,
		},
		{
			name: "05",
			b: createThreadedBinaryTree(
				2, 2, &ThreadedNode{
					Data: 1,
					LeftChild: &ThreadedNode{
						Data: 2,
						LeftChild: &ThreadedNode{
							Data: 4,
							LeftChild: &ThreadedNode{
								Data: 8,
							},
						},
						RightChild: &ThreadedNode{
							Data: 5,
							LeftChild: &ThreadedNode{
								Data: 10,
							},
						},
					},
					RightChild: &ThreadedNode{
						Data: 3,
						LeftChild: &ThreadedNode{
							Data: 6,
							RightChild: &ThreadedNode{
								Data: 13,
							},
						},
					},
				},
			),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.b.CountEmptyPoints(tt.b.Root)
				if got != tt.want {
					t.Errorf("want empty point num error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 创建二叉树
func createThreadedBinaryTree(dg int, dp int, n *ThreadedNode) *ThreadedBinaryTree {
	if n != nil {
		DealThreadedNode(n)
	}
	t := &ThreadedBinaryTree{
		Root:   n,
		Degree: dg,
		Depth:  dp,
	}
	return t
}

// 处理结点
func DealThreadedNode(n *ThreadedNode) {
	if n == nil || (n.LeftChild == nil && n.RightChild == nil) {
		return
	}
	if n.LeftChild != nil {
		n.LeftChild.Parent = n
		DealThreadedNode(n.LeftChild)
	}
	if n.RightChild != nil {
		n.RightChild.Parent = n
		DealThreadedNode(n.RightChild)
	}
}
