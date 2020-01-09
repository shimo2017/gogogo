package main

import "fmt"

func main() {
	fmt.Println("continue-----------")
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("break-----------")
	for i := 0; i < 10; i++ {
		if i > 3 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	fmt.Println("sum----------------")
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
