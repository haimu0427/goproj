package main

import (
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}
type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	student := []Student{
		{"Alice", 20, 88.5},
		{"Bob", 22, 92.0},
		{"Charlie", 21, 85.0},
	}
	sort.Sort(StudentSlice(student))
}
