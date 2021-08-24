package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
	防止多次递交表单

解决方案是在表单中添加一个带有唯一值的隐藏字段。在验证表单时，先检查带有该唯一值的表单是否已经递交过了。
如果是，拒绝再次递交；如果不是，则处理表单进行逻辑处理。另外，如果是采用了 Ajax 模式递交表单的话，当表单递交后，
通过 javascript 来禁用表单的递交按钮。

*/

func formRepeat(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("method:", r.Method) // 获取请求的方法

	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("./12_goWeb/gtpl/formRepeat.gtpl")
		if err != nil {
			fmt.Println("template.ParseFiles: ", err)
		}
		log.Println(t.Execute(w, token))
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断

		err := r.ParseForm()  // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 的合法性

		} else {
			// 不存在 token 报错
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
		fmt.Println("t0:", token)
	}
}
func main() {
	http.HandleFunc("/", formRepeat)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err: ", err)
	}
}
