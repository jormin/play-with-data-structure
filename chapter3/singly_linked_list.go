package chapter3

import (
	"fmt"
	"strings"
)

// 结点
type Node struct {
	Value interface{} `json:"data" remark:"存储数据"`
	Next  *Node
}

// 链式存储线性表
type SinglyLinkedList struct {
	Root   *Node `json:"root" remark:"根结点"`
	Length int   `json:"length" remark:"线性表长度"`
}

// 初始化操作，建立一个空的线性表L
func (l *SinglyLinkedList) InitList() Status {
	// 生成头结点，头结点的数据域存储的是线性表的长度
	l.Length = 0
	l.Root = nil
	return OK
}

// 线性表是否为空，若为空返回true，否则返回false
func (l *SinglyLinkedList) ListEmpty() bool {
	fmt.Println(l)
	return l.Root == nil
}

// 将线性表清空
func (l *SinglyLinkedList) ClearList() Status {
	l.InitList()
	return OK
}

// 将线性表L中的第i个位置元素值返回给e
func (l *SinglyLinkedList) GetElem(i int, e *ElemType) Status {
	if l.Length == 0 || i < 1 || i > l.Length {
		return Error
	}
	n := l.Root
	num := 1
	for n != nil && num < i {
		n = n.Next
		num++
	}
	if n == nil {
		return Error
	}
	*e = n.Value.(ElemType)
	return OK
}

// 在线性表L中查找与给定值e相等的元素，查找成功返回该元素的序号表示成功，否则返回0表示失败
func (l *SinglyLinkedList) LocateElem(e ElemType) int {
	n := l.Root
	num := 1
	for n != nil {
		if n.Value.(ElemType) == e {
			return num
		}
		num++
		n = n.Next
	}
	return 0
}

// 在线性表L的第i个位置插入新元素e
func (l *SinglyLinkedList) ListInsert(i int, e ElemType) Status {
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	pn, res := l.GetPrevNodeByNum(i)
	if res == Error {
		return Error
	}
	cn, res := l.GetNodeByNum(i)
	if res == Error {
		return Error
	}
	n := &Node{
		Value: e,
		Next:  cn,
	}
	l.insert(n, pn)
	l.Length++
	return OK
}

// 删除线性表L的第i个位置的元素，并用e返回其值
func (l *SinglyLinkedList) ListDelete(i int, e *ElemType) Status {
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	pn, res := l.GetPrevNodeByNum(i)
	if res == Error {
		return Error
	}
	cn, res := l.GetNodeByNum(i)
	if res == Error {
		return Error
	}
	l.remove(cn, pn)
	*e = cn.Value.(ElemType)
	l.Length--
	return OK
}

// 插入新结点
func (l *SinglyLinkedList) insert(n, pre *Node) {
	if pre == nil {
		n.Next = l.Root
		l.Root = n
	} else {
		n.Next = pre.Next
		pre.Next = n
	}
}

// 删除结点
func (l *SinglyLinkedList) remove(n, pre *Node) {
	if pre == nil {
		l.Root = n.Next
	} else {
		pre.Next = n.Next
	}
}

// 返回线性表L的元素个数
func (l *SinglyLinkedList) ListLength() int {
	return l.Length
}

// 打印链表
func (l *SinglyLinkedList) String() string {
	all := []string{"head"}
	n := l.Root
	for n != nil {
		all = append(all, fmt.Sprintf("%v", n.Value))
		n = n.Next
	}
	return fmt.Sprintf("length: %d, link: %v", l.Length, strings.Join(all, " -> "))
}

// 根据序号获取前结点
func (l *SinglyLinkedList) GetPrevNodeByNum(i int) (*Node, Status) {
	if i < 1 || i > l.Length {
		return nil, Error
	}
	if i == 1 {
		return nil, OK
	}
	n := l.Root
	num := 1
	for num < i-1 {
		n = n.Next
		num++
	}
	return n, OK
}

// 根据序号获取结点
func (l *SinglyLinkedList) GetNodeByNum(i int) (*Node, Status) {
	if i < 1 || i > l.Length {
		return nil, Error
	}
	n := l.Root
	if i == 1 {
		return n, OK
	}
	num := 1
	for num < i {
		n = n.Next
		num++
	}
	return n, OK
}
