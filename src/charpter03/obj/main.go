package main

// golang是一种面向对象(接口)的编程语言
import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}
type Dog struct {
	Name  string
	Age   int
	Color string
	Hobby string
	slice []int
}
type Person struct {
	Name string
	Sex  bool
	Age  int
}

func (c Cat) Meow() {
	fmt.Println(c.Name, "says Meow!")
}

func main() {
	cat := Cat{
		Name:  "Tom",
		Age:   3,
		Color: "Black",
	}
	fmt.Println(cat)
	cat.Meow()
	dog := Dog{
		Name:  "Jerry",
		Age:   5,
		Color: "Brown",
		Hobby: "Chasing Cats",
	}
	dog.slice = []int{1, 2, 3, 4, 5}
	fmt.Println(dog)

	var dog2 Dog
	dog2.Name = "Buddy"
	dog2.Age = 4
	fmt.Println(dog2)

	//3
	p := new(Dog)
	var p2 *Dog = new(Dog)
	p.Name = "Max"
	p.Age = 2
	(*p2).Name = "Bella"
	(*p2).Age = 3
	fmt.Println(p)
	fmt.Println(p2)

	//4
	var person *Person = &Person{}
	(*person).Name = "John"
	(*person).Sex = true
	(*person).Age = 30
	fmt.Println(person)
	fmt.Println(*person)

	//5

}
