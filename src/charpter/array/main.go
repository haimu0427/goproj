package main

import (
	"fmt"
)

func main() {
	var m [5]int
	fmt.Println(m)
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	var arr3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	var arr4 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)
	var arr5 = [...]int{1: 200, 0: 100, 2: 300, 3: 400, 4: 500}
	fmt.Println(arr5)
	for _, v := range arr5 {
		fmt.Println(v)
	}
	heroes := [5]string{"Batman", "Superman", "Wonder Woman", "Flash", "Green Lantern"}
	fmt.Println(heroes)
	for i, v := range heroes {
		fmt.Printf("hero %v is %v\n", i, v)
	}
}
