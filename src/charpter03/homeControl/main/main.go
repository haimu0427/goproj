package main

import (
	"fmt"
	"homeControl/admin"
	"homeControl/familyFinance"
)

func main() {
	//初始化两个用户
	admin.InitUsers()

	fmt.Println("欢迎使用家庭收支记录软件")
	fmt.Println("请先登录")
	if LoginLoop() != nil {
		Panel(admin.Users[0], admin.Users[1])
	}
}

// Panel 显示家庭收支记录软件的主界面
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
func LoginLoop() *familyFinance.FamilyFinance {
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
				// // 同步用户列表到admin包
				// admin.Initadmin.Users(admin.Users)
				// admin.AdminPanel()
				// // 从admin包获取更新后的用户列表
				// admin.Users = admin.Getadmin.Users()
			}
		case 1:
			for {
				fmt.Println("请输入用户名:")
				fmt.Println("请输入密码:")
				var username, password string
				fmt.Scanln(&username)
				fmt.Scanln(&password)
				if (username == admin.Users[0].User.Name && password == admin.Users[0].User.Password) ||
					(username == admin.Users[1].User.Name && password == admin.Users[1].User.Password) {
					fmt.Println("登录成功")
					return admin.Users[0]
				} else {
					fmt.Println("用户名或密码错误，请重新输入")
					return nil
				}
			}

		case 2:
			admin.Users = append(admin.Users, familyFinance.CreateUser())
		case 3:
			return nil
		}
	}
}
