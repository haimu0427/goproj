package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "marry"
	fmt.Println("test()", p.Name)
}
func (p *Person) test2() {
	p.Name = "dog"
	fmt.Println("test2()", p.Name)
}
func main() {
	p := Person{Name: "John"}
	p.test()
	(&p).test()
	fmt.Println("main()", p.Name) // Output: main() John
	// The method test() does not modify the original Person instance
	p.test2()
	(&p).test2()
	fmt.Println("main()", p.Name) // Output: main() dog
}
