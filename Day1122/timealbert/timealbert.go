package timealbert

import(
	"fmt"
	"time"
	"strconv"
)

func ShowTime()  {
	//打印当前时间
	now := time.Now()
	//year := now.Year()
	fmt.Printf("now is %v, now type is %T\n", now,now)
	//println(year)
	fmt.Printf(now.Format("2006-01-02"))

	//休眠打印
	for i := 0; i < 2; i++ {
		println(i)
		time.Sleep(time.Second)
	}

	println("===============")
	//测试运行时间
	var str string;
	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	println(end-start)
}