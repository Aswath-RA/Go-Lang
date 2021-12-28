package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := sqrtn(-78)
	if err != nil {
		fmt.Println("Error Message : ", err)
	} else {
		fmt.Println(result)
	}
}
func sqrtn(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative number")

	} else {
		return math.Sqrt(x), nil
	}
}
