package main

import(
	"strconv"
)

type Student struct{
	name string
	age int
	subname string
}

//继承
type SmallStudent struct{
	Student //继承
	score float64
}

func (smallStu *SmallStudent)PrintScore(){
	println(smallStu.score)
	println(smallStu.name)
}

//结构体方法
func (stu *Student)PrintStuInfo(){
	println(stu.name+stu.subname+strconv.Itoa(stu.age))
}

func main()  {
	var stuOne Student
	stuOne.name = "AlbertZhao"
	stuOne.subname = "Albert"
	stuOne.age = 25
	stuOne.PrintStuInfo()
	var smallStu SmallStudent
	smallStu.name = "Jack"
	smallStu.age = 22
	smallStu.subname ="Tom"
	smallStu.score = 85
	smallStu.PrintScore()
	smallStu.PrintStuInfo()

	var mod Usb;
	//手机实现了接口
	phone := Phone{}
	mod = phone
	mod.Start()
	mod.Stop()
	//相机实现了接口
	camera := Camera{}
	mod = camera
	mod.Start()
	mod.Stop()
}

type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}

type Camera struct{

}

func (phone Phone)Start(){
	println("手机开始工作")
}

func (phone Phone)Stop(){
	println("手机停止工作")
}

func (camera Camera)Start(){
	println("相机开始工作")
}

func (camera Camera)Stop(){
	println("相机停止工作")
}