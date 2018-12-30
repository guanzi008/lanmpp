package main

import (
	"flag"
	"fmt"
	//"github.com/xiaoqidun/goini"
)

func main() {
	//ini := goini.NewGoINI()
	//if err := ini.LoadFile(`H:\go\guanzi008\test2.ini`); err != nil {
	//	fmt.Println("配置解析错误", err)
	//	return
	//} else {
	flag.Parse()
	//var a string = ini.GetString("a", "a", "a")
	//var b string = ini.GetString("hi", "url", "a")
	// var b = flag.String("qq","env","env")
	//		flag.Parse()
	//		flag.String("qq","gg","gg")
	//		flag.Parse()

	//flag.String("qq","gg","gg")
	url := flag.String("a", "a", "a")
	fmt.Println(url)

}

//
//func main() {
//	xuhao()
//cmd := exec.Command(a, b)
//ret, _ := cmd.CombinedOutput()
//data := string(ret)
//fmt.Println(data,)

//}
