package main

import (
	"fmt"
)

const ArraySize = 7

//Hashtable
type Hashtable struct {
	array [ArraySize]*bucket
}
type bucket struct { //In this hashtable we use seprate chaining method
	head *bucketnode
}
type bucketnode struct {
	key  string
	next *bucketnode
}

func main() {
	myhashtable := Init() //Initalizing the bucket(Linked list in each node of array)
	list := []string{
		"Aswath",
		"Rock",
		"Bruce",
	}
	for _, v := range list {
		myhashtable.Insert(v) //Inserting the value in node
	}
	myhashtable.Delete("Bruce") //Deleting the value in node

	myhashtable.Search("Bruce")

	myhashtable.Search("Aswath") //Searching the value in node

}
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}

	}
	return result
}
func (h *Hashtable) Insert(k string) { //Inserting the value using hash functions
	index := hash(k)
	h.array[index].insert(k)
}
func (h *Hashtable) Search(k string) bool { //Searching the value using hash functions
	index := hash(k)
	return h.array[index].search(k)

}
func (h *Hashtable) Delete(k string) { //Deleting  the value using hash functions
	index := hash(k)
	h.array[index].delete(k)
}
func hash(k string) int { //Hash Function
	sum := 0
	for _, val := range k {
		sum += int(val)
	}
	return sum
}
func (b *bucket) insert(k string) { //Inserting value into the bucket
	if !b.search(k) {
		newNode := &bucketnode{key: k}
		newNode.next = b.head //setting up the head
		b.head = newNode
	} else {
		fmt.Println("Already exist")
	}

}
func (b *bucket) search(k string) bool { //Searching the value in bucket
	currentnode := b.head
	for currentnode != nil {
		if currentnode.key == k {
			return true
		}
		currentnode = currentnode.next
	}
	return false
}
func (b *bucket) delete(k string) { //Deleting the value in bucket
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previosnode := b.head
	for previosnode.next != nil {
		if previosnode.next.key == k {
			previosnode.next = previosnode.next.next
		}
		previosnode = previosnode.next
	}
}
