package main

import (
	"fmt"
)

type stack struct {
	items []int
}

func main() {
	fmt.Println("This is Stack in golang")

	//calling Push
	mystacks := stack{}
	fmt.Println(mystacks)
	mystacks.push(100)
	mystacks.push(200)
	fmt.Println(mystacks)
	mystacks.pop()
	fmt.Println(mystacks)

}

//push
func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

//pop it will return the removed item

func (s *stack) pop() int {
	l := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return l

}
