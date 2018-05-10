package main

import "fmt"

//主入口函数
//在1.9版本中无需定义int以及float32，float64系统自动识别
func main(){
	var a=10;
	var b=10.5;
	fmt.Println(a,b);//整形加上浮点型会抱错
	test();

}
var bools bool;//全局声明；
var b1,b2=true,false;//不要类型的声明
func test(){
	fmt.Println(bools);
}