package main

import "fmt"

func main() {
	//关于Go中，指针只起到取地址和取值两个功能，前者&，后者*
	n := 18
	fmt.Println(&n)
	fmt.Println(*(&n))

	//Go中空指针赋值问题
	//var a *int //没有初始化，对应的地址为nil，*a毫无意义，如何避免new关键字
	//*a = 100
	//fmt.Println(*a)

	//new函数申请一个内存地址
	var a = new(int)
	*a = 100
	fmt.Println(*a)

	//make也是用于内存分配，区别于new，只用于slice、map以及chan的内存创建，而且它返回的类型就是
	//这三个类型本身，而不是他们的指针类型，因为这三种类型是引用类型，所以没必要返回他们的指针。
	//make函数是无法替代的，在初始化那三个类型，然后才可以对其进行操作。
	sourceSlice := make([]string, 5, 10)
	sourceSlice[0] = "ClaireLi"
	fmt.Println(sourceSlice)

	sourceMap := make(map[string]int, 10)
	sourceMap["ClaireLi"] = 10
	fmt.Println(sourceMap)

	//声明map字典
	scoreMap := make(map[string]int, 10) //最好预估一个容量，避免在程序运行期间再动态扩容
	scoreMap["小明"] = 100
	scoreMap["张三"] = 99
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])

	//map的初始化,声明时填充元素
	userInfo := map[string]string{"username": "AlbertZhao", "password": "albertzhao"}
	fmt.Println(userInfo)
	fmt.Println(userInfo["username"])

	//判断键值是否存在
	value, ok := userInfo["username"] //ok是golang中约定的
	if ok {
		fmt.Println("键值存在")
		fmt.Println(value)
	}

	//map的遍历
	for key, value := range userInfo {
		fmt.Printf("键为%v,值为%v\n", key, value)
	}
	//只取map中的key值
	for key := range userInfo {
		fmt.Printf("键为%v\n", key)
	}
	//删除map中的元素 delete(map,key)
	delete(userInfo, "username")
	fmt.Println(userInfo)
	//按照指定顺序遍历map
}
