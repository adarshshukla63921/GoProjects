package main

import (
	"fmt"
	"math"
)

// allows for declaration of variables inside of a if condition, also said variable is avaible for use in else block.
func myPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v <= lim {
		return v
	}
	return lim
}

// go does not need () for conditions but {} are needed.
func ageCalc(age int) {
	if age <= 18 {
		fmt.Println("Child")
	} else {
		fmt.Println("Adult")
	}
}
func main() {
	fmt.Println("Using Control Flow in Go.")
	fmt.Println(myPow(2.0, 5.0, 64.0))
	ageCalc(10)
}
