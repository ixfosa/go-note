package main

import "fmt"

// 定义：
	// 链表由一个个数据节点组成的，它是一个递归结构，要么它是空的，要么它存在一个指向另外一个数据节点的引用。

type LinkNode struct {
	Data interface{}
	NextNode *LinkNode
}

func main() {
	node1 := &LinkNode{
		Data: 1,
	}

	node2 := &LinkNode{
		Data: 2,
	}

	node3 := &LinkNode{
		Data: 3,
	}

	node1.NextNode = node2
	node2.NextNode = node3

	// 按顺序打印数据
	cur := node1
	len := 0
	for cur != nil {
		fmt.Println(cur.Data)
		len++
		cur = cur.NextNode
	}
	fmt.Println("len: ", len)
}
