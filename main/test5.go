package main

import (
	"os"
	"fmt"
	"net"
	"strings"
	"bufio"
)

func main(){
	//test596()
	test599()
}
func test598(){
	//fmt.Println("-----",len(os.Args),"===",&os.Stderr)
	//if len(os.Args) != 2{
	//	fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	//	os.Exit(1)
	//}
	//name := os.Args[1]
	//fmt.Println("*=====*",name)
	ip,_ := net.InterfaceAddrs()
	for _,v := range ip{
		//fmt.Println(v)
		if ipnet,ok := v.(*net.IPNet);ok&&!ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil {
				fmt.Println("123333:", ipnet.IP.String())
			}
		}
	}
	fmt.Println(ip)
	addr :=net.ParseIP(ip[0].Network())
	fmt.Println(addr)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
func test597(){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
//TCP 客户端
func test599(){

	//返回一个拥有 默认size 的reader，接收客户端输入
	reader := bufio.NewReader(os.Stdin)

	for{
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取数据异常")
			continue;
		}
		//去除两端的空格
		input = strings.TrimSpace(input)
		if len(input)<=0{
			continue
		}
		networkType := networkType{"udp","10.3.24.213:7777"}
		go networkType.cConnHander594(input)
	}
}
func test596(){
	service:=":7777"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	if err != nil {
		panic("出错5")
	}
	listen,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil{
		panic("出错6")
	}
	for{
		conn,err := listen.Accept()
		if err != nil{
			continue
		}
		//daytime := time.Now().String()
		//conn.Write([]byte(daytime))
		//conn.Close()
		go connHandler595(conn)
	}
}
func connHandler595(conn net.Conn){
	if conn == nil {
		panic("无效的socket连接")
	}
	buf := make([]byte,4096)
	for{
		//数据读取完毕
		var cnt int;
		if cnt, err := conn.Read(buf);cnt ==0&&err ==nil {
			conn.Close()
			break
		}
		//对读取数据去空格操作
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		//去除读取数据的内部空格
		cInputs := strings.Split(inStr," ")
		//取第一条数据
		fCommand := cInputs[0]
		fmt.Println("客户端传输->" + fCommand)
		switch fCommand {
		case "ping": conn.Write([]byte("服务器端回复-> pong\n"))
		case "hello":conn.Write([]byte("服务器端回复-> world\n"))
		default:
			conn.Write([]byte("服务器端回复" + fCommand + "\n"))
		}
		//conn.Close()  客户端的连接关闭，此连接也将被强制关闭
		fmt.Printf("来自 %v 的连接关闭\n", conn.RemoteAddr())
	}

}

type networkType struct {
	networkType string
	address string
}
//客户端模拟发送请求代理函数
func (n networkType) cConnHander594(input string){
		//1.1与服务端建立连接
		conn,err := net.Dial(n.networkType,n.address)
		if err != nil{
			fmt.Println("客户端建立连接失败")
			return
		}
		//1.4 最后关闭连接
		defer func(){
			conn.Close()
		}()
		//接受服务端回传的信息
		buf := make([]byte,1024)

		//1.2向服务端发送请求数据
		if conn == nil{
			fmt.Println(input)
		}
		_,err = conn.Write([]byte(input))
		if err != nil{
			fmt.Println(err)
			return
		}
		//1.3读取服务端响应的数据
		cnt,err := conn.Read(buf)
		if err != nil{
			fmt.Printf("客户端读取数据失败 %s\n", err)
			return
		}

		fmt.Print("服务器端回复：" + string(buf[0:cnt]))
}