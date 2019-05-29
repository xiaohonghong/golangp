package main

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/beevik/etree"
	"encoding/json"

	"errors"
	"strconv"
)

func main(){
	//parseXml();
	//parseXXMLEtree();
	//writeXmlEtree();
	//testJSON();
	//testJson2()
	//testJson3();
	testJson99()
}
type server struct{
	XMLname    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
	serverOpen xml.Name `xml:"serverOpen"`
}
type Recurlyservers struct {
	XMLName	xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs		[]server `xml:"server"`
	Description	string `xml:",innerxml"`
}
func parseXml(){
	file,err := os.Open("F:/TongXun/golangp/main/servers.xml")//for read access
	checkErr1(err)
	defer file.Close()
	data,err := ioutil.ReadAll(file)
	checkErr1(err)
	v:=Recurlyservers{}
	err = xml.Unmarshal(data,&v)
	checkErr1(err)
	//fmt.Println(v)
	s := v.Svs
	for index,s1 :=range s{
		fmt.Println("",index)
		fmt.Println("",s1)
	}


}
func checkErr1(err error){
	if err !=nil {
		panic(err)
	}
}
func parseXXMLEtree(){
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("F:/TongXun/golangp/main/servers.xml");err != nil{
		panic(err)
	}
	servers := doc.SelectElement("servers")
	//fmt.Println(servers)
	server:=servers.SelectElements("server")
	for _,s := range server{
		if s1 := s.SelectElement("serverOpen",);s1 != nil {
			fmt.Println("ui:",s1.SelectElement("ui").Text())
			fmt.Println("uL:",s1.SelectElement("ul").Text())
		}
		fmt.Println("222serverName",s.SelectElement("serverName").Text());
	}
}
func writeXmlEtree(){

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)
	people := doc.CreateElement("People")
	people.CreateComment("These are all known people")

	jon := people.CreateElement("Person")
	jon.CreateAttr("name", "Jon")
	jon.CreateText("你好！")

	sally := people.CreateElement("Person")
	sally.CreateAttr("name", "Sally")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
type server2 struct {
	ServerName string `json:"ServerName"`
	ServerIp   string `json:"ServerIp"`
}
type server1 struct{
	server2
	ServerArr []server2
}
type servers struct {
	Servers1 []server1
}
func testJSON(){
	var ss servers
	str := `{"servers1":[{"serverName":"北京","serverIp":"127.0.0.9","serverArr":[{"serverName":"IIY","serverIp":"128.0.0.2"},{"serverName":"PPT","serverIp":"128.0.0.5"}]},{"serverName":"SHANGHAI","serverIp":"127.0.0.10"}]}`
	json.Unmarshal([]byte(str),&ss);
	fmt.Println(ss)
	}
	//类型断言 + interface{}类型结息JSON
func testJson2(){
	str := []byte(`{"Name":"Wednesday","Age":6,"Books":["科学技术","语文书","数学书"],"Parents":[{"Name":"Paper","Age":36},{"Name":"Lisa","Age":35}]}`)
	var f interface{}
	json.Unmarshal(str,&f)
	m := f.(map[string] interface{})
	for k,v := range m{
		switch vv:=v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for _, up := range vv {

				if k2,err:= up.(map[string] interface{});err {
					for k3, v2 := range k2 {
						switch vv2 := v2.(type) {
						case string:
							fmt.Println(k3, "is string", vv2)
						case float64:
							fmt.Println(k3, "is string", vv2)
						default:
							fmt.Println("错误的数据类型")
						}
					}
				}else{
					fmt.Println("up",up)
				}
			}
		default:
			fmt.Println("错误的数据类型")
		}
	}
}
type ss struct {
	bbs interface{}
}
func (s *ss) get(str string)(*ss){
	var s1 ss
	if k,err := s.bbs.(map[string] interface{});err{
		s1.bbs = k[str]
	}
	return &s1
}
func (s *ss) MustString() string{
	return s.bbs.(string)
}
func (s *ss) Int() (int){
	return s.bbs.(int);
}
func (s *ss) Array() ([]interface{}){
	bbs := s.bbs.([]interface{})
	return bbs;
}
//创建对象
func NewJson(bte []byte) (ss,error){
	var f  interface{};
	var s ss;
	json.Unmarshal(bte,&f)
	if f == nil {
		return ss{},errors.New("JSON字符串格式转换异常！请检查是否是正确的json格式");
	}
	s.bbs = f;
	return s,nil
}
func testJson3(){
	str := `{"Test":"tttddd","hhh":{"k1":"ddd","k2":"uiy","ddeaz":["zzdd","oixjd","111",2333]},"jjj":["ddd","zzz",1233]}`;
	if sss,err := NewJson([]byte(str));err==nil{
		fmt.Println((sss.get("hhh").get("ddeaz").Array()))
		bbb := sss.get("hhh").get("ddeaz").Array()
		for k,v := range bbb{
			println(fmt.Sprintf("%v=%v", k, v))
		}
	}else{
		fmt.Println(err)
	}
}
func testJson4(){
	var m1 map[string] interface{}
	m1["sss"] = "xxx";
	ss,_ := json.Marshal(m1);
	fmt.Println(string(ss))
}
//字符串转换
func testJson99(){
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 0, 10)
	str = strconv.AppendInt(str,334455,10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}