package main
import "fmt"

//主入口函数首字母小写为私有大写为公有
var a="菜鸟教程"
var b string="www.baidu.com"
var c bool=true
//类型相同的变量,非全局变量
var vname1, vname2, vname3=1,2,3
func main (){
	//声明一个局部变量不适用会抱错
	g,h:="love","No1";//这种声明方式只能在函数体内 并且变量名没被声明过

	fmt.Println(a,b,c);
	fmt.Println(g,h);
}
