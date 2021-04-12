package chapter4

import (
	"reflect"
	"testing"
)

// 测试初始化
func TestLinkedStack_InitStack(t *testing.T) {
	tests := []struct {
		name string
		s    *LinkedStack
		want Status
	}{
		{
			name: "01",
			s:    new(LinkedStack),
			want: OK,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.InitStack()
				if got != tt.want {
					t.Errorf("InitStack() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试清理
func TestLinkedStack_ClearStack(t *testing.T) {
	tests := []struct {
		name       string
		s          *LinkedStack
		want       Status
		wantTop    *Node
		wantLength int
	}{
		{
			name:       "01",
			s:          makeLinkedStack(1, 2, 3, 4),
			want:       OK,
			wantTop:    nil,
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.ClearStack()
				if got != tt.want {
					t.Errorf("ClearStack() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if tt.s.Top != tt.wantTop {
						t.Errorf("want top error, got %v, want %v", tt.s.Top, tt.wantTop)
					}
					if tt.s.Length != tt.wantLength {
						t.Errorf("want length error, got %v, want %v", tt.s.Length, tt.wantLength)
					}
				}
			},
		)
	}
}

// 测试销毁
func TestLinkedStack_DestroyStack(t *testing.T) {
	tests := []struct {
		name       string
		s          *LinkedStack
		want       Status
		wantTop    *Node
		wantLength int
	}{
		{
			name:       "01",
			s:          makeLinkedStack(1, 2, 3, 4),
			want:       OK,
			wantTop:    nil,
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.DestroyStack()
				if got != tt.want {
					t.Errorf("DestroyStack() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if tt.s.Top != tt.wantTop {
						t.Errorf("want top error, got %v, want %v", tt.s.Top, tt.wantTop)
					}
					if tt.s.Length != tt.wantLength {
						t.Errorf("want length error, got %v, want %v", tt.s.Length, tt.wantLength)
					}
				}
			},
		)
	}
}

// 测试获取栈顶
func TestLinkedStack_GetTop(t *testing.T) {
	tests := []struct {
		name     string
		s        *LinkedStack
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			s:    makeLinkedStack(),
			want: Error,
		},
		{
			name:     "02",
			s:        makeLinkedStack(1, 2, 3, 4),
			want:     OK,
			wantElem: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.s.GetTop(&e)
				if got != tt.want {
					t.Errorf("GetTop() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if e != tt.wantElem {
						t.Errorf("want elem error, got %v, want %v", e, tt.wantElem)
					}
				}
			},
		)
	}
}

// 测试入栈
func TestLinkedStack_Push(t *testing.T) {
	tests := []struct {
		name      string
		s         *LinkedStack
		e         ElemType
		want      Status
		wantStack *LinkedStack
	}{
		{
			name:      "01",
			s:         makeLinkedStack(1, 2, 3, 4),
			e:         5,
			want:      OK,
			wantStack: makeLinkedStack(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.Push(tt.e)
				if got != tt.want {
					t.Errorf("Push() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if !reflect.DeepEqual(tt.s, tt.wantStack) {
						t.Errorf("want stack error, got %v, want %v", tt.s, tt.wantStack)
					}
				}
			},
		)
	}
}

// 测试出栈
func TestLinkedStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		s         *LinkedStack
		want      Status
		wantTop   ElemType
		wantStack *LinkedStack
	}{
		{
			name: "01",
			s:    makeLinkedStack(),
			want: Error,
		},
		{
			name:      "02",
			s:         makeLinkedStack(1, 2, 3, 4),
			want:      OK,
			wantTop:   4,
			wantStack: makeLinkedStack(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.s.Pop(&e)
				if got != tt.want {
					t.Errorf("Pop() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if e != tt.wantTop {
						t.Errorf("want top error, got %v, want %v", e, tt.wantTop)
					}
					if !reflect.DeepEqual(tt.s, tt.wantStack) {
						t.Errorf("want stack error, got %v, want %v", tt.s, tt.wantStack)
					}
				}
			},
		)
	}
}

// 测试判断栈是否为空
func TestLinkedStack_StackEmpty(t *testing.T) {
	type fields struct {
		Top    *Node
		Length int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s := &LinkedStack{
					Top:    tt.fields.Top,
					Length: tt.fields.Length,
				}
				if got := s.StackEmpty(); got != tt.want {
					t.Errorf("StackEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试获取栈长度
func TestLinkedStack_StackLength(t *testing.T) {
	tests := []struct {
		name       string
		s          *LinkedStack
		wantLength int
	}{
		{
			name:       "01",
			s:          makeLinkedStack(),
			wantLength: 0,
		},
		{
			name:       "02",
			s:          makeLinkedStack(1, 2, 3, 4),
			wantLength: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				n := tt.s.StackLength()
				if n != tt.wantLength {
					t.Errorf("want length error, got %v, want %v", n, tt.wantLength)
				}
			},
		)
	}
}

// 生成链式栈
func makeLinkedStack(es ...ElemType) *LinkedStack {
	s := new(LinkedStack)
	if len(es) == 0 {
		s.InitStack()
		return s
	}
	for _, v := range es {
		s.Top = &Node{
			Value: v,
			Next:  s.Top,
		}
		s.Length++
	}
	return s
}
