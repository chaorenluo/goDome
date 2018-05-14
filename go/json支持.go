package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//我们使用两个结构体来演示自定义数据类型的json数据编码和解码
type Response struct {
	Page   int
	Fruits []string
}
type Response1 struct {
	Page   int
	Fruits []string
}

func main() {
	//首先我们看一下将基础数据类型编码为JSON数据
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//这里将切片和字典编码为json数组或对象
	slD := []string{"静静", "密码"}
	slB, _ := json.Marshal(slD)
	fmt.Println(string(slB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//JSON包可以自动地编码自定义数据类型。结果将只包括自定义
	// 类型中的可导出成员的值并且默认情况下，这些成员名称都作为JSON数据的键

	reslD := &Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	reslB, _ := json.Marshal(reslD)
	fmt.Println(string(reslB))

	//我们可以使用tag来自定义编码后的json键的名称

	reslC := &Response1{
		Page:   2,
		Fruits: []string{"1", "2"}}
	resQ, _ := json.Marshal(&reslC)
	fmt.Println(string(resQ))

	//解码为json数据为Go的值
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 我们需要提供一个变量来存储解码后的JSON数据，这里
	// 的`map[string]interface{}`将以Key-Value的方式
	// 保存解码后的数据，Value可以为任意数据类型
	var dat map[string]interface{}

	//解码过程中并检查相关错误

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	// 为了使用解码后map里面的数据，我们需要将Value转换为
	// 它们合适的类型，例如我们将这里的num转换为期望的float64
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的数据需要一些类型转换
	strs := dat["strs"].([]interface{})

	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将JSON解码为自定义数据类型，这有个好处是可以
	// 为我们的程序增加额外的类型安全并且不用再在访问数据的时候
	// 进行类型断言
	str := `{"page":1,"fruits":["apple","罗嘉明"]}`
	res := &Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

	// 上面的例子中，我们使用bytes和strings来进行原始数据和JSON数据
	// 之间的转换，我们也可以直接将JSON编码的数据流写入`os.Writer`
	// 或者是HTTP请求回复数据。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	//自己测试
	fmt.Println("------------------------")
	//编码json
	data := []string{"静静", "明明"}
	das, _ := json.Marshal(data)
	fmt.Println(string(das))

	//编码 map
	maps := map[string]int{"apple": 1, "node": 2}
	mapsd, _ := json.Marshal(maps)
	fmt.Println(string(mapsd))

	//自定义类型json编码
	datar := &Response{Page: 1, Fruits: []string{"静静", "明明"}}
	datars, _ := json.Marshal(datar)
	fmt.Println(string(datars))

	//解码json
	bytd := []byte(`{"num":1.2,"start":["a","b"]}`) //需要解码的json
	var mapds map[string]interface{}                //存放解码后的json
	json.Unmarshal(bytd, &mapds)
	//json数据解码类型转换
	float32d := mapds["num"].(float64)
	strinf := mapds["start"].([]interface{})
	fmt.Println(float32d)
	fmt.Println(strinf[0])

	//json解码为自定义数据类型
	strd := `{"page":1,"fruits":["apple","罗嘉明"]}`

	sds := &Response{}
	json.Unmarshal([]byte(strd), &sds)
	fmt.Println(sds.Fruits[1])
}
