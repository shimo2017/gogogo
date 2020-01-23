package main

import "fmt"

// Vertex ...
type Vertex struct {
	X, Y int
}

// Area ...
func (v Vertex) Area() int {
	return v.X * v.Y
}

// Area ...
func Area(v Vertex) int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))

	fmt.Println(v.Area())

	fmt.Println("-------------------------")

	v.Scale(10)
	fmt.Println(v.Area())
}
