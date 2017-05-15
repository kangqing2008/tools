package main

import "fmt"

type Data struct{
	Id		int64
	Name	string
	Age		int8
}

func main() {
	var object interface{}
	var data1,data2 *Data
	data1 = &Data{Id:1,Name:"kangqing",Age:34}
	//data2 = Data{Id:2,Name:"mengminghui",Age:33}
	object = data1
	data2  = object.(Data)
	fmt.Println(data1,data2)
	data2.Age = 121
	fmt.Println(data1,data2)
}
