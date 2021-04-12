package queue

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试初始化
func TestLinkedQueue_InitQueue(t *testing.T) {
	tests := []struct {
		name string
		q    *LinkedQueue
		want Status
	}{
		{
			name: "01",
			q:    new(LinkedQueue),
			want: OK,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.InitQueue()
				if got != tt.want {
					t.Errorf("InitQueue() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试清空
func TestLinkedQueue_ClearQueue(t *testing.T) {
	tests := []struct {
		name      string
		q         *LinkedQueue
		want      Status
		wantQueue *LinkedQueue
	}{
		{
			name:      "01",
			q:         makeLinkedQueue(1, 2, 3, 4),
			want:      OK,
			wantQueue: makeLinkedQueue(),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.ClearQueue()
				if got != tt.want {
					t.Errorf("ClearQueue() = %v, want %v", got, tt.want)
				}
				fmt.Println(tt.q, tt.wantQueue)
				if got == OK && !reflect.DeepEqual(tt.q, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.q, tt.wantQueue)
				}
			},
		)
	}
}

// 测试销毁
func TestLinkedQueue_DestroyQueue(t *testing.T) {
	tests := []struct {
		name      string
		q         *LinkedQueue
		want      Status
		wantQueue *LinkedQueue
	}{
		{
			name:      "01",
			q:         makeLinkedQueue(1, 2, 3, 4),
			want:      OK,
			wantQueue: makeLinkedQueue(),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.DestroyQueue()
				if got != tt.want {
					t.Errorf("DestroyQueue() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.q, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.q, tt.wantQueue)
				}
			},
		)
	}
}

// 测试获取头元素
func TestLinkedQueue_GetHead(t *testing.T) {
	tests := []struct {
		name     string
		q        *LinkedQueue
		want     Status
		wantElem ElemType
	}{
		{
			name: "01",
			q:    makeLinkedQueue(),
			want: Error,
		},
		{
			name:     "02",
			q:        makeLinkedQueue(1, 2, 3, 4),
			want:     OK,
			wantElem: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.q.GetHead(&e)
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
func TestLinkedQueue_EnQueue(t *testing.T) {
	tests := []struct {
		name      string
		q         *LinkedQueue
		e         ElemType
		want      Status
		wantQueue *LinkedQueue
	}{
		{
			name:      "01",
			q:         makeLinkedQueue(1, 2, 3, 4),
			e:         5,
			want:      OK,
			wantQueue: makeLinkedQueue(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.EnQueue(tt.e)
				if got != tt.want {
					t.Errorf("EnQueue() = %v, want %v", got, tt.want)
				}
				if got == OK && !reflect.DeepEqual(tt.q, tt.wantQueue) {
					t.Errorf("want queue error, got %v, want %v", tt.q, tt.wantQueue)
				}
			},
		)
	}
}

// 测试出队
func TestLinkedQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name      string
		q         *LinkedQueue
		e         ElemType
		want      Status
		wantQueue *LinkedQueue
		wantElem  ElemType
	}{
		{
			name: "01",
			q:    makeLinkedQueue(),
			want: Error,
		},
		{
			name:      "02",
			q:         makeLinkedQueue(1, 2, 3, 4),
			want:      OK,
			wantQueue: makeLinkedQueue(2, 3, 4),
			wantElem:  1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var e ElemType
				got := tt.q.DeQueue(&e)
				if got != tt.want {
					t.Errorf("DeQueue() = %v, want %v", got, tt.want)
				}
				if got == OK {
					if !reflect.DeepEqual(tt.q, tt.wantQueue) {
						t.Errorf("want queue error, got %v, want %v", tt.q, tt.wantQueue)
					}
					if e != tt.wantElem {
						t.Errorf("want elem error, got %v, want %v", e, tt.wantElem)
					}
				}
			},
		)
	}
}

// 测试判断是否为空
func TestLinkedQueue_QueueEmpty(t *testing.T) {
	tests := []struct {
		name string
		q    *LinkedQueue
		want bool
	}{
		{
			name: "01",
			q:    makeLinkedQueue(),
			want: true,
		},
		{
			name: "02",
			q:    makeLinkedQueue(1, 2, 3, 4),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.QueueEmpty()
				if got != tt.want {
					t.Errorf("QueueEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 测试获取长度
func TestLinkedQueue_QueueLength(t *testing.T) {
	tests := []struct {
		name string
		q    *LinkedQueue
		want int
	}{
		{
			name: "01",
			q:    makeLinkedQueue(),
			want: 0,
		},
		{
			name: "02",
			q:    makeLinkedQueue(1, 2, 3, 4),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.q.QueueLength()
				if got != tt.want {
					t.Errorf("QueueLength() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 生成链式队列
func makeLinkedQueue(es ...ElemType) *LinkedQueue {
	q := new(LinkedQueue)
	if len(es) == 0 {
		q.InitQueue()
		return q
	}
	n := &Node{
		Value: 0,
		Next:  nil,
	}
	for i, v := range es {
		n.Next = &Node{
			Value: v,
			Next:  nil,
		}
		n = n.Next
		if i == 0 {
			q.Front = n
		}
		if i == len(es)-1 {
			q.Rear = n
		}
	}
	q.Length = len(es)
	return q
}
