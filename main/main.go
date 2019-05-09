package main

import "fmt"

func main(){
	fmt.Println("Hello Word");
	var str = "Hello Word"
	fmt.Println(str);
	var a,b = 12,13;
	if(a>b){
		fmt.Println("本次比较值最大的是：",a)
	}else{
		fmt.Println("本次比较值最大的是：",b)
	}
	switch a {
	case 10:fmt.Println("本次输出值：10");
	case 12:fmt.Println("本次输出值：13");
	default:
		fmt.Println("本次没有匹配的值");
	}

	for(a<=15){
		fmt.Println("本次循环变量值为：",a);
		a++
	}


}