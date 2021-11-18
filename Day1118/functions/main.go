package main

import(
	"fmt"
)

func main()  {
	fmt.Println(sum1(6,7))
	fmt.Println(sum2(6,6,6))
	fmt.Println(sum3(6,6))
	sum4(6,6,5,6,10)
	fmt.Println(sumsub(1,2))
	printInfo()
	fmt.Println(f1())
	al := x3(x1,x2)
	al()

	f:= AddUpper()
	println(f(1))
	println(f(1))
}

func sum1(x int,y int) int  {
	return x+y
}

//命名返回值就相当于在函数中声明了一个变量 ref = x+y
func sum2(x int,y int,z int)(ref int)  {
	ref =  x+y+z
	return //使用命名返回值可以省略return ref中的ref
}

//参数的类型简写
func sum3(x,y int)int{
	return x+y
}

//可变长参数 其实这边传入的是切片x []int 
//重载实现
func sum4(x ...int){
	fmt.Printf("我是可变长参数%v\n",x)
}

//多返回值
func sumsub(x,y int)(int,int){
	return x+y,x-y
}

//defer延迟执行，将defer后面的语句延迟到函数即将返回的时候再执行
//多个defer是栈操作，压栈和出栈
//defer多用于函数结束之前释放资源
func printInfo(){
	fmt.Println("aaa")
	defer fmt.Println("bbbb")
	fmt.Println("cccc")
}

func f1()int{
	x:=5
	defer func(){
		x++
		fmt.Printf("函数执行了，此时X的值为%v\n",x)
	}()   //自执行函数
	return x
}

//函数类型与变量  联想C#中的委托
func x1(){
	println("Hello")
}

func x2()int{
	return 557
}

func x5()int{
	println("我是x5")
	return 10086
}

func x3(y func(),z func()int) func()int{
	y()
	z()
	defer func(){
		println("我是匿名函数")
	}()
	func(){
		println("我也是匿名函数")
	}()
	return x5
}

//累加器
func AddUpper() func(int)int{
	var n int = 10
	return func(x int)int {
		n = n+x
		return n
	}
}
