package chapter3

import (
	"encoding/json"
	"fmt"
)

const (
	// 静态链表线性表初始化最大分配量
	StaticLinkedMaxSize = 20
)

// 静态链表数据域
type StaticLinkedData struct {
	Value  ElemType `json:"value" remark:"存储数据"`
	Cursor int      `json:"cursor" remark:"游标"`
}

// 实现String函数
func (s *StaticLinkedData) String() string {
	b, _ := json.Marshal(s)
	return fmt.Sprintf("%s", b)
}

// 静态链表线性表
type StaticLinkedList struct {
	Data   [StaticLinkedMaxSize]StaticLinkedData `json:"data" remark:"数组，存储静态链表数据域"`
	Length int                                   `json:"length" remark:"线性表当前长度"`
}

// 实现String函数
func (l *StaticLinkedList) String() string {
	b, _ := json.Marshal(l)
	return fmt.Sprintf("%s", b)
}

// 初始化操作，建立一个空的线性表L
func (l *StaticLinkedList) InitList() Status {
	l.Data = [StaticLinkedMaxSize]StaticLinkedData{}
	// 除了尾元素，每一个元素的游标都指向下一个元素的下标
	for i := 0; i < StaticLinkedMaxSize-1; i++ {
		l.Data[i].Cursor = i + 1
	}
	// 尾元素的游标指向第一个元素
	l.Data[StaticLinkedMaxSize-1].Cursor = 0
	// 线性表长度为0
	l.Length = 0
	return OK
}

// 线性表是否为空，若为空返回true，否则返回false
func (l *StaticLinkedList) ListEmpty() bool {
	return l.Length == 0
}

// 将线性表清空
func (l *StaticLinkedList) ClearList() Status {
	l.InitList()
	return OK
}

// 将线性表L中的第i个位置元素值返回给e
func (l *StaticLinkedList) GetElem(i int, e *ElemType) Status {
	if l.Length == 0 || i < 1 || i > l.Length {
		return Error
	}
	// 先查询第一个元素的下标
	ni := l.Data[StaticLinkedMaxSize-1].Cursor
	num := 1
	for num < i {
		ni = l.Data[ni].Cursor
		num++
	}
	*e = l.Data[ni].Value
	return OK
}

// 在线性表L中查找与给定值e相等的元素，查找成功返回该元素的序号表示成功，否则返回0表示失败
func (l *StaticLinkedList) LocateElem(e ElemType) int {
	if l.Length == 0 {
		return 0
	}
	// 先查询第一个元素的下标
	ni := l.Data[StaticLinkedMaxSize-1].Cursor
	num := 1
	for num <= l.Length {
		if l.Data[ni].Value == e {
			return num
		}
		ni = l.Data[ni].Cursor
		num++
	}
	return 0
}

// 在线性表L的第i个位置插入新元素e
func (l *StaticLinkedList) ListInsert(i int, e ElemType) Status {
	// 如果当前线性表长度已达最大值则返回失败
	if l.Length == StaticLinkedMaxSize {
		return Error
	}
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	// 先查询第一个元素的下标
	ni := l.Data[StaticLinkedMaxSize-1].Cursor
	// 第一个元素
	n := l.Data[ni]
	num := 1
	// 找到插入位置的前一个元素
	for n.Cursor != 0 && num < i-1 {
		ni = l.Data[ni].Cursor
		num++
	}
	// 当前该位置元素的游标
	ci := l.Data[ni].Cursor
	// 获取第一个备用结点
	bi := l.Data[0].Cursor
	// 更改首结点的游标为下一个备用结点
	l.Data[0].Cursor = l.Data[bi].Cursor
	// 更改当前选中的备用结点的值
	l.Data[bi].Value = e
	// 更改该结点的游标
	l.Data[bi].Cursor = ci
	// 更改前结点的游标
	l.Data[ni].Cursor = bi
	l.Length++
	return OK
}

// 删除线性表L的第i个位置的元素，并用e返回其值
func (l *StaticLinkedList) ListDelete(i int, e *ElemType) Status {
	// 如果当前线性表长度为0则返回失败
	if l.Length == 0 {
		return Error
	}
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	// 先查询第一个元素的下标
	ni := l.Data[StaticLinkedMaxSize-1].Cursor
	num := 1
	n := l.Data[ni]
	// 找到插入位置的前一个元素
	for n.Cursor != 0 && num < l.Length-2 {
		ni = l.Data[ni].Cursor
		num++
	}
	// 前一位置元素的游标
	ci := l.Data[ni].Cursor
	// 更改前结点的游标
	l.Data[ni].Cursor = l.Data[ci].Cursor
	// 更改当前结点
	*e = l.Data[ci].Value
	l.Data[ci] = StaticLinkedData{}
	// 更改首个备用结点的下标
	if l.Data[0].Cursor > i {
		l.Data[ci].Cursor = l.Data[0].Cursor
		l.Data[0].Cursor = i
	}
	l.Length--
	return OK
}

// 返回线性表L的元素个数
func (l *StaticLinkedList) ListLength() int {
	return l.Length
}
