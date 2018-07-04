package main

func main(){



}

/*
	函数

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



*/