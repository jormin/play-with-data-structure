package queue

import (
	"encoding/json"
	"fmt"
)

const (
	// 数组最大长度
	MaxSize = 20
)

// 双向循环队列
type RecycleList struct {
	Front int               `json:"front" remark:"队头"`
	Rear  int               `json:"rear" remark:"队尾"`
	Data  [MaxSize]ElemType `json:"data" remark:"数据"`
}

// 打印
func (r *RecycleList) String() string {
	b, _ := json.Marshal(r)
	return fmt.Sprintf("%s", b)
}

// 初始化
func (r *RecycleList) InitQueue() Status {
	r.Front = 0
	r.Rear = 0
	r.Data = [MaxSize]ElemType{}
	return OK
}

// 销毁
func (r *RecycleList) DestroyQueue() Status {
	return r.InitQueue()
}

// 清空
func (r *RecycleList) ClearQueue() Status {
	return r.InitQueue()
}

// 是否为空
func (r *RecycleList) QueueEmpty() bool {
	return r.Front == r.Rear
}

// 获取头元素
func (r *RecycleList) GetHead(e *ElemType) Status {
	if r.QueueEmpty() {
		return Error
	}
	*e = r.Data[r.Front]
	return OK
}

// 入队
func (r *RecycleList) EnQueue(e ElemType) Status {
	// 队列已满
	if (r.Rear+1)%MaxSize == r.Front {
		return Error
	}
	r.Data[r.Rear] = e
	r.Rear = (r.Rear + 1) % MaxSize
	return OK
}

// 出队
func (r *RecycleList) DeQueue(e *ElemType) Status {
	if r.QueueEmpty() {
		return Error
	}
	*e = r.Data[r.Front]
	r.Data[r.Front] = 0
	r.Front = (r.Front + 1) % MaxSize
	return OK
}

// 队列长度
func (r *RecycleList) QueueLength() int {
	return (r.Rear - r.Front + MaxSize) % MaxSize
}

// 绝对值
func abs(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}
