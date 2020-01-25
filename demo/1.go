package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
		if len(os.Args) > 1 {
			req, _ := http.NewRequest("GET", "https://whois.aite.xyz/?ajax&client&domain="+os.Args[1], nil)
			res, _ := http.DefaultClient.Do(req)
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body)) //fmt.Println(res)
			os.Exit(0)
		} else {
			fmt.Println("")
			os.Exit(0)
		}
	}

//名称: api1q
//ID:114850
//Token:b6cb4634446b58007d895a3a916ed81d
//创建时间: 2019-09-04 21:25:22
//curl -X POST https://dnsapi.cn/Record.Modify -d 'login_token=114850,b6cb4634446b58007d895a3a916ed81d&format=json&domain=1q.tn&record_id=454815654&sub_domain=api&value=3.2.2.2&record_type=A&record_line= 默认'