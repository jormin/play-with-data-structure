package queue

const (
	// 操作成功
	OK = 1
	// 操作失败
	Error = 0
)

// 操作状态，1为成功，0为失败
type Status int

// 线性表元素类型，当前设定为int
type ElemType int

type Queue interface {
	// 初始化
	InitQueue() Status
	// 销毁
	DestroyQueue() Status
	// 清空
	ClearQueue() Status
	// 是否为空
	QueueEmpty() bool
	// 获取头元素
	GetHead(e *ElemType) Status
	// 入队
	EnQueue(e ElemType) Status
	// 出队
	DeQueue(e *ElemType) Status
	// 队列长度
	QueueLength() int
}
