package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value: ", v)
	}
	fmt.Fprintf(w, "hello long")
}
func main() {

	http.HandleFunc("/", sayHello)  // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

/*
在浏览器输入 http://localhost:9090
可以看到浏览器页面输出了 hello long
服务器端输出
	map[url_long:[111 222]]
	path /
	Scheme
	[111 222]
	key: url_long
	value:  [111 222]
	map[]
	path /favicon.ico
	Scheme
	[]


http://localhost:9090/?url_long=111&url_long=222
	Scheme
	[111 222]
	key: url_long
	value:  [111 222]
	map[]
	path /favicon.ico
	Scheme
	[]
*/