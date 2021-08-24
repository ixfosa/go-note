package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	for k, v :=range r.Form{
		fmt.Printf("key:%v, value:%v\n", k, v)
	}
	fmt.Println(w, "hello!")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method) // 获取请求的方法

	if r.Method == "GET" {
		t, err := template.ParseFiles("./gtpl/login.gtpl")
		if err != nil {
			fmt.Println("template.ParseFiles: ", err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()  // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main011() {
	//http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

/*
request.Form 是一个 url.Values 类型，里面存储的是对应的类似 key=value 的信息，下面展示了可以对 form 数据进行的一些操作:
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

Tips:
	Request 本身也提供了 FormValue () 函数来获取用户提交的参数。如 r.Form ["username"] 也可写成 r.FormValue ("username")。
	调用 r.FormValue 时会自动调用 r.ParseForm，所以不必提前调用。r.FormValue 只会返回同名参数中的第一个，若参数不存在则返回空字符串。
*/
