package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Gender string
	Id     int
	Score  int
}

func (stu *Student) TellYourself() string {
	return fmt.Sprintf("Name: %s, Age: %d, Gender: %s, Id: %d, Score: %d",
		stu.Name, stu.Age, stu.Gender, stu.Id, stu.Score)
}

func main() {
	stu := Student{
		Name:   "Alice",
		Age:    20,
		Gender: "Female",
		Id:     1,
		Score:  90,
	}
	fmt.Println(stu.TellYourself())
	fmt.Print()
}
