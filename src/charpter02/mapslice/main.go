package main

import (
	"fmt"
)

func main() {
	slice_map := make([]map[string]string, 1, 10)
	slice_map[0] = make(map[string]string)
	slice_map[0]["name"] = "John"
	slice_map[0]["age"] = "20"
	slice_map = append(slice_map, map[string]string{
		"name": "Bob",
		"age":  "25",
	})

	slice_map = append(slice_map, map[string]string{
		"name": "Alice",
		"age":  "30",
	})

	fmt.Println("Slice of maps:", slice_map)

	map1 := make(map[int]int)
	map1[3] = 100
	map1[4] = 20
	map1[1] = 30
	map1[2] = 25
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("map1[%d] = %d \n", i, map1[i])
	// }
	for k, v := range map1 {
		fmt.Printf("Key: %d, Value: %d\n", k, v)
	}
}
