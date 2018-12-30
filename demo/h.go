package main

import "fmt"

type xuhao struct {
	id  int
	url string
}

func main() {
	//	xuhao{id:100,url:"odata.cc"}
	fmt.Println(xuhao{id: 100, url: "odata.cc"})
	var xuhao1 xuhao
	xuhao1.id = 101
	xuhao1.url = "1q.tn"
	fmt.Println(xuhao1.id)
	Run(&xuhao1)
}
func Run(st_tmp *xuhao) {
	fmt.Println(st_tmp.id)
}
