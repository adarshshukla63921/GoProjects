package main 

import "fmt"

func main() {
	fmt.Println("Defer in Go")

	var someNumber int = 5

	// defer delays the execution of the below line till the end, of the surrounding function, but stores the needed values for exec.
	defer fmt.Println("original val :",someNumber)
	someNumber=10
	fmt.Println("updated value :",someNumber)

	// Defers func calls are pushed onto a stack,  they are executed in last in first out order. 
}