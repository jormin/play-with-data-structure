package queue

import (
	"reflect"
	"testing"
)

// 测试初始化
func TestRecycleList_InitQueue(t *testing.T) {
	tests := []struct {
		name string
		r    *RecycleList
		want Status
	}{
		{
			name: "01",
			r:    new(RecycleList),
			want: OK,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.r.InitQueue()
				if got != tt.want {
					t.Errorf("InitQueue() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试清空
func TestRecycleList_ClearQueue(t *testing.T) {
	tests := []struct {
		name      string
		r         *RecycleList
		want      Status
		wantQueue *RecycleList
	}{
		{
			name:      "01",
			r:         makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want:      OK,
			wantQueue: makeRecycleList(0, 0, []ElemType{}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.r.ClearQueue()
				if got != tt.want {
					t.Errorf("ClearQueue() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.r, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.r, tt.wantQueue)
				}
			},
		)
	}
}

// 测试销毁
func TestRecycleList_DestroyQueue(t *testing.T) {
	tests := []struct {
		name      string
		r         *RecycleList
		want      Status
		wantQueue *RecycleList
	}{
		{
			name:      "01",
			r:         makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want:      OK,
			wantQueue: makeRecycleList(0, 0, []ElemType{}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.r.DestroyQueue()
				if got != tt.want {
					t.Errorf("DestroyQueue() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.r, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.r, tt.wantQueue)
				}
			},
		)
	}
}

// 测试获取头元素
func TestRecycleList_GetHead(t *testing.T) {
	tests := []struct {
		name     string
		r        *RecycleList
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			r:    makeRecycleList(0, 0, []ElemType{}),
			want: Error,
		},
		{
			name:     "02",
			r:        makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want:     OK,
			wantElem: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.r.GetHead(&e)
				if got != tt.want {
					t.Errorf("GetHead() = %v, want %v", got, tt.want)
				}
				if got == OK && e != tt.wantElem {
					t.Errorf("want elem error, got %v, want %v", e, tt.wantElem)
				}
			},
		)
	}
}

// 测试入队
func TestRecycleList_EnQueue(t *testing.T) {
	tests := []struct {
		name      string
		r         *RecycleList
		e         ElemType
		want      Status
		wantQueue *RecycleList
	}{
		{
			name:      "01",
			r:         makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			e:         8,
			want:      OK,
			wantQueue: makeRecycleList(17, 5, []ElemType{1, 2, 3, 4, 5, 6, 7, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.r.EnQueue(tt.e)
				if got != tt.want {
					t.Errorf("EnQueue() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.r, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.r, tt.wantQueue)
				}
			},
		)
	}
}

// 测试出队
func TestRecycleList_DeQueue(t *testing.T) {
	tests := []struct {
		name      string
		r         *RecycleList
		want      Status
		wantQueue *RecycleList
		wantElem  ElemType
	}{
		{
			name: "01",
			r:    makeRecycleList(0, 0, []ElemType{}),
			want: Error,
		},
		{
			name:      "02",
			r:         makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want:      OK,
			wantQueue: makeRecycleList(18, 4, []ElemType{2, 3, 4, 5, 6, 7}),
			wantElem:  1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.r.DeQueue(&e)
				if got != tt.want {
					t.Errorf("DeQueue() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if !reflect.DeepEqual(tt.r, tt.wantQueue) {
						t.Errorf("want queue error, got %v, want %v", tt.r, tt.wantQueue)
					}
					if e != tt.wantElem {
						t.Errorf("want elem error, got %v, want %v", tt.r, tt.wantQueue)
					}
				}
			},
		)
	}
}

// 测试判断队列是否为空
func TestRecycleList_QueueEmpty(t *testing.T) {
	tests := []struct {
		name string
		r    *RecycleList
		want bool
	}{
		{
			name: "01",
			r:    makeRecycleList(0, 0, []ElemType{}),
			want: true,
		},
		{
			name: "02",
			r:    makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				res := tt.r.QueueEmpty()
				if res != tt.want {
					t.Errorf("QueueEmpty() = %v, want %v", res, tt.want)
				}
			},
		)
	}
}

// 测试获取队列长度
func TestRecycleList_QueueLength(t *testing.T) {
	tests := []struct {
		name string
		r    *RecycleList
		want int
	}{
		{
			name: "01",
			r:    makeRecycleList(0, 0, []ElemType{}),
			want: 0,
		},
		{
			name: "02",
			r:    makeRecycleList(0, 4, []ElemType{1, 2, 3, 4}),
			want: 4,
		},
		{
			name: "03",
			r:    makeRecycleList(17, 4, []ElemType{1, 2, 3, 4, 5, 6, 7}),
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				res := tt.r.QueueLength()
				if res != tt.want {
					t.Errorf("QueueEmpty() = %v, want %v", res, tt.want)
				}
			},
		)
	}
}

// 生成双向循环队列
func makeRecycleList(front int, rear int, es []ElemType) *RecycleList {
	r := new(RecycleList)
	if len(es) == 0 {
		r.InitQueue()
		return r
	}
	r.Front = front
	r.Rear = rear
	i := r.Front
	for _, v := range es {
		r.Data[i] = v
		i = (i + 1) % MaxSize
	}
	return r
}
