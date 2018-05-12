//结构体 type设定结构体的名称 struct语句定义一个新的数据类型
package main

import (
	"fmt"
)

type book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 book //声明Book1为book类型
	var Book2 book //声明book2为book类型
	Book1.title = "Go语言"
	Book1.author = "www.runoob.com"
	Book1.book_id = 694564
	Book1.subject = "Go语言教程"

	Book2.title = "python语言"
	Book2.author = "www.runoob.com"
	Book2.book_id = 694564
	Book2.subject = "Go语言教程"

	fmt.Println("Book 1 title%s\n", Book1.title)
	fmt.Println("Book 2 title%s\n", Book2.title)
	test(Book1)
	test1(&Book2)
	fmt.Println("Book 1 title%s\n", Book1.title)
}

//结构体作为参数传递
func test(parms book) {
	parms.title = "java语言"
	parms.author = "www.runoob.com"
	parms.book_id = 694564
	parms.subject = "java语言教程"
	fmt.Println("Book 2 title%s\n", parms.title)
}

//你可以定义指向结构体的指针类似于其他指针变量，
func test1(book *book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
