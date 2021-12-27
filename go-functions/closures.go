package main

import "fmt"

func main() {
	adder := func(x int, y int) int {
		return x + y
	}

	fmt.Println(adder(5, 10))

}
