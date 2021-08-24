package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
要使表单能够上传文件，首先第一步就是要添加 form 的 enctype 属性，enctype 属性有如下三种情况:
	application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
	multipart/form-data   不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
	text/plain    空格转换为 "+" 加号，但不对特殊字符编码。


	处理文件上传我们需要调用 r.ParseMultipartForm，里面的参数表示 maxMemory，调用 ParseMultipartForm 之后，
上传的文件存储在 maxMemory 大小的内存里面，如果文件大小超过了 maxMemory，那么剩下的部分将存储在系统的临时文件中。
我们可以通过 r.FormFile 获取上面的文件句柄，然后实例中使用了 io.Copy 来存储文件。

	获取其他非文件字段信息的时候就不需要调用 r.ParseForm，因为在需要的时候 Go 自动会去调用。
而且 ParseMultipartForm 调用一次之后，后面再次调用不会再有效果。

上传文件主要三步处理：
	表单中增加 enctype="multipart/form-data"
	服务端调用 r.ParseMultipartForm, 把上传的文件存储在内存和临时文件中
	使用 r.FormFile 获取文件句柄，然后对文件进行存储等处理。


*/
func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method) // 获取请求的方法

	if "GET" == r.Method {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("./12_goWeb\\gtpl\\upload.gtpl")

		if err != nil {
			fmt.Println("template.ParseFiles err: ", err)
			return
		}
		t.Execute(w, token)
	} else {
		// // ParseMultipartForm parses a request body as multipart/form-data.
		// func (r *Request) ParseMultipartForm(maxMemory int64) error
		r.ParseMultipartForm(32 << 20)

		// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		// 文件 handler 是 multipart.FileHeader, 里面存储了如下结构信息
		// type FileHeader struct {
		//	Filename string
		//	Header   textproto.MIMEHeader
		//	Size     int64
		//	content []byte
		//	tmpfile string
		// }

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("r.FormFile", err)
			return
		}

		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)  //map[Content-Disposition:[form-data; name="uploadfile"; filename="循环航班代码.png"] Content-Type:[image/png]]
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println("os.OpenFile err: ", err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main051() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil)
}

func postFile(filename ,targetUrl string) error  {
	bodyBuf := &bytes.Buffer{}

	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("bodyWriter.CreateFormFile err: ", err)
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return err
	}
	defer fh.Close()

	// io.copy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("io.Copy err: ", err)
		return err
	}

	contentType := bodyWriter.FormDataContentType()

	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("http.Post err: ", err)
		return err
	}

	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err: ", err)
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

// 客户端如何向服务器上传一个文件的例子，客户端通过 multipart.Write 把文件的文本流写入一个缓存中，
// 然后调用 http 的 Post 方法把缓存传到服务器。
func main()  {
	target_url := "http://localhost:9090/upload"
	filename := "./循环航班代码.png"
	postFile(filename , target_url)
}