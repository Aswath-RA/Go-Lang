package main 

import "fmt"

func main() {
	//For loop
	
	//simple for loop program printing hello for 5 times.
	
	for i:=0;i<5;i++ {
		fmt.Println("hello")
	}
	
	//for loop with continue and break
	fmt.Println()
	
	for i:=0;i<10;i++ {
		if i==3 {
			continue // 3 will be skipped
		}else if i==7 {
			fmt.Println("loop terminated reached 7  ")
			break
		}else {
			fmt.Println("value of i is : ",i)
		}
	}
	fmt.Println()
	
	//Nested For Loop
	 
	 animals :=[2] string {"cow","elephant"}
	 fruits :=[3] string {"orange","apple","jackfruit"}
	 
	 for i:=0;i<len(animals);i++ {
	 	for j:=0;j<len(fruits);j++ {
	 	fmt.Println(animals[i],fruits[j])
	 	}
	 }
	 
	 //for loop using range
	 fmt.Println("")
	 
	 for _,value := range(fruits) {
	 	fmt.Println(value)
	 }
	  
}