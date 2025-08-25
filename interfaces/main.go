package main

type Shape interface {
	String() string
}

type Circle struct {
}

func (c Circle) String() string {
	return "Circle"
}

type Rectangle struct {
}

func (r Rectangle) String() string {
	return "Rectangle"
}

func ProcessShape(s Shape) string {
	return "this is a " + s.String()
}

func main() {
	c := Circle{}
	r := Rectangle{}

	println(ProcessShape(c))
	println(ProcessShape(r))
}
