package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Rectangle struct {
	LeftTop     Point `json:"left_top"`     // LeftTop Point
	RightBottom Point `json:"right_bottom"` // RightBottom Point
}
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	rect := Rectangle{
		LeftTop:     Point{0, 0},
		RightBottom: Point{10, 10},
	}
	fmt.Println("Rectangle:", rect)

	// Creating a Monster instance
	monster := Monster{
		Name:  "Goblin",
		Age:   5,
		Skill: "Stealth",
	}
	fmt.Println("Monster:", monster)

	json, _ := json.Marshal(monster)
	fmt.Println(string(json))

}
