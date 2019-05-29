package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)
type userinfo struct {
	uid int
	username string
	depart string
	creatd string
}
func main(){
	//fmt.Println("测试主函数调用")
	http.HandleFunc("/",sayHello)//设置访问路由
	http.HandleFunc("/login",login)//设置访问路由
	//err := http.ListenAndServe(":9090",nil)
	err := http.ListenAndServe(":9090",nil)//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//for {
	//	time.Sleep(time.Second*5)
	//	go getBTCPrice();
	//}
	//getBTCPrice();
	//testPrintPage();

	//dbUtil()

}
func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET"{
		t,_ := template.ParseFiles("F:/TongXun/golangp/main/login.gtpl")
		t.Execute(w,nil)
	}else {
		//默认情况下，Handler里面是不
		//会自动解析form的，必须显式的调用r.ParseForm()后才能对这个表单数据进行操作
		r.ParseForm()
		//请求的是表单提交
		//var username = r.FormValue("username")
		//var password = r.FormValue("password")
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		u := userinfo{username:r.FormValue("username"),depart:r.FormValue("password")}
		u.dbUtil()

		expiration := time.Now().AddDate(1,0,0)
		//expiration.Year()+1
		cookie := http.Cookie{Name:"地址",Value:"222",Expires: expiration}
		http.SetCookie(w,&cookie);

	}
}
func (user *userinfo) dbUtil(){
	var username = user.username
	var password = user.depart
	db,err := sql.Open("mysql","root:root@tcp(localhost:3306)/webcron?charset=utf8");
	if(err != nil){
		fmt.Println(err)
		panic(err)
	}
	//插入数据
	stmt,err := db.Prepare("Insert into userinfo(username,departname,created) values(?,?,?)")
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	res,err := stmt.Exec(username,username,"2019-05-16")
	checkErr(err)
	id,err := res.LastInsertId();
	fmt.Println("最新插入数据的ID:",id)
	//更新数据
	stmt,err = db.Prepare("UPDATE userinfo set departname = ? where uid = ?")
	checkErr(err)
	res,err = stmt.Exec(password,id)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("更新数据成功条数：",affect)

	//查询数据
	stmt,err = db.Prepare("select * from userinfo where uid = ?")
	checkErr(err)
	rows,err := stmt.Query(id)
	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("UID:",uid)
		fmt.Println("用户名：",username)
		fmt.Println("机构：",department)
		fmt.Println("创建时间",created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")

	checkErr(err)

	res, err = stmt.Exec(id)

	checkErr(err)

	affect, err = res.RowsAffected()

	checkErr(err)
	fmt.Println("删除成功条数：",affect)

	db.Close()
}
func checkErr(er error){
	if (er != nil){
		fmt.Println(er)
		panic(er)
	}
}
func sayHello(w http.ResponseWriter,r *http.Request){
	r.ParseForm()  //解析参数
	fmt.Println(r.Form)//输出到服务端的打印信息
	fmt.Println("Scheme",r.URL.Scheme)
	fmt.Println("Path",r.URL.Path)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("Key:",k)
		fmt.Println("Value",v)
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func displayLoading(silent bool){
	if silent {return}
	loading := []string{"loading    ","loading .  ","loading .. ","loading ..."}
	for i := 0;;{
		fmt.Printf("\r%c[1;0;32m%s%c[0m",0x1B, loading[i],0x1B)
		time.Sleep(time.Millisecond * 150)
		i = i%3+1
	}
}
func getBTCPrice(){

	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil{

	}
	body, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	err1:= 	json.Unmarshal(body,&result)
	if err1 != nil{}
	var buyArry []float64;
	var i=0;
	for k,v := range result{
		if k != "CNY" {
			continue
		}
		v3 := v.(map[string]interface{})
		//fmt.Printf("币种：%v 15m价格：%f 买入价：%f 卖出价：%f 最终价：%f 币种符号：%v \n",k,v3["15m"],v3["buy"],v3["sell"],v3["last"],v3["symbol"])
		if k == "CNY" {
			break
		}
		buyArry[i] = float64(v3["buy"].(float64));
		testPrintPage(buyArry);
	}
}
func testPrintPage(data []float64){
	//var data []float64 = []float64{11.2,12.3,14.5,18.9,8.0}
	//b := 0
	for i:=0;i<len(data);i++{

		//for k:=i;k<len(data);k++{
		//	if(data[b]>=data[k]){
		//		b = i
		//	}else{
		//		b = k
		//	}
		//}
		var bi int = int(data[i]/0.01)/100;
		for l := 1;l <= bi;l++ {
			if l == bi {
				fmt.Printf("#\n")
				continue
			}
			fmt.Printf("*")
		}

	}
	//var bi int = int(data[b]/0.01)/100;
	//for l := 1;l <= bi;l++ {
	//	if l == bi {
	//		fmt.Printf("#\n")
	//		continue
	//	}
	//	fmt.Printf("*")
	//}
	//fmt.Printf("*",data[b])

}
