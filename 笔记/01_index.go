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



	/*
		Println和Printf区别
		Println：一段一段处理，自动加换行
		Printf：格式化输出，把a放在%d的位置
			%c：字符方式打印
			%d：以整型方式打印

	*/
	fmt.Println("a=",a)

	fmt.Printf("a=%d\n",a)


	/*
	8.多重初始化
	a,b,c := 1,2,3

		两个变量值交换
		i := 10
		j := 20
		i,j = j,i

	9.匿名变量
		_ 匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
		如函数返回多个返回值，可以只取其中1个或多个
		func test()(a,b,c int){
			return 1,2,3
		}
		_,d,e = test()

	10.不同类型的变量定义或常量定义,可自动推倒类型
		var a int
		var b float64
		-------------
		a,b = 10,3,14
		-------------
		var (
			a int
			b float64
		)
		-------------
		var (
			a = 1
			b =1.1
		)

		const i int = 9
		const j float64 = 1.11
		--------------
		const (
			i int = 9
			j float64 = 1.11
		)
		--------------
		const (
			i  = 9
			j  = 1.11
		)


	11.iota枚举
		（1）iota常量自动生成器，每个一行，自动累加1
		（2）iota给常量赋值使用
		（3）iota遇到const，重置为0
			const (
				a = iota   //0
				b = iota   //1
				c = iota   //2
			)

			const d = iota //重置为0
		（4）可以只写一个iota
			const (
				a = iota
				b
				c
			)
		(5)如果是同一行，值都一样
			const (
				i = iota  // 0
				j1,j2,j3 = iota,iota,iota // 1
				k = iota     // 2
			)

	*/


	/*
	12.数据类型
		bool类型，没有初始化，默认值为false
			var a bool = true
			var b =	true
			c := true
		浮点类型，float64比float32存储小数更准确
			var f1 float32
			f1 = 3.14
			f2 := 3.14
		字符类型，byte，单个英文，用单引号''
			1.单引号
			2.往往都只有一个字符，转义字符除外，'\n'
		字符串类型，string
			1.双引号
			2.字符串有1个或多个字符组成
			3.字符串都隐藏一个结束符 '\0'
				str = "a" //由a和\0组成了一个字符串
		复数类型，complex


	13.格式化输出
		%T：操作变量所属类型
		%d：整型格式
		%s：字符串格式
		%c：字符格式
		%f：浮点型格式
		%v：自动匹配格式输出


	14.输入的使用
		阻塞等待用户的输入
		fmt.Scanf("%f",&a)  //别忘了&
		fmt.Scan(&a)

		fmt.Println(a)

	15.类型别名
		type bigint int64  //给int64起一个别名bigint
			var a bigint  相当于 var a int64
		type (
			longint int64
			char byte
			)

	16.其他运算符
		& 	取地址运算符   &a   变量a的地址
		*	取值运算符     *a	   指针变量a所指向内存的值


	17. if 和 { 就是条件，昨括号和if在同一行
		if s == ''{

		}
		也可以支持1个初始化语句，初始化语句和判断条件以分号分隔
		if a := 10;a == 10{

		}

		if a:=10;a == 10 {

		}else if a > 10 {

		}else{

		}

	18.switch
		num :=1
		switch num { 		//switch num:=10;num {
		case 1:
			fmt.Println("1")
			break	//go语言保留了break关键字，跳出switch，不写也行，默认就包含
			fallthrough		//不跳出switch，后面的无条件执行

		case 2:
			fmt.Println("2")
			//break
		case 3:
			fmt.Println("3")
			//break
		default:
			fmt.Println("default")
		}

		case后面可以放条件
		score:=95
		switch {
		case score > 90:
			fmt.Println("优秀")
		case score < 90:
			fmt.Println("良好")
		default:
			fmt.Println("其他")
		}


	19.循环语句：for	，range-迭代器
		for 初始化条件;判断条件;条件变化{
		}

		累加示例
		sum := 0
		for i := 1;i <= 100;i++{
			sum = sum + i
		}

		迭代配合数组或者切片
		迭代会默认返回两个值，一个元素位置，一个元素本身
		str := "abc"
		for i,data := range str{

		}
		for i := range str{	//第二个返回值，默认丢弃，返回元素的位置（下标）

		}
		for i,_ := range str{	//第二个返回值，默认丢弃，返回元素的位置（下标）

		}

	20.跳转语句
		break可用于for,switch,select
		continue仅能用于for

		i:=0
		for {	//for后面不写任何东西，这个循环条件永远为真，死循环
			i++

			if i==5{
				break	//跳出循环，如果嵌套多个循环，跳出最近的那个内循环
				continue	//跳出本次循环，写一次继续
			}

		}

	21.goto--可以用在任何地方，但不能跨函数使用，

	*/





}
