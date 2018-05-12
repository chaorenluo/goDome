package main

import "fmt"

func main() {
	// var a, b int
	// a = 100
	// b = 200
	// var c, d string
	// ret := max(a, b)
	// fmt.Println(ret)
	// c, d = swap("罗佳明", "静静")
	// fmt.Println(c, d)
	// fmt.Println("交换之前的值%d\n", a)
	// fmt.Println("交换之前的值%d\n", b)
	// //&a指向a指针,a变量的内存地址
	// //&b指向b指针,b变量的内存地址
	// swap(&a, &b)
	// fmt.Println("交换之后的值%d\n", a)
	// fmt.Println("交换之后的值%d\n", b)
	//函数作为值
	// fc := func(x int) int {
	// 	fmt.Println("打野")
	// 	return x
	// }
	// fmt.Println(fc(10))

	//匿名函数

	// gets := getstring()
	// fmt.Println(gets())
	// fmt.Println(gets())
	// fmt.Println(gets())
	// gets1 := getstring()
	// fmt.Println(gets1())

	//传递参数的匿名函数
	gets := getstring(10, 50)
	fmt.Println(gets())

}

//一个标准的函数
// func function_name([paramter list]) [return_types]{
// 	//函数体
// }

//函数返回两个数的最大值
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//返回多个值
// func swap(x, y string) (string, string) {
// 	return x, y
// }

//函数值传递值即将实际参数复制一份传递到函数中,这样在函数中如果对参数进行修改不会影响

// func swap(x, y int) int {
// 	var tem int
// 	tem = x
// 	x = y
// 	y = tem
// 	return tem
// }

//函数引用传递值在调用函数传递参数的时候实际是吧参数的内存地址传递都到函数中
//那么函数对参数进行修改将会影响实际参数
// func swap(x *int, y *int) {
// 	var temp int
// 	temp = *x
// 	*x = *y
// 	*y = temp
// }

//匿名函数 匿名函数可以直接使用函数内的变量
// func getstring() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

//带参数的闭包方法
func getstring(x int, y int) func() (int, int) {
	a, b := x, y
	return func() (int, int) {
		a = a * 10
		b = b * 100
		return a, b
	}
}
