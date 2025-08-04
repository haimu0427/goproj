package main

import (
	"fmt"
)

func main() {
	var year int
	var month int
	var monthDays [13]int = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for {
		fmt.Printf("请输入年:")
		fmt.Scanf("%d\n", &year)
		fmt.Printf("请输入月:")
		fmt.Scanf("%d\n", &month)
		if month < 1 || month > 12 {
			fmt.Println("输入月份不正确!")
			continue
		} else {
			if isLeapYear(year) && month == 2 {
				fmt.Println("这个月的天数为:29")
			} else {
				fmt.Println("这个月的天数为:", monthDays[month])
			}
		}

	}
}

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}
