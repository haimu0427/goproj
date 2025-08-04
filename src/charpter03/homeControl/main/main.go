package main

import (
	"fmt"
	"homeControl/admin"
	"homeControl/familyFinance"
)

var users []*familyFinance.FamilyFinance

func main() {
	//初始化两个用户
	users = append(users, familyFinance.SignUp("张三", "123456", 1000))
	users = append(users, familyFinance.SignUp("李四", "654321", 2000))
	fmt.Println("欢迎使用家庭收支记录软件")
	fmt.Println("请先登录")
	if LoginLoop() {
		Panel(users[0], users[1])
	}
}
func Panel(f1, f2 *familyFinance.FamilyFinance) {
	for {
		var keyboard int
		fmt.Println("--------------家庭收支记录软件--------------")
		fmt.Println("              1. 收支明细")
		fmt.Println("              2. 登记收入")
		fmt.Println("              3. 登记支出")
		fmt.Println("              4. 转账功能")
		fmt.Println("              5. 退出软件")
		fmt.Println(" 请选择你的操作(1-5):")
		fmt.Scanln(&keyboard)
		switch keyboard {
		case 1:
			fmt.Println("你选择了收支明细")
			f1.ShowDetials()
		case 2:
			fmt.Println("你选择了登记收入")
			f1.ShowInDetials()
		case 3:
			fmt.Println("你选择了登记支出")
			f1.ShowOutDetials()
		case 4:
			fmt.Println("开始转账功能")
			familyFinance.Transfer(f1, f2)
		case 5:
			fmt.Println("你选择了退出软件")
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
func LoginLoop() bool {
	for {
		fmt.Println("--------------登录界面--------------")
		fmt.Println("0. 管理员操作")
		fmt.Println("1. sign in")
		fmt.Println("2. sign up")
		fmt.Println("3. exit")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			fmt.Println("管理员操作")
			fmt.Println("请输入管理员密码:")
			var adminPassword string
			fmt.Scanln(&adminPassword)
			if adminPassword == admin.AdminPassword {
				fmt.Println("管理员登录成功")
				admin.AdminPanel()
			}
		case 1:
			for {
				fmt.Println("请输入用户名:")
				fmt.Println("请输入密码:")
				var username, password string
				fmt.Scanln(&username)
				fmt.Scanln(&password)
				if (username == users[0].Name && password == users[0].Password) ||
					(username == users[1].Name && password == users[1].Password) {
					fmt.Println("登录成功")
					return true
				} else {
					fmt.Println("用户名或密码错误，请重新输入")
					return false
				}
			}

		case 2:
			for {
				var username, password, password2 string
				fmt.Println("请输入用户名:")
				fmt.Scanln(&username)
				fmt.Println("请输入密码:")
				fmt.Scanln(&password)
				fmt.Println("请再次输入密码:")
				fmt.Scanln(&password2)
				if password == password2 {
					fmt.Println("需要存钱吗? 需要请按0, 不需要请按1")
					var balance int
					var choice int
					fmt.Scanln(&choice)
				loop:
					switch choice {
					case 0:
						fmt.Println("请输入存入金额:")
					case 1:
						balance = 0
					default:
						fmt.Println("输入错误，请重新输入")
						goto loop
					}
					fmt.Scanln(&balance)
					users = append(users, familyFinance.SignUp(username, password, balance))
					fmt.Println("注册成功")

					return true
				}
			}
		case 3:
			return false
		}
	}
}
