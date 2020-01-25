package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args)>1{
		url := "https://apikey.net/?type=json&ip="
		domain := os.Args[1]
		req, _ := http.NewRequest("GET", url+domain, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		//fmt.Println(res)
		fmt.Println(string(body))
	} else {
		fmt.Println("")
	}
}
//if os.Args[1]==""{
//return
//}

