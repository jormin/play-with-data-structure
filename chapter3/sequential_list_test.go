package chapter3

import (
	"reflect"
	"testing"
)

// 测试初始化线性表
func TestSequentialInitList(t *testing.T) {
	tests := []struct {
		name string
		want Status
	}{
		{
			name: "normal",
			want: OK,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l := &SequentialList{}
				if got := l.InitList(); got != tt.want {
					t.Errorf("InitList() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试线性表是否为空函数
func TestSequentialListEmpty(t *testing.T) {
	tests := []struct {
		name string
		l    *SequentialList
		want bool
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			want: true,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.l.ListEmpty(); got != tt.want {
					t.Errorf("ListEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试将线性表清空
func TestSequentialClearList(t *testing.T) {
	tests := []struct {
		name string
		l    *SequentialList
		want Status
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			want: OK,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			want: OK,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.l.ClearList(); got != tt.want {
					t.Errorf("ClearList() = %v, want %v", got, tt.want)
				}
				if tt.l.Length != 0 {
					t.Errorf("tt.l.Length = %v, want %v", tt.l.Length, 0)
				}
			},
		)
	}
}

// 测试将线性表L中的第i个位置元素值返回给e
func TestSequentialGetElem(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		l        *SequentialList
		args     args
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			args: args{
				i: 0,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				i: 0,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "03",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				i: 11,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "04",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				i: 10,
			},
			want:     OK,
			wantElem: 10,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				if got := tt.l.GetElem(tt.args.i, &e); got != tt.want {
					t.Errorf("GetElem() = %v, want %v", got, tt.want)
				}
				if e != tt.wantElem {
					t.Errorf("elem = %v, want %v", e, tt.wantElem)
				}
			},
		)
	}
}

// 测试在线性表L中查找与给定值e相等的元素
func TestSequentialLocateElem(t *testing.T) {
	type args struct {
		e ElemType
	}
	tests := []struct {
		name string
		l    *SequentialList
		args args
		want int
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			args: args{
				e: 1,
			},
			want: 0,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				e: 11,
			},
			want: 0,
		},
		{
			name: "03",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				e: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.l.LocateElem(tt.args.e); got != tt.want {
					t.Errorf("LocateElem() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试添加元素
func TestSequentialListInsert(t *testing.T) {
	type args struct {
		i int
		e ElemType
	}
	tests := []struct {
		name     string
		l        *SequentialList
		args     args
		want     Status
		wantList *SequentialList
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 20,
			},
			args: args{
				i: 20,
				e: 1,
			},
			want: Error,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{0, 1, 2},
				Length: 3,
			},
			args: args{
				i: 0,
				e: 1,
			},
			want: Error,
		},
		{
			name: "03",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{0, 1, 2},
				Length: 3,
			},
			args: args{
				i: 5,
				e: 1,
			},
			want: Error,
		},
		{
			name: "04",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				i: 10,
				e: 1,
			},
			want: OK,
			wantList: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 10},
				Length: 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.l.ListInsert(tt.args.i, tt.args.e)
				if got != tt.want {
					t.Errorf("ListInsert() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.l, tt.wantList) {
					t.Errorf("list = %v, want list %v", tt.l, tt.wantList)
				}
			},
		)
	}
}

// 测试删除元素
func TestSequentialListDelete(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		l        *SequentialList
		args     args
		want     Status
		wantElem ElemType
		wantList *SequentialList
	}{

		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			args: args{
				i: 1,
			},
			want: Error,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{0, 1, 2},
				Length: 3,
			},
			args: args{
				i: 0,
			},
			want: Error,
		},
		{
			name: "03",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{0, 1, 2},
				Length: 3,
			},
			args: args{
				i: 5,
			},
			want: Error,
		},
		{
			name: "04",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			args: args{
				i: 9,
			},
			want:     OK,
			wantElem: 9,
			wantList: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 10},
				Length: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.l.ListDelete(tt.args.i, &e)
				if got != tt.want {
					t.Errorf("ListDelete() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if e != tt.wantElem {
						t.Errorf("elem = %v, want elem %v", e, tt.wantElem)
					}
					if !reflect.DeepEqual(tt.l, tt.wantList) {
						t.Errorf("list = %v, want list %v", tt.l, tt.wantList)
					}
				}
			},
		)
	}
}

// 测试获取线性表长度
func TestSequentialListLength(t *testing.T) {
	tests := []struct {
		name string
		l    *SequentialList
		want int
	}{
		{
			name: "01",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{},
				Length: 0,
			},
			want: 0,
		},
		{
			name: "02",
			l: &SequentialList{
				Data:   [SequentialMaxSize]ElemType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Length: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.l.ListLength(); got != tt.want {
					t.Errorf("ListLength() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
