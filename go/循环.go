package main

import "fmt"

func main() {
	test1()
}
func test() {
	var a, b int
	a = 1
A:
	for a < 100 {
		a++
		for b = 2; b < a; b++ {
			if a%b == 0 {
				goto A
			}
		}
		fmt.Println(b)
	}
}

//goto语句无条件转移到过程中指定的行
func test1() {
	a := 1

A:
	for a < 20 {
		if a == 15 {
			goto A

		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}

}
