package test

import(
	"fmt"
)

func main()  {
	//作用域age只在条件判断语句中生效
	if age:=19;age>18{
		fmt.Println("hello")
	}

	//forr 第一个占位符是索引
	s := "albertzhao"
	for i, v := range s {
		fmt.Printf("%v,%c\n",i,v)
	}

	n:=1
	switch n{
	case 1,4,5,6:
		fmt.Println("Monday")
		//fallthrough执行完当前分支后会执行下一个分支
		fallthrough
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Firday")
	}

	var arrAge [1]int
	numArray:=[3]int{10,2}
	arrAge[0] = 1
	fmt.Println(arrAge[0])
	fmt.Println(numArray)

	//自动推断数组长度
	cityArray := [...]string{"上海","北京"}
	fmt.Println(cityArray)
	fmt.Println(len(cityArray))

	//通过索引来初始化数组
	cityArray2 := [...]string{1:"上海",3:"北京"}
	fmt.Println(cityArray2)
	fmt.Println(len(cityArray2))
}