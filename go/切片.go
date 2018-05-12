package main

import (
	"fmt"
)

func main() {
	test4()
}

//go语言切片就是一个动态数组
//数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

func test() {
	//定义一个未指定大小的数组来定义切片
	//var identifier []type
	//切片不需要指名长度或使用make()函数来创建切片
	//简写为
	//slicel:=make([]type,len)
	//make([]T, length, capacity)

	//初始化切片
	//s := []int{1, 2, 3}
	//直接初始化切片 []表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	//s := arr[:]
	//初始化切片s是数组arr的引用
	// s := arr[startIndex:endIndex]
	// //将arr中从下标startIndex到endIndex-1下的元素创建一个新的切片
	// s := arr[startIndex]
	// //缺省endIndex时表示一直到arr的最后一个元素
	// s := arr[:endIndex]
	// //缺省startIndex是将表示 从arr的第一个元素开始
	// s1 := s[startIndex:endIndex]
	// //通过切片s初始化切片s1
	// s := make([]int, len, cap)

	//len()和cap()函数
	var nums = make([]int, 3, 5)
	//切片的长度
	fmt.Println(len(nums))
	//切片的容量
	fmt.Println(cap(nums))
	fmt.Println(nums)
}

//一个切片在未初始化的时候默认为nil空,长度为0
func test2() {
	var numbers []int
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	if numbers == nil {
		fmt.Println("切片是空的")
	}
}

//切片截取
func test3() {
	//创建切片
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	dayin(a)
	//打印原始切片
	fmt.Println("a==", a)
	//打印自切片从1-4但不包括4
	fmt.Println("a[1:4]==", a[1:4])
	//默认下限为0
	fmt.Println("a[:3]==", a[:3])
	//默认上限为len(s)
	fmt.Println("a[4:]==", a[4:])

	a1 := make([]int, 0, 5)
	dayin(a1)
	//打印子切片索引 0包含到索引2不包含
	a2 := a[:2]
	dayin(a2)
}

//append() copy()函数 copy是拷贝 append()是追加
func test4() {
	var a []int
	dayin(a)
	//允许追加空切片
	a = append(a, 0)
	dayin(a)
	//向切片添加一个元素
	a = append(a, 1)
	dayin(a)
	//同时追加多个元素
	a = append(a, 2, 3, 4, 5, 6, 9, 10, 11, 12)
	dayin(a)
	//创建切片是之前的两倍容量
	a1 := make([]int, len(a), cap(a)*2)
	//拷贝吧a拷贝到a1里面
	copy(a1, a)
	dayin(a1)
}
func dayin(a []int) {
	fmt.Println("切片的大小为%x\n", len(a))
	fmt.Println("切片的最大长度为%x\n", cap(a))
}
