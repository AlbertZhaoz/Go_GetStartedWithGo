package main

import "fmt"

func main()  {
	stuNameSlice := []string{"albertzhao","albertchen","albertli"}
	for _, v := range stuNameSlice {
		fmt.Println(v)
	}
	stuNameSlice = append(stuNameSlice,"albertzhang")
	for _, v := range stuNameSlice {
		fmt.Println(v)
	}

	fmt.Println("==============")

	//这里的append的原理是从[0-1) [2-End)两个拼起来
	//删除第一个元素 append(slice[:1],slice[2:]...)注意这里有三个点
	stuNameSlice = append(stuNameSlice[:1],stuNameSlice[2:]...)
	fmt.Println(stuNameSlice)
	
}