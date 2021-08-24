package main

import (
	"errors"
	"fmt"
)

// 顺序表，全名顺序存储结构，是线性表的一种。

// 线性表用于存储逻辑关系为“一对一”的数据，顺序表自然也不例外。

// 顺序表对数据的物理存储结构也有要求。
// 顺序表存储数据时，会提前申请一整块足够大小的物理空间，然后将数据依次存储起来，存储时做到数据元素之间不留一丝缝隙。


// slice 是对底层数组的抽象和控制。它是一个结构体：
	//	type slice struct {
	//	    array unsafe.Pointer
	//	    len   int
	//	    cap   int
	//	}


type  SeqList struct {
	ptr *[]interface{}   	// 指向线性表空间指针
	len int      			// 线性表长度
	cap int					// 表容量
}

// 顺序表初始化
func NewSeqList(cap int) *SeqList {
	if cap <= 0 {
		return nil
	}
	ptr := make([]interface{}, cap)
	return &SeqList{
		ptr: &ptr,
		len: 0,
		cap: cap,
	}
}

// 判空
func (l *SeqList) Init() bool {
	if l.len == 0 {
		return true
	}
	return false
}


// 判空
func (l *SeqList) IsEmpty() bool {
	if l.len == 0 {
		return true
	}
	return false
}


// 判满
func (l *SeqList) IsFull() bool {
	if l.len == l.cap {
		return true
	}
	return false
}


// 返回长度
func (l *SeqList) Len() int {
	return l.len
}

// 返回容量
func (l *SeqList) Cap() int {
	return l.cap
}


func (l *SeqList) Append(elem interface{}) {
	if l.IsFull() {
		l.grow()
	}
	(*l.ptr)[l.len] = elem
	l.len++
}


func (l *SeqList) grow() {
	newPtr := make([]interface{}, l.cap * 2)
	copy(newPtr, *l.ptr)
	l.ptr = &newPtr
	l.cap = l.cap * 2
}


// 顺序表插入元素
	//向已有顺序表中插入数据元素，根据插入位置的不同，可分为以下 3 种情况：
		//- 插入到顺序表的表头；
		//- 在表的中间位置插入元素；
		//- 尾随顺序表中已有元素，作为顺序表中的最后一个元素；

	//虽然数据元素插入顺序表中的位置有所不同，但是都使用的是同一种方式去解决，即：通过遍历，找到数据元素要插入的位置，然后做如下两步工作：
		// 将要插入位置元素以及后续的元素整体向后移动一个位置；
		// 将元素放到腾出来的位置上；
// 插入元素,index为插入的位置，elem为插入值
func (l *SeqList) Insert (index int, elem interface{}) bool {
	if index < 0 || index > l.len || l.IsFull(){
		return false
	}
	// 先将index位置元素以及之后的元素后移一位
	for i := l.len - 1; i >= index; i-- {
		(*l.ptr)[i + 1] = (*l.ptr)[i]
	}
	(*l.ptr)[index] = elem
	l.len++
	return true
}

// 顺序表删除元素
	// 从顺序表中删除指定元素，实现起来非常简单，只需找到目标元素，并将其后续所有元素整体前移 1 个位置即可。
	// 后续元素整体前移一个位置，会直接将目标元素删除，可间接实现删除元素的目的。
func (l *SeqList) Del (index int) bool {
	if index < 0 || index > l.len {
		return false
	}

	for i := index; i < l.len; i++ {
		(*l.ptr)[i] = (*l.ptr)[i + 1]
	}
	l.len--
	return true
}




func (l *SeqList) Get (index int) interface{} {
	if index < 0 || index > l.len {
		errors.New("index < 0 || index > st.len")
	}
	return (*l.ptr)[index]
}

// 顺序表更改元素的实现过程是：
	// 找到目标元素；
	// 直接修改该元素的值；
func (l *SeqList) Set (index int, elem interface{}) bool {
	if index < 0 || index > l.len {
		return false
	}
	(*l.ptr)[index] = elem
	return true
}

// 顺序表查找元素
//顺序表中查找目标元素，可以使用多种查找算法实现，比如说二分查找算法、插值查找算法等。
func (l *SeqList) Contain (elem interface{}) bool {
	for i := 0; i < l.len; i++ {
		if (*l.ptr)[i] == elem{
			return true
		}
	}
	return false
}


func (l *SeqList) Print() {
	if l.len == 0 {
		return
	}
	str := ""
	for idx, val := range *l.ptr {
		if idx == 0 && val != nil {
			str += fmt.Sprintf("[%d", val)
		} else {
			if val != nil {
				str += fmt.Sprintf(" %d", val)
			}
		}
	}
	str += "]"
	fmt.Println(str)
}

func main() {
	list := NewSeqList(5)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(1)

	fmt.Println(list.Insert(9, 0))

	list.Insert(0, 9)
	list.Print()

	list.Del(0)
	list.Print()

	fmt.Println(list.Len())
	fmt.Println(list.Cap())

	fmt.Println(list.Get(0))

	list.Set(0, 2)
	list.Print()

	fmt.Println(list.Contain(3))
	fmt.Println(list.Contain(2))
}
