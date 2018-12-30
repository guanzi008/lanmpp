package main

import (
	"fmt"
	"net/http"
)

func main() {
	//	flag.Parse()
	//生成client 参数为默认
	client := &http.Client{}
	//	var a = flag.String("qq","env","env")
	//生成要访问的url
	url := "https://ip.cn/"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//	stdout := os.Stdout
	//	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)

}
