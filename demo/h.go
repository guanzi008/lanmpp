package main

import (
	"fmt"
	"github.com/golang/net/proxy"
	"io/ioutil"
	"log"
	"net/http"
	"os"

)

func main() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9742", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	if resp, err := httpClient.Get("https://www.google.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}


















//package main
//
//import "fmt"
//
//type xuhao struct {
//	id  int
//	url string
//}
//
//func main() {
//	//	xuhao{id:100,url:"odata.cc"}
//	fmt.Println(xuhao{id: 100, url: "odata.cc"})
//	var xuhao1 xuhao
//	xuhao1.id = 101
//	xuhao1.url = "1q.tn"
//	fmt.Println(xuhao1.id)
//	Run(&xuhao1)
//}
//func Run(st_tmp *xuhao) {
//	fmt.Println(st_tmp.id)
//}
