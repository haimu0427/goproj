package main

import (
	"errors"
	"fmt"
	_ "time"
)

func main() {
	num1 := 10
	num2 := new(int)
	fmt.Printf("num1:%T %v %v\n", num1, num1, &num1)
	fmt.Printf("num2: %T %v %v", num2, num2, &num2)
	err := readconfig("config.yam")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Config file read successfully")
	}

}

func readconfig(name string) (err error) {
	if name == "config.yaml" {
		return nil
	} else {
		return errors.New("config file is not correct")
	}
}
