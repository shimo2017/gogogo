package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(redius float64) float64 {
	return func(redius float64) float64 {
		return pi * redius * redius
	}
}

func main() {
	cunter := incrementGenerator()
	fmt.Println(cunter())
	fmt.Println(cunter())
	fmt.Println(cunter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}
