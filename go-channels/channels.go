package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2) //Buffer channel of value 2

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		value, ischannelopen := <-ch
		fmt.Println("Value of the channel: ", value)
		fmt.Println("Is Channel open: ", ischannelopen)
		fmt.Println(<-ch)

	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 6
		ch <- 7
		close(ch)

	}(ch, wg)

	wg.Wait()
}
