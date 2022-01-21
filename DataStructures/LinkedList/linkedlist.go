package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head   *node
	length int
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 56}
	node2 := &node{data: 4}
	node3 := &node{data: 6}
	node4 := &node{data: 90}
	mylist.Insert(node1)
	mylist.Insert(node2)
	mylist.Insert(node3)
	mylist.Insert(node4)
	fmt.Println(mylist) //It will print the last one adress and length of list
	mylist.Printdata()
	mylist.Deletedata(6) //Deleting the value with node 3
	mylist.Printdata()

}
func (l *linkedlist) Insert(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedlist) Printdata() {
	printvar := l.head
	for l.length != 0 {
		fmt.Print("\t", printvar.data)
		printvar = printvar.next
		l.length--

	}
	fmt.Println()
}
func (l *linkedlist) Deletedata(val int) {
	if l.length == 0 { //if the value is 0
		return
	}
	if l.head.data == val { //if the deleting value is head move the head to next node
		l.head = l.head.next
		l.length--
	}
	todelete := l.head
	for todelete.next.data != val { //Previous data to delete
		if todelete.next.next == nil { //no  same value to delete
			return
		}
		todelete = todelete.next

	}
	todelete.next = todelete.next.next
	l.length--
}
