package main

import "fmt"

func main() {
	const(
		n1 = iota //实现类似枚举的效果
		n2
		n3
		n4
	)
	const(
		a1 = iota //0
		a2 = 100
		a3        //2,上面新增了一行，计数加1，此处又加1为2
		a4
	)
	const(
		b1,b2 = iota+1,iota+2 //此处iota都是0,没有新增一行,b1=1,b2=2
		b3,b4 = iota+1,iota+2 //b3=2,b4=3
	)
	

	fmt.Println("我是iota累计四次的数:",n4)

	fmt.Println("albert")
	var sl string = "Hello albert"
	fmt.Println(sl)
	DeclareVariable()
}

func DeclareVariable()  {
	//批量声明变量
	var(
		name string = "Albert"
		age int = 25
	)
	//简短变量声明，只能在函数中使用
	author:="AlbertZhao"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("The name of author is %s",name)
	fmt.Println("The name of author is %s",author)
	x,_:=Foo()
	fmt.Println(x)
}

func Foo()(int,string)  {
	return 25,"ZackYang"
}
