package main

import (
	"fmt"
	"sort"
)

/*****************sort包介绍**********************/
/*
Go语言标准库sort包中实现了几种基本的排序算法：插入排序、快排和堆排序，但在使用sort包进行排序时无需具体考虑使用那种排序方式。
	func insertionSort(data Interface, a, b int)
	func heapSort(data Interface, a, b int)
	func quickSort(data Interface, a, b, maxDepth int)

sort.Interface接口定义了三个方法，注意sort包中接口Interface这个名字，是大写字母I开头，不要和interface关键字混淆，这里就是一个接口名而已。
	type Interface interface {
		// Len 为集合内元素的总数
		Len() int
		// 如果index为i的元素小于index为j的元素，则返回true，否则false
		Less(i, j int) bool
		// Swap 交换索引为 i 和 j 的元素
		Swap(i, j int)
	}
这三个方法分别是：
	获取数据集合长度的Len()方法
	比较两个元素大小的Less()方法
	交换两个元素位置的Swap()方法。只要实现了这三个方法，就可以对数据集合进行排序，sort包会根据实际数据自动选择高效的排序算法。

sort包原生支持[]int、[]float64和[]string三种内建数据类型切片的排序操作，即不必我们自己实现相关的Len()、Less()和Swap()方法。

以[]int为例，我们看看在sort包中的是怎么定义排序操作的:
	type IntSlice []int

	先通过 []int 来定义新类型IntSlice，然后在IntSlice上定义三个方法，Len()，Less(i, j int)，Swap(i, j int)，
实现了这三个方法也就意味着实现了sort.Interface。

	方法 func (p IntSlice) Sort() 通过调用 sort.Sort(p) 函数来实现排序。而p因为是sort.Interface类型，但IntSlice实现了这三个接口方法，
也是sort.Interface类型，因此可以直接调用得到排序结果。其他[]float64和[]string的排序也基本上按照这种方式来实现。

	func (p IntSlice) Len() int           { return len(p) }
	func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
	func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
	func (p IntSlice) Sort() { Sort(p) }
*/

// []int，[]string排序的实例：
// 默认结果都是升序排列
func main151() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println("a:", a) // a: [-14 -1 3 4 5 9 11]

	str := []string{"long", "ixfosa", "zhong"}
	sort.Strings(str)
	fmt.Println("str:", str) // str: [ixfosa long zhong]

	// func Reverse(data Interface) Interface
	// func Sort(data Interface)
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Printf("After reverse: %v\n", str) // After reverse: [zhong long ixfosa]

	// 如果我们想对一个 sortable object 进行逆序排序，可以自定义一个type。但 sort.Reverse 帮你省掉了这些代码。
	b := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Ints(b)
	fmt.Println("b:", b) // b: [1 2 3 4 5 6 7 8 9]
	sort.Sort(sort.Reverse(sort.IntSlice(b))) // b: [9 8 7 6 5 4 3 2 1]
	fmt.Println("b:", b)
}


/*****************相关方法：**********************/
/*
将类型为float64的slice以升序方式排序
	func Float64s(a []float64)

判定是否已经进行排序func Float64s(a []int)
	func Float64sAreSorted(a []float64) bool

Ints 以升序排列 int 切片。
	func Ints(a []int)

判断 int 切片是否已经按升序排列。
	func IntsAreSorted(a []int) bool

IsSorted 判断数据是否已经排序。包括各种可sort的数据类型的判断．
	func IsSorted(data Interface) bool


Strings 以升序排列 string 切片。
	func Strings(a []string)

判断 string 切片是否按升序排列
	func StringsAreSorted(a []string) bool

// search使用二分法进行查找，Search()方法回使用“二分查找”算法来搜索某指定切片[0:n]，
// 并返回能够使f(i)=true的最小的i（0<=i<n）值，并且会假定，如果f(i)=true，则f(i+1)=true，
// 即对于切片[0:n]，i之前的切片元素会使f()函数返回false，i及i之后的元素会使f()
// 函数返回true。但是，当在切片中无法找到时f(i)=true的i时（此时切片元素都不能使f()
// 函数返回true），Search()方法会返回n（而不是返回-1）。

// Search 常用于在一个已排序的，可索引的数据结构中寻找索引为 i 的值 x，例如数组或切片。
// 这种情况下实参 f一般是一个闭包，会捕获所要搜索的值，以及索引并排序该数据结构的方式。
func Search(n int, f func(int) bool) int

// SearchFloat64s 在float64s切片中搜索x并返回索引如Search函数所述.
// 返回可以插入x值的索引位置，如果x不存在，返回数组a的长度切片必须以升序排列
	func SearchFloat64s(a []float64, x float64) int

// SearchInts 在ints切片中搜索x并返回索引如Search函数所述. 返回可以插入x值的
// 索引位置，如果x不存在，返回数组a的长度切片必须以升序排列
	func SearchInts(a []int, x int) int

// SearchFloat64s 在strings切片中搜索x并返回索引如Search函数所述. 返回可以
// 插入x值的索引位置，如果x不存在，返回数组a的长度切片必须以升序排列
	func SearchStrings(a []string, x string) int

// 其中需要注意的是，以上三种search查找方法，其对应的slice必须按照升序进行排序，
// 否则会出现奇怪的结果．

// Sort 对 data 进行排序。它调用一次 data.Len 来决定排序的长度 n，调用 data.Less
// 和 data.Swap 的开销为O(n*log(n))。此排序为不稳定排序。他根据不同形式决定使用
// 不同的排序方式（插入排序，堆排序，快排）。
	func Sort(data Interface)

// Stable对data进行排序，不过排序过程中，如果data中存在相等的元素，则他们原来的
// 顺序不会改变，即如果有两个相等元素num, 他们的初始index分别为i和j，并且i<j，
// 则利用Stable对data进行排序后，i依然小于ｊ．直接利用sort进行排序则不能够保证这一点。
	func Stable(data Interface)
 */


/******************** 自定义sort.Interface排序*******************/
/*
	如果是具体的某个结构体的排序，就需要自己实现Interface了。
数据集合（包括自定义数据类型的集合）排序需要实现sort.Interface接口的三个方法，
即：Len()，Swap(i, j int)，Less(i, j int)，数据集合实现了这三个方法后，即可调用该包的Sort()方法进行排序。
Sort(data Interface) 方法内部会使用quickSort()来进行集合的排序。quickSort()会根据实际情况来选择排序方法。

任何实现了 sort.Interface 的类型（一般为集合），均可使用该包中的方法进行排序。这些方法要求集合内列出元素的索引为整数。
 */

type person struct {
	Name string
	Age int
}

type personSlice []person

func (s personSlice) Len() int {
	return len(s)
}

func (s personSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s personSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func main152()  {
	a := personSlice{
		person{"ixfosa", 22},
		person{"long", 21},
		person{"zhong", 20},
	}
	fmt.Println(a) // [{ixfosa 22} {long 21} {zhong 20}]
	sort.Sort(a)
	fmt.Println(a) // [{zhong 20} {long 21} {ixfosa 22}]

	sort.Stable(a)
	fmt.Println(a) // [{zhong 20} {long 21} {ixfosa 22}]
}


/****************** sort.Slice*********************/
/*
	利用sort.Slice 函数，而不用提供一个特定的 sort.Interface 的实现，而是 Less(i，j int) 作为一个比较回调函数，
可以简单地传递给 sort.Slice 进行排序。这种方法一般不建议使用，因为在sort.Slice中使用了reflect。
	func Slice(slice interface{}, less func(i, j int) bool)
 */
type student struct {
	Name string
	Age int
}

func main()  {
	a := []student{
		student{"ixfosa", 22},
		student{"long", 21},
		student{"zhong", 20},
	}
	sort.Slice(a, func(i, j int) bool {
		return a[1].Age > a[j].Age
	})
	fmt.Println(a) // [{ixfosa 22} {long 21} {zhong 20}]
}
