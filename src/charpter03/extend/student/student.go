package student

type Student struct {
	Name  string
	Age   int
	Score int
}
type Pupli struct {
	Student
	Zuoye string
}
type Graduate struct {
	Student
	Lunwen string
}
