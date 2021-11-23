package timealbert

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func ShowTime() {
	//打印当前时间
	now := time.Now()
	//year := now.Year()
	fmt.Printf("now is %v, now type is %T\n", now, now)
	//println(year)
	fmt.Printf(now.Format("2006-01-02"))

	//休眠打印
	for i := 0; i < 2; i++ {
		println(i)
		time.Sleep(time.Second)
	}

	println("===============")
	//测试运行时间
	var str string
	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	println(end - start)

	err := test2("d", "dga", "dga", "dga", "dga", "dga", "dga", "dga", "dga", "dga", "dga")
	if err == nil {
		println("aaaaa")
		panic(err)
	}
	println("xxx")
}

func test2(args ...string) (err error) {
	//使用defer+recover来捕获和处理异常
	defer func() {
		err := recover() //recover内置函数，可以捕获到异常
		if err != nil {
			println("错误是：")
			println(err)
		}
	}()
	if args[100] == "" {
		return errors.New("字符串错误")
	} else {
		return nil
	}

}
