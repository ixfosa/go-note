package main

import "fmt"

// 下面代码有什么不规范的地方吗？
func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	// 检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。
	if v := x["two"]; v == "" {
		fmt.Println("no entry") // no entry
	}

	// 最可靠的操作是使用访问 map 时返回的第二个值。
	if v, ok := x["two"]; ok {
		fmt.Println("no entry", v) // no entry
	}
}
