package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//切片声明1
	var stuList []string
	//切片插值
	stuList = append(stuList, "AlbertZhao", "AlbertChen")
	stuList = append(stuList, "AlbertZhao", "AlbertChen")
	//切片遍历，index为索引，可以用_来占位 v是元素，range后是遍历的对象
	for index, v := range stuList {
		fmt.Println(index, v)
	}

	//切片声明2
	teaList := make([]string, 3, 10)
	for _, v := range teaList {
		fmt.Println(v)
	}
	fmt.Println(len(teaList))

	//连接切片
	totalList := append(stuList, teaList...)
	fmt.Println(len(totalList))
	fmt.Println("===========")

	//切片的赋值拷贝,copy(dst,src []Type) 会拷贝源目标的最小长度到目的的最小长度，最小长度由dst决定
	//min(len(dst),len(src))
	copytotalList := make([]string, 5, 10)
	copy(copytotalList, totalList)
	fmt.Println(totalList)
	fmt.Println(copytotalList)
	for _, v := range copytotalList {
		fmt.Println(v)
	}
	//从切片中删除元素 删除切片index元素，a = append(a[:index],a[index+1:]...)
	//删除copytotalList第1个元素AlbertChen
	copytotalList = append(copytotalList[:1], copytotalList[2:]...)
	fmt.Println(copytotalList)

	fmt.Println("【小赵和小陈的对话】")
	fmt.Println("小陈：想听一只小狗的故事嘛？")
	fmt.Println("小赵：我认识？")
	fmt.Println("小陈：你不认识，我昨天在楼下遇到的，是一只小白狗")
	fmt.Println("小赵：像史努比嘛？")

	//排序,排序这种长度不定的数组要先转换为切片再进行排序
	var arrAge = [...]int{3, 7, 5, 6, 9, 5, 4, 1}
	sort.Ints(arrAge[:])
	fmt.Println(arrAge)
	for i := 0; i < len(arrAge); i++ {
		fmt.Println(arrAge[i])
	}

	//strings.Contains
	var cmd = "git commit"
	if strings.Contains(cmd, "commit") {
		fmt.Println("包含-commit")
	} else if strings.Contains(cmd, "-m") {
		fmt.Println("包含-m")
	} else {
		fmt.Println("不包含")
	}

	switch cmd {
	case "git commit -m":
		fmt.Println("相等")
	case "git commit":
		fmt.Println("不相等啊")
		fallthrough
	default:
		fmt.Println("不相等")
	}
}
