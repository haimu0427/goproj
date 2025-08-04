package main

import (
	"factory/model"
	"fmt"
)

func main() {
	stu := model.NewStudent("张三", 18, 100)
	fmt.Println(stu)
	fmt.Println(stu.GetName())
	fmt.Println(stu.GetAge())
}
