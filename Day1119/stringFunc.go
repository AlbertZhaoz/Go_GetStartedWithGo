package main

import (
	"strings"
	"fmt"
	"albert1119/util"
	"math/rand"
	"time"
)

func main()  {
	str := "Hello世界"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n",r[i])
	}

	//替换字符串中
	str2:= "go go go to home"
	tempStr2 := strings.Replace(str2, "go", "golang", 2)
	println(tempStr2)

	//删除字符串首尾空格
	str3 := "  hello albert  "
	tempStr3 := strings.TrimSpace(str3)
	//%q可以用双引号将字符串引起来
	fmt.Printf("I'm tempStr3:%q\n", tempStr3)

	//判断字符串是否以指定开头或结尾
	str4:= "This is free city"
	fmt.Println(strings.HasPrefix(str4, "This"))
	fmt.Println(strings.HasSuffix(str4, "city"))

	f2:= utils.Cal(2,3)  //5
	println(f2(4)) //5+4=9

	GuessInt()


}

func GuessInt(){
	rand.Seed(time.Now().Unix())
	x:= rand.Intn(100)
	println(x)
	println("您总共有10次机会，请猜测数：")
	var y int	
	var flag bool
	for i := 9; i >-1; i-- {
		fmt.Scan(&y)
		if y<x{
			fmt.Printf("啊呀，你猜小了,您还有%v机会，请再次输入数字：\n", i)
		}else if y>x{
			fmt.Printf("啊呀，你猜大了,您还有%v机会，请再次输入数字：\n", i)
		}else{
			println("恭喜你猜对了")
			flag = true
			break;
		}
	}
	if !flag{
		fmt.Printf("Sorry，机会用完了，这个数是%v", x)		
	}
	
}