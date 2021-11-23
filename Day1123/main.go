package main

import(
	"fmt"
	"encoding/json"
)

//定义一个Cat结构体
//打Tag进行小写的转换
type Cat struct{
	Name string `json:"nAme"`
	SubName string `json:"subname"`
	Age int
	Color string `json:"color"`
}

type Point struct{
	x,y int
}

type Rect struct{
	leftUp,rightDown Point
}

func main()  {
	r1:= Rect{Point{1,2},Point{3,4}}
	fmt.Printf("r1.leftUp.x:%p,r1.leftup.y:%p,r1.rightDown.x:%p,r1.rightDown.y:%p\n", &r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)

	//匿名结构体
	var user struct{Name string;Age int}
	user.Name = "AlbertZhao"
	user.Age = 25
	println(user.Name)
	//创建Cat变量
	var catOne Cat;
	catOne.Name = "Xiaochen"
	catOne.SubName = "XiaochenMiao"
	catOne.Color = "White"
	catOne.Age = 2
	fmt.Println(catOne)

	data,err := json.Marshal(catOne)
	if err!=nil{
		defer func(){
			err2 := recover()
			if err2!=nil{
				fmt.Println(err2)
				//这里可以将错误信息发送给管理员
			}
		}()
		fmt.Println("json encoding err  ",err)		
	}
	fmt.Println("Json后的数据",string(data))




}