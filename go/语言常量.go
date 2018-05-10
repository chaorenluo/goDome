package main
import "fmt"

//常量当做枚举使用
// const(
// 	num1=0
// 	num2=2
// 	num3=3
// )
//iota相当于sql里面的自动增长组件
const(
	a=iota//0
	b//1
	c//2
	d="ha" //独立的值 iota+1
	e //ha如果上面的值是独立的值的话下面所有的值就和上面的值一样 同时iota+1 直到遇到iota赋值才会显示iota的累加值
	f=100 //iota+1
	g //iota+1
	h=iota
)
//常量的例子
func main(){
	const WIDTH int=10;
	const HEIGHT int=20;
	const a1,b1,c1=1,false,"多重赋值";
	fmt.Println(WIDTH*HEIGHT);
	fmt.Println(a1,b1,c1);
	//fmt.Println(num2);
	fmt.Println(a,b,c,d,e,f,g,h);
}