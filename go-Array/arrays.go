package main

import "fmt"

func main() {
	// Array is a collection of homogenus elements ,Which can be accessed by index value.
	
	/*
	     Types of Array Declaration.
	     
	     1) var array-name = [length] data-type {values}
	     2) array-name = [length] data-type {values}
	     3) array-name = [...] data-type {values} 
	     4) array-name = [length] data-type {index:value}     
	     
	*/  
	
	  var arr = [5] int {1,2,3,4,5}
	  var colours =[2] string {"Orange","Blue"} 
	  
	  animals := [2] string {"Dog","Cat"}
	  
	  bikes := [...] string {"Yahama","BMW","Benelli"}
	  
	  cars := [3] string {1:"Audi",2:"Benz"}
	  
	  fmt.Println(arr)
	  fmt.Println(colours)
	  fmt.Println(animals)
	  fmt.Println(bikes)
	  fmt.Println(cars)
	  
	  animals[1]="Tiger"

	  fmt.Println(animals)
	  
	  fmt.Println("Length of Bikes: ",len(bikes))
	  fmt.Println("Length of Cars[1] : ",len(cars[1]))
	  
	  
}
