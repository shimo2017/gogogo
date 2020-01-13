package main

import "fmt"

func main() {
	// Q1 以下のスライスから一番小さい数を探して出力するコードを書いてください
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	// 回答
	minVal := 0
	nextVal := 0
	for k, v := range l {
		if k == 0 {
			minVal = v
			continue
		}
		nextVal = v
		if nextVal <= minVal {
			minVal = nextVal
		}

	}
	fmt.Println(minVal)

	fmt.Println("------------------------")

	// 以下の果物の価格の合計を出力するコードを書いてください
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	// 回答
	summary := 0
	for _, v := range m {
		summary = summary + v
	}
	fmt.Println(summary)
}
