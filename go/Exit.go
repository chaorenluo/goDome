//使用ox.Exit立即退出程序并给状态值

package main

import (
	"fmt"
	"os"
)

func main() {

	//当使用os.Exit()退出程序的时候 defer清尾函数不会执行
	defer test()

	//退出程序并设置状态值
	os.Exit(3)
}
func test() {
	fmt.Println("清尾方法")
}
