package admin

import (
	"fmt"
	"homeControl/familyFinance"
)

var AdminPassword string = "admin123"

// 全局用户列表，供admin包使用
var Users []*familyFinance.FamilyFinance

func AdminPanel() {
	for {
		fmt.Println("--------------管理员操作界面--------------")
		fmt.Println("1. 创建用户")
		fmt.Println("2. 查看用户信息")
		fmt.Println("3. 修改用户信息")
		fmt.Println("4. 删除用户")
		fmt.Println("5. 返回上一级")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("创建用户")
			Users = append(Users, familyFinance.CreateUser()) //这里的几个用户用的不是同一个, 需要修改
		case 2:
			showUsers()
		case 3:
			fmt.Println("修改用户信息")
			var index int
			var name, password string
			var balance int
			fmt.Print("请输入用户序号: ")
			fmt.Scanln(&index)
			fmt.Print("请输入新用户名: ")
			fmt.Scanln(&name)
			fmt.Print("请输入新密码: ")
			fmt.Scanln(&password)
			fmt.Print("请输入新余额: ")
			fmt.Scanln(&balance)
			ModifyUser(index-1, name, password, balance)
		case 4:
			fmt.Println("删除用户")
			// 调用删除用户的函数
		case 5:
			fmt.Println("返回上一级")
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}

// InitUsers 初始化用户列表（供main包调用）
func InitUsers() {
	Users = append(Users, familyFinance.SignUp("张三", "123456", 1000))
	Users = append(Users, familyFinance.SignUp("李四", "654321", 2000))
}

func GetUsers() []*familyFinance.FamilyFinance {
	return Users
}
func showUsers() {
	fmt.Println("查看用户信息")
	fmt.Println("--------------用户信息列表--------------")
	if len(Users) == 0 {
		fmt.Println("暂无用户")
	}
	fmt.Printf("%4s %15s %10s\n", "序号", "用户名", "余额")
	for i, user := range Users {
		fmt.Printf("%4d %15s %10d\n", i+1, user.Name, user.Balance)
	}
}
func DeleteUser(index int) {
	if index < 0 || index >= len(Users) {
		fmt.Println("删除用户失败，用户不存在")
		return
	}
	fmt.Println("确认是否删除用户？(y/n)")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm == "y" || confirm == "Y" {
		Users = append(Users[:index], Users[index+1:]...)
	}
	fmt.Println("用户删除成功")
}
func ModifyUser(index int, name string, password string, balance int) {
	if ExistUser(index) || ExistUser(name) {
		user := Users[index]
		user.Name = name
		user.Password = password
		user.Balance = balance
		fmt.Println("用户信息修改成功")
	}
}

// 可以使用接口和类型断言来实现类似重载的效果
func FindUser(param interface{}) *familyFinance.FamilyFinance {
	switch v := param.(type) {
	case int:
		if v < 0 || v >= len(Users) {
			fmt.Println("查找用户失败，用户不存在")
			return nil
		}
		return Users[v]
	case string:
		for _, user := range Users {
			if user.Name == v {
				return user
			}
		}
		fmt.Println("查找用户失败，用户不存在")
		return nil
	default:
		fmt.Println("参数类型错误")
		return nil
	}
}
func ExistUser(param interface{}) bool {
	switch v := param.(type) {
	case int:
		if v < 0 || v >= len(Users) {
			fmt.Println("查找用户失败，用户不存在")
			return false
		}
		return true
	case string:
		for _, user := range Users {
			if user.Name == v {
				return true
			}
		}
		fmt.Println("查找用户失败，用户不存在")
		return false
	default:
		fmt.Println("参数类型错误")
		return false
	}
}
