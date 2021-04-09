package chapter4

import (
	"encoding/json"
	"fmt"
)

const (
	// 顺序存储线性表初始化最大分配量
	SequentialMaxSize = 20
	// 空栈顶
	EmptyTop = -1
)

// 顺序存储结构
type SequentialList struct {
	Data [SequentialMaxSize]ElemType `json:"data" remark:"数组，存储数据元素"`
	Top  int                         `json:"top" remark:"栈顶"`
}

// 字符串
func (s *SequentialList) String() string {
	b, _ := json.Marshal(s)
	return fmt.Sprintf("%s", b)
}

// 初始化
func (s *SequentialList) InitStack() Status {
	s.Data = [SequentialMaxSize]ElemType{}
	s.Top = EmptyTop
	return OK
}

// 销毁
func (s *SequentialList) DestroyStack() Status {
	return s.InitStack()
}

// 清空
func (s *SequentialList) ClearStack() Status {
	return s.InitStack()
}

// 是否为空
func (s *SequentialList) StackEmpty() bool {
	return s.Top == EmptyTop
}

// 获取栈顶
func (s *SequentialList) GetTop(e *ElemType) Status {
	if s.StackEmpty() {
		return Error
	}
	*e = s.Data[s.Top]
	return OK
}

// 入栈
func (s *SequentialList) Push(e ElemType) Status {
	if s.Top >= SequentialMaxSize-1 {
		return Error
	}
	s.Top++
	s.Data[s.Top] = e
	return OK
}

// 出栈
func (s *SequentialList) Pop(e *ElemType) Status {
	if s.StackEmpty() {
		return Error
	}
	*e = s.Data[s.Top]
	s.Data[s.Top] = 0
	s.Top--
	return OK
}

// 栈的长度
func (s *SequentialList) StackLength() int {
	return s.Top + 1
}
