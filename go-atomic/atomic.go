package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wait sync.WaitGroup
var count int64

func increment(s string) {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println(s, i)
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	//wait.Wait()
	fmt.Println("last count value ", count)

	wait.Wait()
}
