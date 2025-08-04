package model

type student struct {
	name  string
	age   int
	score int
}

func NewStudent(name string, age int, score int) *student {
	return &student{
		name:  name,
		age:   age,
		score: score,
	}
}
func (s *student) GetAge() int {
	return s.age
}
func (s *student) GetName() string {
	return s.name
}
