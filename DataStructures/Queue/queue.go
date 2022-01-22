package main

import "fmt"

type Queue struct {
	items []int
}

func main() {
	fmt.Println("This is Queue  in golang")
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.EnQueue(100)
	myQueue.EnQueue(200)
	myQueue.EnQueue(300)

	fmt.Println(myQueue)

	myQueue.DeQueue()

	fmt.Println(myQueue)

}
func (q *Queue) EnQueue(i int) {
	q.items = append(q.items, i)
}
func (q *Queue) DeQueue() int {
	removedvalue := q.items[0]
	q.items = q.items[1:]
	return removedvalue
}
