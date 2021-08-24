package main

import "fmt"

// 分而治之，就是把一个复杂的问题分成两个或更多的相同或相似的子问题。
// 直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。
// 分治法一般使用递归来求问题的解。

// 递归: 递归就是不断地调用函数本身。

// 比如求阶乘 1 * 2 * 3 * 4 * 5 *...* N：
func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	// 函数不断地调用本身，并且还乘以一个变量：n * Rescuvie(n-1)，这是一个递归的过程。
	// 因为递归式使用了运算符，每次重复的调用都使得运算的链条不断加长，系统不得不使用栈进行数据保存和恢复。
	return Rescuvie(n - 1) * n
}
/*
	会反复进入一个函数，它的过程如下:
		Rescuvie(5)
		{5 * Rescuvie(4)}
		{5 * {4 * Rescuvie(3)}}
		{5 * {4 * {3 * Rescuvie(2)}}}
		{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
		{5 * {4 * {3 * {2 * 1}}}}
		{5 * {4 * {3 * 2}}}
		{5 * {4 * 6}}
		{5 * 24}
		120
 */

// 尾递归
	// 如果每次递归都要对越来越长的链进行运算，那速度极慢，并且可能栈溢出，导致程序奔溃。 所以有另外一种写法，叫尾递归：

	// 尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。

	// 如果一个函数中所有递归形式的调用都出现在函数的末尾，我们称这个递归函数是尾递归的。
	// 当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，这个递归调用就是尾递归。
	// 尾递归函数的特点是在回归过程中不用做任何操作，这个特性很重要，因为大多数现代的编译器会利用这种特点自动生成优化的代码

func RescuvieTail(n, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n - 1, a * n)
}
/*
	递归过程如下:
		RescuvieTail(5, 1)
		RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
		RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
		RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
		RescuvieTail(1, 60*2)=RescuvieTail(1, 120)
		120
 */


// 例子：斐波那契数列
	// 斐波那契数列是指，后一个数是前两个数的和的一种数列。如下：
	// 1 1 2 3 5 8 13 21 ... N-1 N 2N-1

// 递推
func FBA1(n int) int {
	if n <= 2 {
		return 1
	}

	res := 0
	f, g := 0, 1

	for i := 1; i < n; i++ {
		res = f + g
		f = g
		g = res
	}
	return res
}


// 递归的求解为：
func FBA2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FBA2(n - 2) + FBA2(n - 1)
}


// 尾递归的求解为：
func FBA3(n, a, b int) int {
	if n <= 2 {
		return b
	}
	return FBA3(n - 1, b, a + b)
}


// 例子：二分查找
	// 在一个已经排好序的数列，找出某个数，如：
	// 1 5 9 15 81 89 123 189 333
func BinarySearch(array []int, target int, l, r int) int {
	for {
		if l > r {
			return -1
		}
		mid := (l + r) / 2
		if target < array[mid] {
			r = mid - 1
		} else if target > array[mid] {
			l = mid + 1
		} else {
			return mid
		}

	}
}
func main() {
	fmt.Println(Rescuvie(5)) // 120
	fmt.Println(RescuvieTail(5, 1))


	fmt.Println("递推: ", FBA1(10))
	fmt.Println("递归: ", FBA2(10))
	fmt.Println("尾递归: ", FBA3(10, 1, 1))

	a := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	fmt.Println("BinarySearch: ", BinarySearch(a, 333, 0, len(a) - 1))
}
