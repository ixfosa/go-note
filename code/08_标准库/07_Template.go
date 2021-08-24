package main

import (
	"html/template"
	"net/http"
)

// Template
	// html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。
	// 它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用text/template包。


/******************模板示例*********************/
/*
	通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。
	模板中的注释引用数据接口的元素（一般如结构体的字段或者字典的键）来控制执行过程和获取需要呈现的值。
 	模板执行时会遍历结构并将指针表示为’.‘（称之为”dot”）指向运行过程中数据结构的当前位置的值。

	用作模板的输入文本必须是utf-8编码的文本。”Action”—数据运算和控制单位—由”“界定；在Action之外的所有文本都不做修改的拷贝到输出中。
	Action内部不能有换行，但注释可以有换行。

 */

func test1(w http.ResponseWriter, r *http.Request)  {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./08_标准库/html/01_test.html")
	if err != nil {
		panic(err)
	}

	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "gogogo!!!")     // http://127.0.0.1:9090/
}

type UserInfo struct {
	Name string
	Gender string
	age int
}

func test2(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("./08_标准库/html/01_test2.html")
	if err != nil {
		panic(err)
	}

	user := UserInfo{
		Name:   "ixfosa",
		Gender: "man",
		age:    22,
	}
	tmpl.Execute(w, user)
}

/*****************模板语法**********************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */



func main() {
	// http.HandleFunc("/", test1)
	http.HandleFunc("/", test2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}