package main

import (
	"encapsulate/model"
	"fmt"
)

func main() {
	p := *model.NewPerson("坤坤")
	p.SetAge(18)
	p.Setsal(18000)
	fmt.Println(p.Name, "爱你呦")
	fmt.Println(p)

}
