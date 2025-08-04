package familyFinance

import (
	"fmt"
	"homeControl/usermanager"
)

type FamilyFinance struct {
	usermanager.User
	Balance int
	Details []string
	Change  []int
}

func SignUp(name string, password string, balance int) *FamilyFinance {
	return &FamilyFinance{
		User: usermanager.User{
			Name:     name,
			Password: password,
		},
		Balance: balance,
		Details: []string{},
		Change:  []int{},
	}

}
func CreateUser() *FamilyFinance {
	for {
		var username, password, password2 string
		var balance int
		fmt.Println("请输入用户名:")
		fmt.Scanln(&username)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)
		fmt.Println("请再次输入密码:")
		fmt.Scanln(&password2)
		if password == password2 {
			fmt.Println("需要存钱吗? 需要请按0, 不需要请按1")
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
			fmt.Println("注册成功")

		}
		return SignUp(username, password, balance)
	}
}

func (f *FamilyFinance) showNowBalance() {
	fmt.Printf("当前账户余额: %d\n", f.Balance)
}
func (f *FamilyFinance) ShowDetials() {
	fmt.Println("--------------显示收支明细--------------")
	if len(f.Details) == 0 {
		fmt.Println("暂无收支记录")
		return
	}
	f.showNowBalance()
	fmt.Printf("%4s %10s %10s %20s\n", "序号", "账户金额", "收支金额", "说明")
	for i, detail := range f.Details {
		fmt.Printf("%4d %10d %10d %20s\n", i+1, f.Balance, f.Change[i], detail)
	}

}
func (f *FamilyFinance) ShowInDetials() {
	var income int
	var note string
	fmt.Println("--------------登记收入--------------")
	fmt.Println("本次收入金额:")
	fmt.Scanln(&income)
	if income < 0 {
		f.ShowOutDetials()
	} else {
		f.Change = append(f.Change, income)
		f.Balance += income
		fmt.Println("本次收入说明:")
		fmt.Scanln(&note)
		f.Details = append(f.Details, note)
		f.showNowBalance()
	}

}

func (f *FamilyFinance) ShowOutDetials() {
	var out int
	var note string
	fmt.Println("--------------登记支出--------------")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&out)
	if out > f.Balance {
		fmt.Println("支出金额不能大于当前余额，请重新输入")
		return
	} else if out < 0 {
		f.ShowInDetials()
	} else {
		f.Change = append(f.Change, -out)
		fmt.Println("本次支出说明:")
		fmt.Scanln(&note)
		f.Details = append(f.Details, note)
		f.Balance -= out
		f.showNowBalance()
	}

}

func Transfer(from, to *FamilyFinance) {
	var amount int
	fmt.Println("开始转账功能, 请输入转账金额:")
	fmt.Scanln(&amount)
	if amount < 0 {
		fmt.Println("转账金额不能为负数，请重新选择")
		return
	} else if amount > from.Balance {
		fmt.Println("转账金额不能大于当前余额，请重新选择")
		return
	} else {
		from.Balance -= amount
		to.Balance += amount
		from.Details = append(from.Details, fmt.Sprintf("转账给%s: %d", to.Details, amount))
		to.Details = append(to.Details, fmt.Sprintf("转账来自%s: %d", from.Details, amount))
	}

}
