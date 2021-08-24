package main

import (
	"fmt"
)

// 1.单链表，就是链表是单向的，可以一直往下找到下一个数据节点，它只有一个方向，它不能往回找。

// 2.双链表，每个节点既可以找到它之前的节点，也可以找到之后的节点，是双向的。

// 3.循环链表，就是它一直往下找数据节点，最后回到了自己那个节点，形成了一个回路。
// 循环单链表和循环双链表的区别就是，一个只能一个方向走，一个两个方向都可以走。

// 循环链表
type Ring struct {
	next, prev *Ring   // 前驱和后驱节点
	Value interface{}  // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
// 因为绑定前驱和后驱节点为自己，没有循环，时间复杂度为：O(1)。
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
// 会连续绑定前驱和后驱节点，时间复杂度为：O(n)。
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取上一个或下一个节点
// 获取前驱或后驱节点，时间复杂度为：O(1)。
func (r *Ring) Next() *Ring {
	if r == nil {
		r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r == nil {
		r.init()
	}
	return r.prev
}


// 获取第 n 个节点
// 因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：
// 因为需要遍历 n 次，所以时间复杂度为：O(n)。
func (r *Ring) Move(n int) *Ring {
	if r == nil {
		r.init()
	}

	switch {
	case n < 0:
		for ; n > 0; n++ {
			r = r.prev
		}
	case n >0:
		for ; n < 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点当前节点，链接一个节点，并且返回当前节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后面的 n 个节点
func (r *Ring) UnLink(n int) *Ring {

	return r
}

func (r *Ring) Print() {
	node := r
	str := "["
	for {
		str += fmt.Sprintf("%v, ",  node.Value)
		node = node.next

		if node == r {
			break
		}
	}
	str = string([]byte(str)[:len(str)-2]) + "]"
	fmt.Println(str)
}
func main() {

	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})


	r.Print()


}
