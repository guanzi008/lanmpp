package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

//定义一个字符串变量，并制定默认值以及使用方式
var a = flag.String("domain", "1q.tn", "域名whois查询")

//定义一个int型字符
//var b  = flag.Int("ins", 1, "ins nums")

func main() {
	// 上面定义了两个简单的参数，在所有参数定义生效前，需要使用flag.Parse()来解析参数
	flag.Parse()
	client := &http.Client{}
	// 测试上面定义的函数
	//	api := "odata.cc"
	url := string(*a)

	//提交请求
	reqest, err := http.NewRequest("GET", "http://whois.aite.xyz/?ajax&client&domain="+url, nil)
	//reqest, err := http.Get("http://whois.aite.xyz/?ajax&client&domain="+url,b)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)

	//fmt.Println("a string flag:",string(*a))
	//fmt.Println("ins num:",rune(*b))
}
