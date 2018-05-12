package main

import (
	"fmt"
)

func main() {
	var parms = []int{10, 20, 30, 40, 50}
	//num := test3(parms, 5)
	fmt.Println(test3(parms, 5))
}

//go语言数组声明
//var variable_name [size] variable_type
// func test() {
// 	//var balance [10]int
// 	//初始化数组
// 	//var balances = [5]int{1, 2, 3, 4, 5}
// 	var tds [10]int
// 	for a := 0; a < 10; a++ {
// 		tds[a] = 100 + a
// 	}
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(tds[i])
// 	}
// }

//二维数组
// func test2() {
// 	//var num = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

// 	//这种写法会报错
// 	// var num = [2][3]int{
// 	// 	{1, 2, 3},
// 	// 	{4, 5, 6}
// 	// }

// 	//这种写法正确的
// 	var num = [2][3]int{{1, 2, 3}, {4, 5, 6}}
// 	for i := 0; i < 2; i++ {
// 		for a := 0; a < 3; a++ {
// 			fmt.Println(num[i][a])
// 		}
// 	}
// }

//数组参数
func test3(arr []int, size int) float32 {

	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg

}
