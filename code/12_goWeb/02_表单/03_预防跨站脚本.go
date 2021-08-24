package main
/*
	动态站点会受到一种名为 “跨站脚本攻击”（Cross Site Scripting, 安全专家们通常将其缩写成 XSS）的威胁，
而静态站点则完全不受其影响。

	攻击者通常会在有漏洞的程序中插入 JavaScript、VBScript、 ActiveX 或 Flash 以欺骗用户。
一旦得手，他们可以盗取用户帐户信息，修改用户设置，盗取 / 污染 cookie 和植入恶意广告等。

对 XSS 最佳的防护应该结合以下两种方法：
	一是验证所有输入数据，有效检测攻击;
	一个是对所有输出数据进行适当的处理，以防止任何已成功注入的脚本在浏览器端运行。

那么 Go 里面是怎么做这个有效防护的呢？Go 的 html/template 里面带有下面几个函数可以帮你转义

	func HTMLEscape (w io.Writer, b [] byte) // 把 b 进行转义之后写到 w
	func HTMLEscapeString (s string) string // 转义 s 之后返回结果字符串
	func HTMLEscaper (args ...interface {}) string // 支持多个参数一起转义，返回结果字符串
*/

func main() {
	
}
