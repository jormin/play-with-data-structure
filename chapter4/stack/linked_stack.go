package stack

import (
	"fmt"
	"strings"
)

// 结点
type Node struct {
	Value ElemType `json:"data" remark:"存储数据"`
	Next  *Node
}

// 链式存储栈
type LinkedStack struct {
	Top    *Node `json:"root" remark:"根结点"`
	Length int   `json:"length" remark:"线性表长度"`
}

// 字符串
func (s *LinkedStack) String() string {
	var all []string
	n := s.Top
	for n != nil {
		all = append(all, fmt.Sprintf("%v", n.Value))
		n = n.Next
	}
	return fmt.Sprintf("length: %d, link: %v", s.Length, strings.Join(all, " -> "))
}

// 初始化
func (s *LinkedStack) InitStack() Status {
	s.Length = 0
	s.Top = nil
	return OK
}

// 销毁
func (s *LinkedStack) DestroyStack() Status {
	return s.InitStack()
}

// 清空
func (s *LinkedStack) ClearStack() Status {
	return s.InitStack()
}

// 是否为空
func (s *LinkedStack) StackEmpty() bool {
	return s.Top == nil
}

// 获取栈顶
func (s *LinkedStack) GetTop(e *ElemType) Status {
	if s.StackEmpty() {
		return Error
	}
	*e = s.Top.Value
	return OK
}

// 入栈
func (s *LinkedStack) Push(e ElemType) Status {
	nt := Node{
		Value: e,
		Next:  s.Top,
	}
	s.Top = &nt
	s.Length++
	return OK
}

// 出栈
func (s *LinkedStack) Pop(e *ElemType) Status {
	if s.StackEmpty() {
		return Error
	}
	*e = s.Top.Value
	s.Top = s.Top.Next
	s.Length--
	return OK
}

// 栈的长度
func (s *LinkedStack) StackLength() int {
	return s.Length
}
