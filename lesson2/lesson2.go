package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	cunter := incrementGenerator()
	fmt.Println(cunter())
	fmt.Println(cunter())
	fmt.Println(cunter())

}
