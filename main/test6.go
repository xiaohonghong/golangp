package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)
func main(){
	//l := listenType{"udp4",":7777"}
	//l.test699()
	l := listenType{"udp",":7777"}
	l.test694()
}
func (l *listenType) test699(){

	tcpAddr,err := net.ResolveTCPAddr("tcp4",":7777")
	if err != nil {
		panic("出错5")
	}
	listen,err := net.ListenTCP(l.network,tcpAddr)
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		conn,err := listen.Accept()
		if err != nil{
			continue
		}
		daytime := time.Now().String()
		fmt.Println(daytime)
		//conn.Write([]byte(daytime))
		//conn.Close()
		go connHandler695(conn)
	}
}
func connHandler695(conn net.Conn){
	if conn == nil {
		panic("无效的socket连接")
	}
	buf := make([]byte,4096)
	for{
		//数据读取完毕
		var cnt int;
		if cnt,_ = conn.Read(buf);cnt == 0 {
			fmt.Println("....")
			//conn.Close()
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
			conn.Write([]byte(fCommand + "\n"))
		}
		//conn.Close()  客户端的连接关闭，此连接也将被强制关闭
		//fmt.Printf("来自 %v 的连接关闭\n", conn.RemoteAddr())
	}

}
type listenType struct {
	network string
	address string
}
func (l *listenType) test694(){

	tcpAddr,err := net.ResolveUDPAddr("udp4",":7777")
	if err != nil {
		panic("出错5")
	}
	conn,err := net.ListenUDP(l.network,tcpAddr)
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		//conn,err := listen.ReadFromUDP()
		//if err != nil{
		//	continue
		//}
		daytime := time.Now().String()
		fmt.Println(daytime)
		//conn.Write([]byte(daytime))
		//conn.Close()
		go connHandler695(conn)
	}
}