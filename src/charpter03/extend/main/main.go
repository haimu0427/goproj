package main

import (
	"extend/student"
	"fmt"
)

func main() {
	// Create instances of student, pupli, and Graduate
	s := student.Student{Name: "John", Age: 20, Score: 90}
	p := student.Pupli{Student: student.Student{
		Name:  "Jane",
		Age:   18,
		Score: 85,
	}, Zuoye: "Math Homework"}
	g := student.Graduate{Student: s, Lunwen: "Thesis on Go"}

	// Print the details
	fmt.Println("Student:", s)
	fmt.Println("Pupli:", p)
	fmt.Println("Graduate:", g)

}
