package main

import (
	"os"
	"fmt"
	"strings"
)

func main(){
	test99()
}
//文件读取操作
func test99(){
	//os.Mkdir("Web.xml",0777)
	os.MkdirAll("web/web-info/make/def",0777)
	//os.Remove("web.xml")
	//os.RemoveAll("web")
	os.Remove("web/web-info/make/def/web.xml")
	if ws,err := os.Create("web/web-info/make/def/web.xml");err == nil{
		ws.Write([]byte("请问之后怎么走！\n"))
		ws.WriteAt([]byte("我该怎么办\n"),100)
		ws.WriteString("无话可说\n")
	}
	if rd,err := os.Open("web/web-info/make/def/web.xml");err == nil{
		bte := make([]byte,1024);
		f,_ := rd.Read(bte)
		fmt.Println(f)
		fmt.Println("", strings.Fields (string(bte)))
	}

}