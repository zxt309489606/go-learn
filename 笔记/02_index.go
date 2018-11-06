package main

func main(){



}

/*
	函数
	函数名称：根据约定，首字母小写为private，大写为public
	参数列表：函数可以是0个或多个参数，参数格式为：变量名 类型，如果有多个参数通过逗号分隔，不支持默认参数
	返回类型：
			1.如果有多个返回值，变量名不是必须的，可以只有类型
			2.如果只有一个返回值且不声明返回值变量，那么可以省略，包括返回值的括号，如 func myfunc() int{}
			3.如果没有返回值，那么直接省略最后的返回信息
			4.如果有返回值，那么必须在函数的内部添加return语句

	0.定义格式
	func FuncName (参数列表) (返回类型){
		//函数体
		return v1,v2
	}

	2.无参无返回值函数的使用：函数名()
	func MyFunc(){
		a := 6
		fmt.Println(a)
	}

	调用：MyFunc()

	3.有参无返回值
	func MyFunc(a int){
		fmt.Println(a)
	}

	func Myfunc(a int,b int){
	}
	func Myfunc(a,b int){
	}
	func Myfunc(a int,b string){
	}

	调用：函数名(所需参数)
	MyFunc(666)

	4.不定参数列表
	...type 不定参数类型，传递的参数可是0个或多个
	func Myfunc(args ...int){
		len(args)	//获取用户传递参数的个数

		//args可以使用for和range迭代
	}

	func Myfunc (a int,args ...int){
		//调动时，固定参数一定要传参，不定参数根据需求传递
		//并且不定参数，一定放在形参中的最后一个参数

	}

	不定参数传递
	func Myfunc1(tmp ...int){

	}
	func Myfunc2(args ...int){
		//全部元素传递给Myfunc1
		Myfunc1(args...)
		//只把后2个参数传递给另外一个函数使用
		Myfunc1(args[:2]...) //0~2（不包括2），传递过去
		Myfunc1(args[2:]...) //从args[2]开始，包括本身，把后面所有元素传递过去

	}
	5.无参有返回值，只有1个返回值
		有返回值的需要return中断函数，通过retuan返回
		func Myfunc(){
			return 666
		}
		给返回值起一个变量名，go推荐的写法，常用写法
		func Myfunc() (result int){
			result = 666
			return
		}

	6.函数有多个返回值
		//
		func Myfunc() (int,int,int){
			return 1,3,4
		}
		调用：a,b,c := Myfunc()
		//go推荐写法
		func Myfunc() (a int,b int,c int){
			a,b,c = 11,22,33
			return
		}
		调用：a,b,c := Myfunc()


	7.函数类型
		函数也是一种数据类型，通过type给一个函数类型起名

		FuncType是一个函数类型,
		type FuncType func(int,int) int //没有函数名，没有{}
		声明一个函数类型变量，变量名叫fTest
		var fTest FuncType
		fTest = Add
		result = fTest(1,2)
		Ftest = Minus
		result = fTest(10,5)

	8.回调函数，函数参数是函数类型，这个函数就是回调函数


	9.匿名函数与闭包，闭包以引用方式捕获外部变量，就说里面变量改了，外面也会改

		定义匿名函数，同时调用，后面的括号代表外部调用
		func(){
			fmt.Println()
		}()

		f := func(i,j int){

		}
		f(1,2)

		func(i,j int){
			fmt.Print()
		}(10,20)

		有参有返回值
		x,y := func(i,j int) (max,min int){
			if i > j{
				max = i
				min = j
			}else{
				max = j
				min = i
			}
			return
		}(10,20)


	10.defer
		用于延迟一个函数或方法的执行，只能出现在函数或方法的内部
		函数结束前调用
		如果有多个defer，他们会以LIFO(后进先出)的顺序执行（就是反向顺序	），不管函数发生什么错误，defer都会执行

	11.局部变量
		执行到定义变量那句话，才开始分配空间，离开作用域自动释放
	   全局变量
		定义在函数外部，全局变量在任何地方都能使用
       不同作用域，允许定义同名变量，使用变量的原则：就近原则

	12.导入包
		import "fmt"
		import "os"

		常用方式
		import (
			"fmt"
			"os"
		)

		.操作
		import . "fmt" //调用函数，无需通过包名，不常用

		给包名起别名
		import new_fmt "fmt"

		忽略此包
		import _ "fmt"
		_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数

	13.工程管理
		同目录
		（1）分文件编程(多个源文件)，必须放在src目录
		（2）设置GOPATH环境变量
		（3）同一个目录包名必须一样
		（4）go env查看go的相关环境路径
		（5）同一个目录，调用别的文件的函数，直接调用即可，无需包名引用
		不同目录


		目录结构
		src：放源代码，
		bin：放可执行程序
		pkg：放平台相关的库
		如果有多个文件多个包，
		（1）配置gopath环境变量，配置src的同级目录的绝对路径
		（2）自动生成bin或pkg目录，需要使用go install 命令，除了要配置gopath环境变量，还要配置gobin环境变量








*/