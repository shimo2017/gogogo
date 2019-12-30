package main

import (
	"fmt"
)

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return result
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	r4 := convert(100)
	fmt.Println(r4)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(2)
}
