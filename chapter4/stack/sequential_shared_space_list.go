package stack

import (
	"encoding/json"
	"fmt"
)

const (
	// 左侧空栈顶
	LeftEmptyTop = -1
	// 右侧栈顶
	RightEmptyTop = SequentialMaxSize
	// 左侧
	Left = 0
	// 右侧
	Right = 1
)

// 顺序存储结构 - 共享空间
type SequentialSharedSpaceList struct {
	Data     [SequentialMaxSize]ElemType `json:"data" remark:"数组，存储数据元素"`
	LeftTop  int                         `json:"left_top" remark:"左侧栈顶"`
	RightTop int                         `json:"right_top" remark:"右侧栈顶"`
}

// 字符串
func (s *SequentialSharedSpaceList) String() string {
	b, _ := json.Marshal(s)
	return fmt.Sprintf("%s", b)
}

// 初始化
func (s *SequentialSharedSpaceList) InitStack() Status {
	s.Data = [SequentialMaxSize]ElemType{}
	s.LeftTop = LeftEmptyTop
	s.RightTop = RightEmptyTop
	return OK
}

// 销毁
func (s *SequentialSharedSpaceList) DestroyStack() Status {
	return s.InitStack()
}

// 清空
func (s *SequentialSharedSpaceList) ClearStack() Status {
	return s.InitStack()
}

// 是否为空
func (s *SequentialSharedSpaceList) StackEmpty() bool {
	return s.LeftTop == LeftEmptyTop && s.RightTop == RightEmptyTop
}

// 获取栈顶
func (s *SequentialSharedSpaceList) GetTop(e *ElemType, pos int) Status {
	if pos == Left {
		// 左侧栈
		if s.LeftTop == LeftEmptyTop {
			return Error
		}
		*e = s.Data[s.LeftTop]
	} else {
		// 右侧栈
		if s.RightTop == RightEmptyTop {
			return Error
		}
		*e = s.Data[s.RightTop]
	}
	return OK
}

// 入栈
func (s *SequentialSharedSpaceList) Push(e ElemType, pos int) Status {
	// 判断两个栈是否都满了
	if s.LeftTop+1 == s.RightTop {
		return Error
	}
	if pos == Left {
		// 左侧栈
		s.LeftTop++
		s.Data[s.LeftTop] = e
	} else {
		// 右侧栈
		s.RightTop--
		s.Data[s.RightTop] = e
	}
	return OK
}

// 出栈
func (s *SequentialSharedSpaceList) Pop(e *ElemType, pos int) Status {
	if pos == Left {
		// 左侧栈
		if s.LeftTop == LeftEmptyTop {
			return Error
		}
		*e = s.Data[s.LeftTop]
		s.Data[s.LeftTop] = 0
		s.LeftTop--
	} else {
		// 右侧栈
		if s.RightTop == RightEmptyTop {
			return Error
		}
		*e = s.Data[s.RightTop]
		s.Data[s.RightTop] = 0
		s.RightTop++
	}
	return OK
}

// 栈的长度
func (s *SequentialSharedSpaceList) StackLength() int {
	return s.LeftTop + 1 + SequentialMaxSize - s.RightTop
}
