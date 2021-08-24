package main

import "fmt"

/*
关于 channel 下面描述正确的是？
	A. 向已关闭的通道发送数据会引发 panic；
	B. 从已关闭的缓冲通道接收数据，返回已缓冲数据或者零值；
	C. 无论接收还是接收，nil 通道都会阻塞；
// ABC
 */

// 下面的代码有几处问题？请详细说明。
	// 有两处问题：
		// 1.直接返回的 T{} 不可寻址；
		// 2.不可寻址的结构体不能调用带结构体指针接收者的方法；
type Test struct {
	n int
}

func (t Test) Set1(n int)  {
	t.n = n
	fmt.Println(t.n)
}

func (t *Test) Set2(n int)  {
	t.n = n
	fmt.Println(t.n)
}

func GetTest() Test {
	return Test{}
}

func main() {
	GetTest().Set1(2) // 2

	// GetTest().Set2(2) // cannot call pointer method on GetTest(),  cannot take the address of GetTest()

	x := GetTest()
	x.Set2(4)   // 4
}
