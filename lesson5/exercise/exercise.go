package main

import "fmt"

func main() {
	q1()

	// 回答
	fmt.Println("10ですー")

	fmt.Println("------------------------------------")

	q2()

	// 回答
	fmt.Println("600ですー")
}

func q1() {
	// Q1 以下の結果は正しく表示されるのか、されないのか
	//    正しく表示される場合は、なにが表示されるか
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j)
}

func q2() {
	// Q2 以下の結果は正しく表示されるのか、されないのか
	//    正しく表示される場合は、なにが表示されるか
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	p2 = p1
	j = *p2 + i
	fmt.Println(j)
}
