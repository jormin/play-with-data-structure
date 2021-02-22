package chapter3

const (
	// 操作成功
	OK = 1
	// 操作失败
	Error = 0
	// 线性表初始化最大分配量
	MaxSize = 20
)

// 操作状态，1为成功，0为失败
type Status int

// 线性表元素类型，当前设定为int
type ElemType int

// 线性表
type List interface {
	// 初始化操作，建立一个空的线性表L
	InitList() Status
	// 线性表是否为空，若为空返回true，否则返回false
	ListEmpty() bool
	// 将线性表清空
	ClearList() Status
	// 将线性表L中的第i个位置元素值返回给e
	GetElem(i int, e *ElemType) Status
	// 在线性表L中查找与给定值e相等的元素，查找成功返回该元素的序号表示成功，否则返回0表示失败
	LocateElem(e ElemType) int
	// 在线性表L的第i个位置插入新元素e
	ListInsert(i int, e ElemType) Status
	// 删除线性表L的第i个位置的元素，并用e返回其值
	ListDelete(i int, e *ElemType) Status
	// 返回线性表L的元素个数
	ListLength() int
}
