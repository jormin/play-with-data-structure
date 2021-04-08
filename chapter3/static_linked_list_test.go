package chapter3

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试初始化线性表
func TestStaticLinkedInitList(t *testing.T) {
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
				l := &StaticLinkedList{}
				if got := l.InitList(); got != tt.want {
					t.Errorf("InitList() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试线性表是否为空函数
func TestStaticLinkedListEmpty(t *testing.T) {
	tests := []struct {
		name string
		l    *StaticLinkedList
		want bool
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			want: true,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
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
func TestStaticLinkedClearList(t *testing.T) {
	tests := []struct {
		name string
		l    *StaticLinkedList
		want Status
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			want: OK,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
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
func TestStaticLinkedGetElem(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		l        *StaticLinkedList
		args     args
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			args: args{
				i: 0,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
			args: args{
				i: 0,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "03",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
			args: args{
				i: 11,
			},
			want:     Error,
			wantElem: 0,
		},
		{
			name: "04",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
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
func TestStaticLinkedLocateElem(t *testing.T) {
	type args struct {
		e ElemType
	}
	tests := []struct {
		name string
		l    *StaticLinkedList
		args args
		want int
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			args: args{
				e: 1,
			},
			want: 0,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
			args: args{
				e: 11,
			},
			want: 0,
		},
		{
			name: "03",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
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
func TestStaticLinkedListInsert(t *testing.T) {
	type args struct {
		i int
		e ElemType
	}
	tests := []struct {
		name     string
		l        *StaticLinkedList
		args     args
		want     Status
		wantList *StaticLinkedList
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			args: args{
				i: 20,
				e: 1,
			},
			want: Error,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				4,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 0},
				},
			),
			args: args{
				i: 0,
				e: 1,
			},
			want: Error,
		},
		{
			name: "03",
			l: makeStaticLinkedList(
				4,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 0},
				},
			),
			args: args{
				i: 5,
				e: 1,
			},
			want: Error,
		},
		{
			name: "04",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
			args: args{
				i: 10,
				e: 1,
			},
			want: OK,
			wantList: makeStaticLinkedList(
				12,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 11},
					{10, 10, 0},
					{11, 1, 10},
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.l.ListInsert(tt.args.i, tt.args.e)
				fmt.Println(tt.l)
				fmt.Println(tt.wantList)
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
func TestStaticLinkedListDelete(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		l        *StaticLinkedList
		args     args
		want     Status
		wantElem ElemType
		wantList *StaticLinkedList
	}{

		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			args: args{
				i: 1,
			},
			want: Error,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				4,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 0},
				},
			),
			args: args{
				i: 0,
			},
			want: Error,
		},
		{
			name: "03",
			l: makeStaticLinkedList(
				4,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 0},
				},
			),
			args: args{
				i: 5,
			},
			want: Error,
		},
		{
			name: "04",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
			args: args{
				i: 9,
			},
			want:     OK,
			wantElem: 9,
			wantList: makeStaticLinkedList(
				9,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 10},
					{10, 10, 0},
				},
			),
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
					fmt.Println(tt.wantList)
					if !reflect.DeepEqual(tt.l, tt.wantList) {
						t.Errorf("list = %v, want list %v", tt.l, tt.wantList)
					}
				}
			},
		)
	}
}

// 测试获取线性表长度
func TestStaticLinkedListLength(t *testing.T) {
	tests := []struct {
		name string
		l    *StaticLinkedList
		want int
	}{
		{
			name: "01",
			l:    makeStaticLinkedList(1, [][]int{}),
			want: 0,
		},
		{
			name: "02",
			l: makeStaticLinkedList(
				11,
				[][]int{
					{1, 1, 2},
					{2, 2, 3},
					{3, 3, 4},
					{4, 4, 5},
					{5, 5, 6},
					{6, 6, 7},
					{7, 7, 8},
					{8, 8, 9},
					{9, 9, 10},
					{10, 10, 0},
				},
			),
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

// 生成静态链表线性表，元素格式为长度为3的数组【下标、值、下结点下标】
func makeStaticLinkedList(fc int, list [][]int) *StaticLinkedList {
	l := &StaticLinkedList{}
	if len(list) == 0 {
		l.InitList()
		return l
	}
	pos := map[int]int{}
	// 最后一个元素的Cursor存放第一个有数值的元素的下标
	l.Data[StaticLinkedMaxSize-1].Cursor = 1
	// 长度为数组长度
	l.Length = len(list)
	// 循环插入数据
	for _, v := range list {
		l.Data[v[0]] = StaticLinkedData{
			Value:  ElemType(v[1]),
			Cursor: v[2],
		}
		pos[v[0]] = 1
	}
	// 循环处理备用结点
	bos := []int{}
	for i := 1; i <= len(l.Data); i++ {
		if _, ok := pos[i]; !ok {
			bos = append(bos, i)
		}
	}
	l.Data[0].Cursor = bos[0]
	for i := 0; i < len(bos)-1; i++ {
		l.Data[bos[i]] = StaticLinkedData{
			Value:  0,
			Cursor: bos[i+1],
		}
	}
	// 尾元素的游标指向第一个元素
	l.Data[StaticLinkedMaxSize-1].Cursor = 1
	return l
}
