package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个文件
	file, err := os.Open("F:\\Repo\\GetStartedWithGolang\\Day1124\\InterfacePractise\\main.go")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Println(file)
	//当函数退出时，要及时关闭file
	defer file.Close()

	//创建一个*Reader,带缓冲的,读一部分处理一部分缓冲的好处
	//默认缓冲大小是4096
	reader := bufio.NewReader(file)
}
