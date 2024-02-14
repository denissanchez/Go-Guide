package main

type Shape interface {
	area() float64
}

type Square struct {
	base float64
}

type Rectangle struct {
	base   float64
	height float64
}

func (c Square) area() float64 {
	return c.base * c.base
}

func (r Rectangle) area() float64 {
	return r.base * r.height
}

func calculate(s Shape) float64 {
	return s.area()
}

func main() {
	s := Square{base: 15}
	r := Rectangle{base: 10, height: 5}

	println(calculate(s))
	println(calculate(r))
}
