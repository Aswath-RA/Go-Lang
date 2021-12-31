package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2) //Buffer channel of value 2

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { //channel way declaration value out of channel
		defer wg.Done()
		fmt.Println(<-ch)

	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) { //channel way declaration value into  channel
		defer wg.Done()
		ch <- 6
		close(ch) //closing channel after use

	}(ch, wg)

	wg.Wait()
}
