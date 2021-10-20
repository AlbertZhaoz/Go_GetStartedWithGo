package main

import (
	"fmt"
	"strings"
)

func main() {
	f1 := 1.725649
	fmt.Printf("此是f1的类型%T\n",f1)
	//fmt.Println("xxx is %f",f1) 
	//PrintIn自带换行，但是不支持格式化输出字符串
	//Print直接输出内容，Printf格式化输出字符串，Println会输出内容的结尾添加一个换行符
	fmt.Printf("xxx is %f\n",f1)
	var b1 bool
	fmt.Printf("The value of b1 is %v\n",b1)
	s1 := `
	世情薄,
	人情恶,
	雨送黄昏花易落.
	`
	fmt.Println(s1)
	//分割
	ret := strings.Split(s1,",")
	fmt.Println(ret)
	//字符串长度
	fmt.Println(len(s1))
	//字符串拼接
	name:="理想"
	world:= "创造世界"
	fmt.Println(name+world)
	sprintfWord := fmt.Sprintf("%s%s",name,world)
	fmt.Println(sprintfWord)
	fmt.Printf("s1的数据类型是%T\n",s1)
	//包含
	fmt.Println(strings.Contains(name,"理"))
	//前缀
	fmt.Println(strings.HasPrefix(sprintfWord,"理"))
	//后缀
	fmt.Println(strings.HasSuffix(sprintfWord,"理"))
	//Join：在字符串数组中，用sep连接起来
	fmt.Println(strings.Join(ret,"+"))
}