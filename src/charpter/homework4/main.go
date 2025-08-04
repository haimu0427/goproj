package main

import (
	"fmt"
	"time"
)

// 这是一道题, 计算判断从1990年1月1日之后的某一天相隔多少天
func main() {
	var year, month, day int
	fmt.Println("请输入日期(格式: YYYY MM DD):")
	fmt.Scanln(&year, &month, &day)

	startDate := time.Date(2006, 5, 24, 0, 0, 0, 0, time.UTC)
	inputDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if inputDate.Before(startDate) {
		fmt.Println("输入的日期必须在2006年5月24日之后")
		return
	}

	daysDiff := inputDate.Sub(startDate).Hours() / 24
	fmt.Printf("从2006年5月24日到%d年%d月%d日相隔%.0f天\n", year, month, day, daysDiff)
}
