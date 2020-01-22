package main

import "fmt"

// Vertex ...
type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func chargeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 100
	fmt.Println(v.X, v.Y)

	// 構造体にセットできる
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	// キー指定なしの場合は順番にセットできる
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	// 初期構造体を出力すると、初期値が設定されている
	v4 := Vertex{}
	fmt.Println(v4)
	fmt.Printf("%T %v\n", v4, v4)

	// 変数宣言の場合、上記と同じ
	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	// ポインタが返ってくる
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	// 上記と同じ意味
	v7 := &Vertex{}
	fmt.Printf("%T, %v\n", v7, v7)

	fmt.Println("-----------------------")
	v8 := Vertex{1, 2, "test"}
	changeVertex(v8)
	fmt.Println(v8)

	v9 := &Vertex{1, 2, "test"}
	chargeVertex2(v9)
	fmt.Println(v9)
}
