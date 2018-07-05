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

	7.有参数也有返回值





*/