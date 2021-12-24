package main

import "fmt"

func main() {
	x := 45
	fmt.Println("Value of x is :", x)
	y := &x

	var ptr *int

	*y = 23
	fmt.Println("Value of x is :", x)

	name := "Aswath"
	fmt.Println("Name is :", name)
	newname := &name
	*newname = "Rock"
	fmt.Println("New Name is :", name)

	ptr = y
	fmt.Println(*ptr)

}
