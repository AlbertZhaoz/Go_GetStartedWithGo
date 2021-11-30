package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	ReadFileOne("F:\\Repo\\GetStartedWithGolang\\Day1124\\InterfacePractise\\main.go")
	println("==========================")
	//打开一个文件
	file, err := os.Open("F:\\Repo\\GetStartedWithGolang\\Day1124\\InterfacePractise\\main.go")

	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Println(file)
	//当函数退出时，要及时关闭file，否则会有内存泄漏
	defer file.Close()

	//创建一个*Reader,带缓冲的,读一部分处理一部分缓冲的好处,可以处理比较大的文件
	//默认缓冲大小是4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //每次读取一行，读到一个换行就结束
		//io.EOF表示文件末尾
		if err == io.EOF {
			fmt.Println("已经读到文件末尾啦,我跳出去啦")
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束....")
}

//小文件读取
func ReadFileOne(path string) {
	//一次性读取
	content, err := ioutil.ReadFile(path) //content []byte
	if err != nil {
		fmt.Println("Read file err=", err)
	}

	fmt.Println(string(content)) //[]byte转string   string转[]byte  []byte(string)
}

//写文件
