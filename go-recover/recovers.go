package main

import (
	"fmt"
)

func main() {
	fmt.Println("err: ", SaveDivide(10, 0))
	fmt.Println(SaveDivide(10, 10))
	fmt.Println("Addition of two numbers: ", adder(10, 5))

}

//
func SaveDivide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
func adder(x, y int) int {
	return x + y
}
