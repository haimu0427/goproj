package admin

import (
	"fmt"
)

var AdminPassword string = "admin123"

func AdminPanel() {
	fmt.Println("--------------管理员操作界面--------------")
	fmt.Println("1. 查看用户信息")
	fmt.Println("2. 修改用户信息")
	fmt.Println("3. 删除用户")
	fmt.Println("4. 返回上一级")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("查看用户信息")
		// 调用查看用户信息的函数
	case 2:
		fmt.Println("修改用户信息")
		// 调用修改用户信息的函数
	case 3:
		fmt.Println("删除用户")
		// 调用删除用户的函数
	case 4:
		return
	default:
		fmt.Println("输入错误，请重新输入")
	}
}
