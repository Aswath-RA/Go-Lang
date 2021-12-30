package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//go messager("Aswath") //go keyword is used for creating goroutines
	go messager("Rock")
	go messager("Aswath")
	wg.Add(2)
	//delay to print the message
	wg.Wait()
}

func messager(name string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("Hi Welcome to Golang", name)
	}
}
