package chapter3

const (
	// 顺序存储线性表初始化最大分配量
	SequentialMaxSize = 20
)

// 顺序存储线性表
type SequentialList struct {
	Data   [SequentialMaxSize]ElemType `json:"data" remark:"数组，存储数据元素"`
	Length int                         `json:"length" remark:"线性表当前长度"`
}

// 初始化操作，建立一个空的线性表L
func (l *SequentialList) InitList() Status {
	l.Data = [SequentialMaxSize]ElemType{}
	l.Length = 0
	return OK
}

// 线性表是否为空，若为空返回true，否则返回false
func (l *SequentialList) ListEmpty() bool {
	return l.Length == 0
}

// 将线性表清空
func (l *SequentialList) ClearList() Status {
	l.InitList()
	return OK
}

// 将线性表L中的第i个位置元素值返回给e
func (l *SequentialList) GetElem(i int, e *ElemType) Status {
	if l.Length == 0 || i < 1 || i > l.Length {
		return Error
	}
	*e = l.Data[i-1]
	return OK
}

// 在线性表L中查找与给定值e相等的元素，查找成功返回该元素的序号表示成功，否则返回0表示失败
func (l *SequentialList) LocateElem(e ElemType) int {
	if l.Length == 0 {
		return 0
	}
	for i, v := range l.Data {
		if v == e {
			return i + 1
		}
	}
	return 0
}

// 在线性表L的第i个位置插入新元素e
func (l *SequentialList) ListInsert(i int, e ElemType) Status {
	// 如果当前线性表长度已达最大值则返回失败
	if l.Length == SequentialMaxSize {
		return Error
	}
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	// 当要插入的元素的位置不在表尾，则需要将当前该位置及之后的元素都向后挪一位
	for j := l.Length - 1; j >= i-1; j-- {
		l.Data[j+1] = l.Data[j]
	}
	l.Data[i-1] = e
	l.Length++
	return OK
}

// 删除线性表L的第i个位置的元素，并用e返回其值
func (l *SequentialList) ListDelete(i int, e *ElemType) Status {
	// 如果当前线性表长度为0则返回失败
	if l.Length == 0 {
		return Error
	}
	// 当要插入的位置比第一位1或者比最后一位大则返回失败
	if i < 1 || i > l.Length+1 {
		return Error
	}
	*e = l.Data[i-1]
	// 将当前位置及之后的元素向前挪
	for j := i; j < l.Length; j++ {
		l.Data[j-1] = l.Data[j]
	}
	l.Data[l.Length-1] = 0
	l.Length--
	return OK
}

// 返回线性表L的元素个数
func (l *SequentialList) ListLength() int {
	return l.Length
}
