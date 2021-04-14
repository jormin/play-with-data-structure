package binarytree

import (
	"fmt"
	"testing"
)

// 空结点
var n1 *Node

// 只有根结点
var n2 = &Node{
	Data: 1,
}

// 根结点只有左子树
var n3 = &Node{
	Data: 1,
	LeftChild: &Node{
		Data: 2,
	},
}

// 根结点有左子树和右子树
var n4 = &Node{
	Data: 1,
	LeftChild: &Node{
		Data: 2,
	},
	RightChild: &Node{
		Data: 3,
	},
}

// 完整结点
var n5 = &Node{
	Data: 1,
	LeftChild: &Node{
		Data: 2,
		LeftChild: &Node{
			Data: 4,
			LeftChild: &Node{
				Data: 8,
			},
			RightChild: &Node{
				Data: 9,
			},
		},
		RightChild: &Node{
			Data: 5,
			LeftChild: &Node{
				Data: 10,
			},
			RightChild: &Node{
				Data: 11,
			},
		},
	},
	RightChild: &Node{
		Data: 3,
		LeftChild: &Node{
			Data: 6,
			LeftChild: &Node{
				Data: 12,
			},
			RightChild: &Node{
				Data: 13,
			},
		},
		RightChild: &Node{
			Data: 7,
			LeftChild: &Node{
				Data: 14,
			},
			RightChild: &Node{
				Data: 15,
			},
		},
	},
}

// 测试前序遍历
func TestDLR(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{
			name: "01",
			n:    n1,
			want: "[]",
		},
		{
			name: "02",
			n:    n2,
			want: "[1]",
		},
		{
			name: "03",
			n:    n3,
			want: "[1 2]",
		},
		{
			name: "04",
			n:    n4,
			want: "[1 2 3]",
		},
		{
			name: "05",
			n:    n5,
			want: "[1 2 4 8 9 5 10 11 3 6 12 13 7 14 15]",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ns := DLR(tt.n)
				got := fmt.Sprintf("%v", ns)
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试中序遍历
func TestLDR(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{
			name: "01",
			n:    n1,
			want: "[]",
		},
		{
			name: "02",
			n:    n2,
			want: "[1]",
		},
		{
			name: "03",
			n:    n3,
			want: "[2 1]",
		},
		{
			name: "04",
			n:    n4,
			want: "[2 1 3]",
		},
		{
			name: "05",
			n:    n5,
			want: "[8 4 9 2 10 5 11 1 12 6 13 3 14 7 15]",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ns := LDR(tt.n)
				got := fmt.Sprintf("%v", ns)
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试后续遍历
func TestLRD(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{
			name: "01",
			n:    n1,
			want: "[]",
		},
		{
			name: "02",
			n:    n2,
			want: "[1]",
		},
		{
			name: "03",
			n:    n3,
			want: "[2 1]",
		},
		{
			name: "04",
			n:    n4,
			want: "[2 3 1]",
		},
		{
			name: "05",
			n:    n5,
			want: "[8 9 4 10 11 5 2 12 13 6 14 15 7 3 1]",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ns := LRD(tt.n)
				got := fmt.Sprintf("%v", ns)
				if got != tt.want {
					t.Errorf("want error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}
