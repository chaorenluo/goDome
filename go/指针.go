package main

import (
	"fmt"
)

func main() {
	//	&放到变量前相当于返回一个变量的内存地址
	// var a int
	// a = 10
	// fmt.Println("a的变量地址为%x\n", &a)
	// var a int = 20 //申明实际变量
	// var ip *int    //申明指针变量
	// ip = &a        //指针变量的储存地址
	// fmt.Println("ip变量的地址是%x\n", &a)
	// //指针变量的储存地址
	// fmt.Println("ip变量储存的指针地址%x\n", ip)
	// //使用指针访问
	// fmt.Println("ip变量值%x\n", *ip)
	// var st *string
	// var s1 string = "罗嘉明"
	// st = &s1
	// if st == nil {
	// 	fmt.Println("st是空指针")
	// } else {
	// 	fmt.Println("st不是空指针")
	// }
	// fmt.Println(st)
	test1()

}

//指针数组
const MAX int = 5

func test() {
	var num = [5]int{1, 2, 3, 4, 5}
	var prt [MAX]*int
	for i := 0; i < 5; i++ {
		prt[i] = &num[i]
	}
	for x := 0; x < 5; x++ {
		fmt.Println(*prt[x])
	}
}

//指向指针的指针
func test1() {
	var a int = 10
	var b *int
	//申明指针指向指针的变量时需要使用**
	var c **int
	var d ***int
	b = &a
	c = &b
	d = &c
	fmt.Println("变量a的值%x\n", a)
	fmt.Println("指针b的值%x\n", *b)
	fmt.Println("指针c的值%x\n", **c)
	fmt.Println("指针d的值%x\n", ***d)
}
