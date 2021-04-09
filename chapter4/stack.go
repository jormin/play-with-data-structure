package chapter4

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

// 栈
type Stack interface {
	// 初始化
	InitStack() Status
	// 销毁
	DestroyStack() Status
	// 清空
	ClearStack() Status
	// 是否为空
	StackEmpty() bool
	// 获取栈顶
	GetTop(e *ElemType) Status
	// 入栈
	Push(e ElemType) Status
	// 出栈
	Pop(e *ElemType) Status
	// 栈的长度
	StackLength() int
}
