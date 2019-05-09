package main

import "fmt"

const(
	id = 1
	name = "tutu"
	sex = "1"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	a1 = iota   //0
	b2          //1
	c3          //2
	d = "ha"   //独立值，iota += 1
	e          //"ha"   iota += 1
	f = 100    //iota +=1
	g          //100  iota +=1
	h = iota   //7,恢复计数
	i          //8

)
/*
	主函数
 */
func main(){
	//sayMessage();
	//testConst();
	//testItota();
	//testItota1();
	//testYunSuan1();
	//testNil2();

	//testArray();
	//testJieGouTi();
	//testQieMian();
	//testMap();
	startCalTxt()
}


func sayMessage(){
	/*
	常量不允许变更，只允许在方法类调用
	 */
	const PROD = "sex"
	fmt.Println("Hello Word")
	var str = "Hello Word"
	fmt.Println(str)
	var a,b = 12,13
	if a>b {
		fmt.Println("本次比较值最大的是：",a)
	}else{
		fmt.Println("本次比较值最大的是：",b)
	}
	switch a {
	case 10:fmt.Println("本次输出值：10")
	case 12:fmt.Println("本次输出值：13")
	default:
		fmt.Println("本次没有匹配的值")
	}

	for a<=15 {
		fmt.Println("本次循环变量值为：",a)
		a++
	}
}
/**
	常量枚举值
 */
func testConst(){
	fmt.Println(id,name,sex)
}
/*
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
 */
func testItota(){
	println("a b c",a,b,c)
}
func testItota1(){
	println("a1 b2 c3 d e f g h i",a,b,c,d,e,f,g,h,i)
}
/**
	测试位运算
 */
 const(
 	tu = 1<<iota
 	te = 1<<iota
 	tb = 1<<iota
 )
func testYunSuan1(){
	var b = 10
	b <<= 2
	fmt.Println(b,tu,te,tb)
}
func testNil2(){
	var cd = 4
	var ce = 4
	var ptr *int
	ptr = &cd
	//var hhh = 123;
	testNil(ptr,ce)
	fmt.Println("指针传递：prt=",cd)
	fmt.Println("值传递：ce=",ce)

	var nn = testFunction(1,3,4,"223")
	fmt.Println("输出返回结果：",nn)
}
func testNil(b *int, a int){
	//var bc int;
	*b = *b+122
	//bc = bc+122;
	a += a+122
	fmt.Println("指正传递：*b+122:",*b)
	fmt.Println("值传递：a+122=",a)
	fmt.Println("指正指向的内存地址：*b=",b)
}
/**
	测试带返回值的函数
 */
func testFunction(a,b,c int,d string)(bool){
	var ss = true
	fmt.Println("分别输出值：a+b+c=",a+b+c,"输出传递的字符串：str=",d)
	return ss
}
/*
	数组语法
 */
func testArray(){
	var sArray = [5]float32{1.22,3.44,9.99,5.66,7.88}
	var i int
	for i=0 ; i<len(sArray); i++ {
		//fmt.Println("输出的第",i,"个值：",sArray[i]);
		fmt.Printf("Element[%d] = %d\n",i,sArray[i])
	}
	fmt.Println(sArray[1])
}
type Books struct {
	title string
	author string
	subject string
	book_id int
}
/*
	测试结构体
 */
func testJieGouTi(){
	var book1,book2 Books
	book1.title = "致爱"
	book1.author = "tuhs"
	book1.subject = "科学"
	book1.book_id = 99
	book2 = Books{"星空之门","兔子先生","科幻",12}

	var bookArray [2]Books
	bookArray[0] =  book1
	bookArray[1] =  book2

	fmt.Println(Books{"星空之门","兔子先生","科幻",12})
	fmt.Println(Books{title:"haha",book_id:13})
	testJieGouTi2(book1)
	testJieGouTi3(bookArray)
}
func testJieGouTi2(book Books){
	fmt.Println(book)
}

func testJieGouTi3(bookArray [2]Books){
	var i int
	for i=0;i< len(bookArray);i++{
		fmt.Printf("输出第【%d】本书，详细信息是:%d\n",i,bookArray[i])
	}
}
/*
	切片(动态化数组)
 */
 func testQieMian(){
 	var arr1 []int
	 var arr2 = make([]int,3,5)
	 arr1 = [] int {1,3}
	 arr2 = [] int {7,6,4,2,8,43,56,32,123,13,66,78}
	 fmt.Println(arr2)
	 fmt.Println(arr1)
	 fmt.Printf("arr2数组的长度【%d】,数组的最大容量是【%d】\n",len(arr2),cap(arr2))

	 fmt.Printf("截取arr2[3:5]得到的结果是：%d\n",arr2[3:5])

	 copy(arr1,arr2[3:5])

	 //var k,v int;
	 //循环处理arr1数组
	 for k, v := range arr1 {
		 fmt.Println(k)
		 fmt.Println(v)
	 }
	 //语言范围 返回的是key 和 value
	 for _,v := range arr2[2:7] {
	 	 arr1 = append(arr1,v)
	 }

	 printSlice(arr1)

	 printSlice(arr2)
 }
 //格式化输出
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
//Map 函数的用法
func testMap(){
	//声明map函数 默认是nil
	var map1 map[string]int
	//使用make 函数 创建实例
	map1 = make(map[string] int)
	var map2 = make(map[string] int)
	map1["ddd"] = 222
	map1["aaa"] = 444
	map1["bbb"] = 555
	map1["ccc"] = 111

	map2["ooo"]=889
	map2["ppp"]=443
	map2["uuu"]=776
	map2["yyy"]=552
	for k, v := range map1 {
		fmt.Printf("输出map1数组Key:[%d],输出数组:value[%d]\n",k,v)
	}
	fmt.Println("\n\n")
	for k, v := range map2 {
		fmt.Printf("输出map2数组Key:[%d],输出数组:value[%d]\n",k,v)
	}
	fmt.Println("\n\n")
	delete(map1, "aaa")
	for k, v := range map1 {
		fmt.Printf("输出map1数组Key:[%d],输出数组:value[%d]\n", k, v)
	}
}
//呼叫出租车
type calTxt interface {
	txt()
	tousu()
}
//卡宴车
type KaYan struct {

}
func (kayan KaYan)txt(){
	fmt.Println("我是卡宴车主，欢迎您乘坐。。。")
}
//奔驰
type BenZi struct {

}
func (benzi BenZi)txt(){
	fmt.Println("我是奔驰车主，欢迎您乘坐。。。")
}

//宝马
type BMZ struct {

}
func (bmz BMZ)txt(){
	fmt.Println("我是宝马车主，欢迎您乘坐。。。")
}
type KeFU struct {

}
func (kefu KeFU)tousu(){
	fmt.Println("请问我有什么能帮助您的吗？")
}


func startCalTxt(){


	//var caltxt calTxt;

	kefu := new(KeFU)
	bmz := new(BMZ)
	caltxt := calTxt(bmz,kefu);

	if v, ok := caltxt.(KeFU); ok {
		v.tousu()
	}


}