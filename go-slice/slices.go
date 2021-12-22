package main

import "fmt"

func main () {
 
 	/* We can declare a slice by following ways
 	  1) creating a new slice 
 	  2) slicing an array 
 	  3) Creating slice using make()
 	 */


	slicea := [] int{1,2,3,4,5}
	fmt.Println("Length of Slice A is    : ",len(slicea))
	fmt.Println("Capacity of Slice A is  : ",cap(slicea)) 
	fmt.Println("")
	
	
	arr:=[5] int {1,2,3,4,5}
	sliceb:=arr[1:3]
	fmt.Println("Length of Slice B  is   : ",len(sliceb))
	fmt.Println("Capacity of Slice B is  : ",cap(sliceb)) 
	fmt.Println("")
	
	slicec:=make([]int ,5,10)
	fmt.Println("Length of Slice C is    : ",len(slicec))
	fmt.Println("Capacity of Slice C  is : ",cap(slicec))
	fmt.Println("")
	
	// Adding values using append()
	
	fmt.Println("Original value of Slice A is  :",slicea)
	slicea=append(slicea,6,7,8)
	fmt.Println("Appended value of Slice A is  :",slicea)
	
	//Changing  value using index
	
	slicea[4]=99
	fmt.Println("Value changed using index[4]  Slice A is  : ",slicea)
	
	
	
	
}