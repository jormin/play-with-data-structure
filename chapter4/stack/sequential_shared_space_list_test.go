package stack

import (
	"reflect"
	"testing"
)

// 测试初始化
func TestSequentialSharedSpaceList_InitStack(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialSharedSpaceList
		want Status
	}{
		{
			name: "01",
			s:    new(SequentialSharedSpaceList),
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
func TestSequentialSharedSpaceList_ClearStack(t *testing.T) {
	tests := []struct {
		name         string
		s            *SequentialSharedSpaceList
		want         Status
		wantLeftTop  int
		wantRightTop int
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			want:         OK,
			wantLeftTop:  LeftEmptyTop,
			wantRightTop: RightEmptyTop,
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
					if tt.s.LeftTop != tt.wantLeftTop {
						t.Errorf("want left top error, got %v, want %v", tt.s.LeftTop, tt.wantLeftTop)
					}
					if tt.s.RightTop != tt.wantRightTop {
						t.Errorf("want right top error, got %v, want %v", tt.s.RightTop, tt.wantRightTop)
					}
				}
			},
		)
	}
}

// 测试销毁
func TestSequentialSharedSpaceList_DestroyStack(t *testing.T) {
	tests := []struct {
		name         string
		s            *SequentialSharedSpaceList
		want         Status
		wantLeftTop  int
		wantRightTop int
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			want:         OK,
			wantLeftTop:  LeftEmptyTop,
			wantRightTop: RightEmptyTop,
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
					if tt.s.LeftTop != tt.wantLeftTop {
						t.Errorf("want left top error, got %v, want %v", tt.s.LeftTop, tt.wantLeftTop)
					}
					if tt.s.RightTop != tt.wantRightTop {
						t.Errorf("want right top error, got %v, want %v", tt.s.RightTop, tt.wantRightTop)
					}
				}
			},
		)
	}
}

// 测试获取栈顶
func TestSequentialSharedSpaceList_GetTop(t *testing.T) {
	tests := []struct {
		name     string
		s        *SequentialSharedSpaceList
		pos      int
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  EmptyTop,
				RightTop: RightEmptyTop,
			},
			pos:  Left,
			want: Error,
		},
		{
			name: "02",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  EmptyTop,
				RightTop: RightEmptyTop,
			},
			pos:  Right,
			want: Error,
		},
		{
			name: "03",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			pos:      Left,
			want:     OK,
			wantElem: 4,
		},
		{
			name: "04",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 15,
			},
			pos:      Right,
			want:     OK,
			wantElem: 5,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.s.GetTop(&e, tt.pos)
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
func TestSequentialSharedSpaceList_Push(t *testing.T) {
	tests := []struct {
		name      string
		s         *SequentialSharedSpaceList
		elem      ElemType
		pos       int
		want      Status
		wantStack *SequentialSharedSpaceList
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
				},
				LeftTop:  9,
				RightTop: 10,
			},
			want: Error,
		},
		{
			name: "02",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			elem: 5,
			pos:  Left,
			want: OK,
			wantStack: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  4,
				RightTop: 16,
			},
		},
		{
			name: "03",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			elem: 5,
			pos:  Right,
			want: OK,
			wantStack: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.s.Push(tt.elem, tt.pos)
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
func TestSequentialSharedSpaceList_Pop(t *testing.T) {
	tests := []struct {
		name      string
		s         *SequentialSharedSpaceList
		pos       int
		want      Status
		wantStack *SequentialSharedSpaceList
		wantElem  ElemType
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  LeftEmptyTop,
				RightTop: RightEmptyTop,
			},
			pos:  Left,
			want: Error,
		},
		{
			name: "02",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  LeftEmptyTop,
				RightTop: RightEmptyTop,
			},
			pos:  Right,
			want: Error,
		},
		{
			name: "03",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			pos:  Left,
			want: OK,
			wantStack: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  2,
				RightTop: 16,
			},
			wantElem: 4,
		},
		{
			name: "04",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			pos:  Right,
			want: OK,
			wantStack: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 17,
			},
			wantElem: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.s.Pop(&e, tt.pos)
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
func TestSequentialSharedSpaceList_StackEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialSharedSpaceList
		want bool
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  LeftEmptyTop,
				RightTop: RightEmptyTop,
			},
			want: true,
		},
		{
			name: "02",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  EmptyTop,
				RightTop: 16,
			},
			want: false,
		},
		{
			name: "03",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
				LeftTop:  3,
				RightTop: RightEmptyTop,
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

func TestSequentialSharedSpaceList_StackLength(t *testing.T) {
	tests := []struct {
		name string
		s    *SequentialSharedSpaceList
		want int
	}{
		{
			name: "01",
			s: &SequentialSharedSpaceList{
				Data:     [SequentialMaxSize]ElemType{},
				LeftTop:  LeftEmptyTop,
				RightTop: RightEmptyTop,
			},
			want: 0,
		},
		{
			name: "02",
			s: &SequentialSharedSpaceList{
				Data: [SequentialMaxSize]ElemType{
					1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 2, 1,
				},
				LeftTop:  3,
				RightTop: 16,
			},
			want: 8,
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
