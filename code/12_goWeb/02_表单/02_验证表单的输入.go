package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*

	开发 Web 的一个原则就是，不能信任用户输入的任何信息，所以验证和过滤用户的输入信息就变得非常重要，
我们经常会在微博、新闻中听到某某网站被入侵了，存在什么漏洞，这些大多是因为网站对于用户输入的信息没有做严格的验证引起的，
所以为了编写出安全可靠的 Web 程序，验证表单输入的意义重大。

必填字段
	内置函数 len 可以获取字符串的长度
	if len(r.Form["username"][0])==0{
		// 为空的处理
	}

	r.Form 对不同类型的表单元素的留空有不同的处理， 对于空文本框、空文本区域以及文件上传，元素的值为空值，
而如果是未选中的复选框和单选按钮，则根本不会在 r.Form 中产生相应条目，如果我们用上面例子中的方式去获取数据时程序就会报错。
所以我们需要通过 r.Form.Get() 来获取值，因为如果字段不存在，通过该方式获取的是空值。但是通过 r.Form.Get() 只能获取单个的值，
如果是 map 的值，必须通过上面的方式来获取。

数字
	想要确保一个表单输入框中获取的只能是数字, 是判断正整数， 我们先转化成 int 类型，然后进行处理

	getint,err:=strconv.Atoi(r.Form.Get("age"))
	if err!=nil{
		// 数字转化出错了，那么可能就不是数字
	}

	// 接下来就可以判断这个数字的大小范围了
	if getint >100 {
		// 太大了
	}
还有一种方式就是正则匹配的方式
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		return false
	}

中文
	有时候我们想通过表单元素获取一个用户的中文名字，但是又为了保证获取的是正确的中文，我们需要进行验证，而不是用户随便的一些输入。
对于中文我们目前有两种方式来验证，可以使用 unicode 包提供的 func Is(rangeTab *RangeTable, r rune) bool 来验证，
也可以使用正则方式来验证，这里使用最简单的正则方式，如下代码所示
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		return false
	}


英文
期望通过表单元素获取一个英文值，例如我们想知道一个用户的英文名，应该是 astaxie，而不是 asta 谢。
我们可以很简单的通过正则验证数据：
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		return false
	}


电子邮件地址
你想知道用户输入的一个 Email 地址是否正确，通过如下这个方式可以验证：
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("no")
	}else{
		fmt.Println("yes")
	}


手机号码
你想要判断用户输入的手机号码是否正确，通过正则也可以验证：
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		return false
	}


下拉菜单
如果我们想要判断表单里面 <select> 元素生成的下拉菜单中是否有被选中的项目。
我们的 select 可能是这样的一些元素
	<select name="fruit">
	<option value="apple">apple</option>
	<option value="pear">pear</option>
	<option value="banana">banana</option>
	</select>
那么我们可以这样来验证
	slice:=[]string{"apple","pear","banana"}
	v := r.Form.Get("fruit")
	for _, item := range slice {
		if item == v {
			return true
		}
	}
	return false


单选按钮
<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女
那我们也可以类似下拉菜单的做法一样
	slice:=[]string{"1","2"}
	for _, v := range slice {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	return false


复选框
有一项选择兴趣的复选框，你想确定用户选中的和你提供给用户选择的是同一个类型的数据。
	<input type="checkbox" name="interest" value="football">足球
	<input type="checkbox" name="interest" value="basketball">篮球
	<input type="checkbox" name="interest" value="tennis">网球
对于复选框我们的验证和单选有点不一样，因为接收到的数据是一个 slice
	slice:=[]string{"football","basketball","tennis"}
	a:=Slice_diff(r.Form["interest"],slice)
	if a == nil{
		return true
	}
	return false
上面这个函数 Slice_diff 包含在我开源的一个库里面 (操作 slice 和 map 的库)，github.com/astaxie/beeku


日期和时间
你想确定用户填写的日期或时间是否有效。例如，用户在日程表中安排 8 月份的第 45 天开会，或者提供未来的某个时间作为生日。
Go 里面提供了一个 time 的处理包，我们可以把用户的输入年月日转化成相应的时间，然后进行逻辑判断
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	fmt.Printf("Go launched at %s\n", t.Local())
获取 time 之后我们就可以进行很多时间函数的操作。具体的判断就根据自己的需求调整。


身份证号码
如果我们想验证表单输入的是否是身份证，通过正则也可以方便的验证，但是身份证有 15 位和 18 位，我们两个都需要验证
	// 验证 15 位身份证，15 位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		return false
	}

	// 验证 18 位身份证，18 位前 17 位为数字，最后一位是校验位，可能为数字或字符 X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		return false
	}
*/

func getParameter(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Method: ", r.Method)
	if "GET" == r.Method {
		t, err := template.ParseFiles("./12_goWeb/gtpl/form.gtpl")
		if err != nil {
			fmt.Println("template.ParseFiles err: ", err)
			return
		}
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("Method: ", r.Method)
		r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		fmt.Fprintln(w, "hello go")

		fmt.Printf("r.Form: %T\n", r.Form)   // r.Form: url.Values type Values map[string][]string

		for k, v := range r.Form {
			fmt.Printf("r.Form: %T, k: %v, v: %v\n",r.Form, k, v)
			// r.Form: url.Values, k: city, v: [夏畈]
			// r.Form: url.Values, k: jianjie, v: [zhong]
			// r.Form: url.Values, k: username, v: [long]
			// r.Form: url.Values, k: password, v: [ixfosa]
			// r.Form: url.Values, k: sex, v: [男]
			// r.Form: url.Values, k: hobby, v: [睡觉 吃饭 打游戏]
		}
	}
}

// 127.0.0.1:9090/get?name=ixfosa&sex=男
func get(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Method: ", r.Method)
	r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	for k, v := range r.Form {
		fmt.Printf("r.Form: %T, k: %v, v: %v\n",r.Form, k, v)
		// r.Form: url.Values, k: name, v: [ixfosa]
		// r.Form: url.Values, k: sex, v: [男]
	}
	fmt.Fprintln(w, "hello go")
}

func main() {
	//http.HandleFunc("/parameter", getParameter)
	http.HandleFunc("/get", get)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("http.ListenAndServe err: ", err)
	}


}
