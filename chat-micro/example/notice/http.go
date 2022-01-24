package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//36.关闭 HTTP 的响应体
//使用 HTTP 标准库发起请求、获取响应时，即使你不从响应中读取任何数据或响应为空，都需要手动关闭响应体。新手很容易忘记手动关闭，或者写在了错误的位置：
func httpClose() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	defer resp.Body.Close()    // resp 可能为 nil，不能读取 Body
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

// 应该先检查 HTTP 响应错误为 nil，再调用 resp.Body.Close() 来关闭响应体：
func httpClosePlus() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	checkError(err)

	defer resp.Body.Close()    // 绝大多数情况下的正确关闭方式
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

// 绝大多数请求失败的情况下，resp 的值为 nil 且 err 为 non-nil。但如果你得到的是重定向错误，那它俩的值都是 non-nil，最后依旧可能发生内存泄露。2 个解决办法：
func httpCloseSuccess() {
	resp, err := http.Get("http://www.baidu.com")

	// 关闭 resp.Body 的正确姿势
	if resp != nil {
		defer resp.Body.Close()
	}

	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

//一些支持 HTTP1.1 或 HTTP1.0 配置了 connection: keep-alive 选项的服务器会保持一段时间的长连接。
//但标准库 “net/http” 的连接默认只在服务器主动要求关闭时才断开，所以你的程序可能会消耗完 socket 描述符。解决办法有 2 个，请求结束后：
//直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
//设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，此时 HTTP 标准库会主动断开连接。
//注意：
//若你的程序要向同一服务器发大量请求，使用默认的保持长连接。
//若你的程序要连接大量的服务器，且每台服务器只请求一两次，那收到请求后直接关闭连接。或增加最大文件打开数 fs.file-max 的值。
func main() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	checkError(err)

	req.Close = true
	//req.Header.Add("Connection", "close")    // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}