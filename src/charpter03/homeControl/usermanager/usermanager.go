package usermanager

import (
	"fmt"
)

type User struct {
	Name     string
	Password string
}

func NewUser(name string, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
func Login() {
	var username, password string
	fmt.Println("请输入用户名:")
	fmt.Scanln(&username)
	fmt.Println("请输入密码:")
	fmt.Scanln(&password)
}
func Usermanager() {
	fmt.Println("1. 查看用户信息")
	fmt.Println("2. 修改用户信息")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("查看用户信息")
	case 2:
		fmt.Println("修改用户信息")

	}
}
