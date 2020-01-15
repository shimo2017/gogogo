package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	var p *int = &n

	fmt.Println(p)

	fmt.Println(*p)

	fmt.Println("---------------------")
	var i int = 100
	one(&i)
	fmt.Println(i)
	fmt.Println("---------------------")
	fmt.Println(&i)
	fmt.Println(*&i)

}
