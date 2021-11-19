package main

import(
	"strings"
)

func main(){
	f1:= makeSuffix(".jpg") //无需每次都传入.jpg
	println(f1("Albert"))
	println(f1("Albert"))
}


//闭包的最佳实践：
//此函数可以接收一个文件后缀名（比如.jpg),并返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg)则返回文件名.jpg，
//如果已经由.jpg后缀，则返回原文件名。
//strings.HasSuffix()：可以判断某个字符串是否有指定的后缀
func makeSuffix(suffix string) func(string)string{
	return func(name string)string{
		if strings.HasSuffix(name,".jpg"){
			return name
		}else{
			return name+suffix
		}
	}
}