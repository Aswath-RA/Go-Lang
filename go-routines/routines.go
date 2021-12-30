package main

import (
	"fmt"
	"time"
)

func main() {
	go messager("Aswath") //go keyword is used for creating goroutines
	go messager("Rock")
	time.Sleep(3 * time.Millisecond) //delay to print the message

}

func messager(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hi Welcome to Golang", name)
	}
}
