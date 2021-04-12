package stack

import (
	"reflect"
	"testing"
)

// 测试初始化
func TestSequentialList_InitStack(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialList
		want Status
	}{
		{
			name: "01",
			s:    new(SequentialList),
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

// 测试清空
func TestSequentialList_ClearStack(t *testing.T) {
	tests := []struct {
		name    string
		s       *SequentialList
		want    Status
		wantTop int
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			want:    OK,
			wantTop: EmptyTop,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.ClearStack()
				if got != tt.want {
					t.Errorf("ClearStack() = %v, want %v", got, tt.want)
				}
				if got == OK && tt.s.Top != tt.wantTop {
					t.Errorf("want top error, got %v, want %v", tt.s.Top, tt.wantTop)
				}
			},
		)
	}
}

// 测试销毁
func TestSequentialList_DestroyStack(t *testing.T) {
	tests := []struct {
		name    string
		s       *SequentialList
		want    Status
		wantTop int
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			want:    OK,
			wantTop: EmptyTop,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.DestroyStack()
				if got != tt.want {
					t.Errorf("DestroyStack() = %v, want %v", got, tt.want)
				}
				if got == OK && tt.s.Top != tt.wantTop {
					t.Errorf("want top error, got %v, want %v", tt.s.Top, tt.wantTop)
				}
			},
		)
	}
}

// 测试获取栈顶
func TestSequentialList_GetTop(t *testing.T) {
	tests := []struct {
		name     string
		s        *SequentialList
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{},
				Top:  -1,
			},
			want: Error,
		},
		{
			name: "02",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
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
				if got == OK && e != tt.wantElem {
					t.Errorf("want elem error, got %v, want %v", e, tt.wantElem)
				}
			},
		)
	}
}

// 测试入栈
func TestSequentialList_Push(t *testing.T) {
	tests := []struct {
		name      string
		s         *SequentialList
		elem      ElemType
		want      Status
		wantStack *SequentialList
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				},
				Top: 19,
			},
			want: Error,
		},
		{
			name: "02",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			elem: 5,
			want: OK,
			wantStack: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4, 5},
				Top:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.Push(tt.elem)
				if got != tt.want {
					t.Errorf("Push() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.s, tt.wantStack) {
					t.Errorf("got stack error, got %s, want %s", tt.s, tt.wantStack)
				}
			},
		)
	}
}

// 测试出栈
func TestSequentialList_Pop(t *testing.T) {
	tests := []struct {
		name      string
		s         *SequentialList
		want      Status
		wantStack *SequentialList
		wantElem  ElemType
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{},
				Top:  -1,
			},
			want: Error,
		},
		{
			name: "02",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			want: OK,
			wantStack: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3},
				Top:  2,
			},
			wantElem: 4,
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
					if !reflect.DeepEqual(tt.s, tt.wantStack) {
						t.Errorf("got stack error, got %s, want %s", tt.s, tt.wantStack)
					}
					if e != tt.wantElem {
						t.Errorf("got elem error, got %v, want %v", e, tt.wantElem)
					}
				}
			},
		)
	}
}

// 测试是否为空栈
func TestSequentialList_StackEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialList
		want bool
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{},
				Top:  -1,
			},
			want: true,
		},
		{
			name: "02",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.StackEmpty()
				if got != tt.want {
					t.Errorf("StackEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestSequentialList_StackLength(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialList
		want int
	}{
		{
			name: "01",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{},
				Top:  -1,
			},
			want: 0,
		},
		{
			name: "02",
			s: &SequentialList{
				Data: [SequentialMaxSize]ElemType{1, 2, 3, 4},
				Top:  3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.StackLength()
				if got != tt.want {
					t.Errorf("StackLength() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
