package queue

import (
	"fmt"
	"strings"
)

// 结点
type Node struct {
	Value ElemType `json:"data" remark:"存储数据"`
	Next  *Node
}

// 链式队列
type LinkedQueue struct {
	Front  *Node `json:"front" remark:"头"`
	Rear   *Node `json:"rear" remark:"尾"`
	Length int   `json:"length" remark:"长度"`
}

// 字符串
func (q *LinkedQueue) String() string {
	var all []string
	n := q.Front
	i := 0
	for n != nil {
		all = append(all, fmt.Sprintf("%v", n.Value))
		n = n.Next
		i++
	}
	return fmt.Sprintf("length: %d, link: %v", i, strings.Join(all, " -> "))
}

// 初始化
func (q *LinkedQueue) InitQueue() Status {
	q.Front = nil
	q.Rear = nil
	q.Length = 0
	return OK
}

// 销毁
func (q *LinkedQueue) DestroyQueue() Status {
	return q.InitQueue()
}

// 清空
func (q *LinkedQueue) ClearQueue() Status {
	return q.InitQueue()
}

// 判断是否为空
func (q *LinkedQueue) QueueEmpty() bool {
	return q.Front == q.Rear
}

// 获取头结点
func (q *LinkedQueue) GetHead(e *ElemType) Status {
	if q.QueueEmpty() {
		return Error
	}
	*e = q.Front.Value
	return OK
}

// 入队
func (q *LinkedQueue) EnQueue(e ElemType) Status {
	q.Rear.Next = &Node{
		Value: e,
		Next:  nil,
	}
	q.Rear = q.Rear.Next
	q.Length++
	return OK
}

// 出队
func (q *LinkedQueue) DeQueue(e *ElemType) Status {
	if q.QueueEmpty() {
		return Error
	}
	*e = q.Front.Value
	q.Front = q.Front.Next
	if q.Length == 1 {
		q.Rear = q.Front
	}
	q.Length--
	return OK
}

// 获取长度
func (q *LinkedQueue) QueueLength() int {
	return q.Length
}
