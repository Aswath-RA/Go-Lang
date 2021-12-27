package main

import (
	"fmt"
	"reflect"
)

type details struct {
	fname   string
	lname   string
	age     int
	balance float64
}

type myType string

func showDetails(i, j interface{}) {
	t1 := reflect.TypeOf(i)
	k1 := t1.Kind()
	t2 := reflect.TypeOf(j)
	k2 := t2.Kind()
	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Println("Type of second interface:", t2)
	fmt.Println("Kind of second interface:", k2)

}

func main() {
	iD := myType("12345678")
	person := details{
		fname:   "Go",
		lname:   "Geek",
		age:     32,
		balance: 33000.203,
	}
	showDetails(person, iD)
}
