package main

import "fmt"

func main() {

	/*
	变量
	1.声明格式
		var 变量名 类型 ---如果变量声明了，必须要使用
	2.只是声明没有初始化的变量，默认值是0
	3.同一个{}里，声明的变量名是唯一的
	4.可同时声明多个变量
	var a,b,c int
	5.变量赋值
	a = 5
	6.变量初始化同时赋值
	var b int = 8
	7.自动推倒类型，只能使用一次用于初始化
		变量名 := 值


	*/

	var a int
	fmt.Println(a)
	b := 8
	fmt.Println(b)

}
