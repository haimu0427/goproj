package main

import (
	"fmt"
)

func main() {
	var intarr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice1 []int = intarr[1:3]
	fmt.Println("Slice1:", slice1, "需要尽量的保持不变的,一般是保持两倍, 但是可以是1.5倍", (slice1), len(slice1))
	fmt.Println("intarr[1]的地址", &intarr[1], "slice自己的地址是:", &slice1, "slice1[0]:地址", &slice1[0], "slice1[1]:地址", &slice1[1])

	var slice2 []int = make([]int, 3)       //对于切片的初始化,可以使用make函数,第一个参数是类型,第二个参数是长度,第三个参数是容量
	var slice3 []int = []int{1, 2, 3, 4, 5} //直接初始化切片
	fmt.Println("Slice2:", slice2, "Slice3:", slice3)
	fmt.Println("slice3 capacity:", cap(slice3), cap(slice2), "slice3 length:", len(slice3), len(slice2))

	var arr []int = []int{1, 2, 3, 4, 5, 6}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}
	slice = arr[1:4]
	slice1 = arr[:3]
	slice2 = arr[2:]
	slice3 = arr[:]
	fmt.Println("Slice1:", slice1, "Slice2:", slice2, "Slice3:", slice3)
	slice_s := slice[:2]
	fmt.Println("Slice:", slice_s)
	// 切片的扩容
	slice = append(slice, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	slice3 = append(slice3, slice3...)

	//切片之间的拷贝操作
	var slice_ = make([]int, 9)
	var slice_1 = make([]int, 9)
	copy(slice_, slice3)
	copy(slice_1, slice3)
	fmt.Println("Slice_:", slice_)
	fmt.Println("Slice_1:", slice_1)

}
