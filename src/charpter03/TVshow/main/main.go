package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type Tv struct {
	Goods
	Brand
	Size int
}
type Tv2 struct {
	*Goods
	*Brand
	Size int
	int  // This is an anonymous field
}

func main() {
	tv := Tv{
		Goods: Goods{
			Name:  "xiaomi tv",
			Price: 2999.99,
		},
		Brand: Brand{
			Name:    "xiaomi",
			Address: "beijing",
		},
		Size: 55,
	}
	fmt.Println("TV Details:", tv)

	tv2 := Tv2{
		Goods: &Goods{
			Name:  "samsung tv",
			Price: 4999.99,
		},
		Brand: &Brand{
			Name:    "samsung",
			Address: "seoul",
		},
		Size: 65,
		int:  99, // Anonymous field
	}
	fmt.Println("TV2 Details:", *tv2.Goods, *tv2.Brand, tv2.Size, tv2.int)

}
