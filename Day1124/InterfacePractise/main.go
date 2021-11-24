package main

import(
	"sort"
	"fmt"
)

//声明Hero结构体
type Hero struct{
	Name string
	Age int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现接口Interface
func (hs HeroSlice)Len() int{
	return len(hs)
}

func (hs HeroSlice)Less(i,j int) bool{
	return hs[i].Age <hs[j].Age
}

func (hs HeroSlice)Swap(i,j int){
	//hs[i],hs[j] = hs[j]:hs[i]这句话等价下面的
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main()  {
	//先定义一个数组/切片
	var intSlice = []int{0,-1,10,7,90}

	//使用系统提供的排序方案
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}