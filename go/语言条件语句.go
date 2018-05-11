package main
import "fmt"

func main(){
	test();
	test1();
	test2();
	test3();

}

//if判断
func test(){
	num :=20;
	if num < 20 {
		fmt.Println(num,"小于20");
	}else{
		fmt.Println(num,"大于20");
	}

}
//switch判断
// const(
// 	a="A"
// 	b="B"
// 	c="C"
// 	d="D"
// )
func test1( ){
	s1:="B";
	switch{
		case s1=="A":
			fmt.Println("成绩优秀");
		case s1=="B":
			fmt.Println("成绩良好");
		case s1=="C":
			fmt.Println("成绩及格");
		case s1=="D":
			fmt.Println("成绩不及格");
	}
}
//Type Switch判断某个interface 中储存的变量类型
func test2(){
	var x interface{}
	switch i:=x.(type){
		case nil:
			fmt.Println("x的类型是%T",i)
		case int:
			fmt.Println("x的类型是int类型")
		case float64:
			fmt.Println("x的类型是float64")
		case bool,string:
			fmt.Println("x的类型是bool或者string")
		default:
			fmt.Println("未知类型")
	}
}
//select是Go的一个控制结构,类似于通信的switch语句
//select随机执行一个科运行的case，如果没有case科运行它将阻塞
//直到case可运行或者有default
func test3(){
	var c1,c2,c3 chan int
	var i1,i2 int
	select{
	case i1=<-c1:
		fmt.Println("received",i1,"from c1\n")
	case c2<- i2:
		fmt.Println("sent",i2,"to c2\n")
	case i3,ok:=(<-c3):
		if ok {
			fmt.Println("received",i3,"from c3\n")
		}else{
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("on communication\n");
	}
}
