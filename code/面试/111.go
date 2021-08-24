package main

import "fmt"

// 下面的代码有什么问题？

type N int

func (n N) value() ()  {
	fmt.Printf("v:%p, %v\n",&n, n)
}

func (n *N) pointer() ()  {
	fmt.Printf("v:%p, %v\n",n, *n)
}

func main() {
	var a N = 25

	// p := &a
	// p1 := &p
	a.value()
	a.pointer()

	// p1.value()
	// p1.pointer()
	// calling method value with receiver p1 (type **N) requires explicit dereference
	// calling method pointer with receiver p1 (type **N) requires explicit dereference
}
