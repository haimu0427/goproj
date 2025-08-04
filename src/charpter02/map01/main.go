package main

import (
	"fmt"
)

func main() {
	//1
	var map01 map[string]string
	fmt.Println(map01)
	map01 = make(map[string]string, 10)
	map01["宋江"] = "及时雨"
	map01["卢俊义"] = "玉麒麟"
	map01["吴用"] = "智多星"
	map01["林冲"] = "豹子头"
	fmt.Println(map01)
	//2
	var map02 = make(map[string]string, 5)
	map02["宋江"] = "及时雨"
	//3
	map03 := map[string]string{
		"宋江":  "及时雨",
		"卢俊义": "玉麒麟",
		"吴用":  "智多星",
	}
	fmt.Println(map03)
	//test
	map04 := make(map[string]bool, 5)
	map04["刘子健"] = true
	map04["zcz"] = true
	map04["zwx"] = true
	map04["zwx"] = false //覆盖

	fmt.Println(map04)
	//查询
	if value, ok := map04["zwx"]; ok {
		fmt.Println("zwx的值是：", value)
	} else {
		fmt.Println("zwx不存在")
	}

	//删除的两个方法
	//1
	for key := range map04 {
		delete(map04, key)
	}
	fmt.Println(map04)

	map04 = make(map[string]bool, 5)
	map04["刘子健"] = true
	map04["zcz"] = true
	map04["zwx"] = true
	map04["zwx"] = false //覆盖

	clear(map04) //清空
	fmt.Println(map04)

}
