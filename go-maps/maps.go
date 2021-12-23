package main
import ("fmt")

func main() {

//Simple map

var mycars  = map[string]int{"Audi":1,"BMW":4,"RollsRoyce":1}
fmt.Println("mycars --> ",mycars)
fmt.Println()
fmt.Println("MY Car Collections")
fmt.Println("How many Audi car : ",mycars["Audi"])
fmt.Println("How many BMW car  : ",mycars["BMW"])

//Declaring maps
var fruits = make(map[int]string)
fmt.Println("Fruits: ",fruits)
fmt.Println()


//Value adding to map
fruits[1]="Orange"
fruits[2]="Strawberry"
fmt.Println("value added to Fruits: ",fruits)
fmt.Println()

//Value deleting from map
delete(fruits,1)
fmt.Println("value Deleted from Fruits: ",fruits)
fmt.Println()

//Iterating in maps
fmt.Println("Printing maps using for loop")
for key,value :=range(mycars) {
	fmt.Println(key,":",value)
}


}