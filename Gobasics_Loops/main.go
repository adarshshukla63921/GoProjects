package main

import "fmt"

func main() {
	
	fmt.Println("Loops in Go.")

	// Go has for loop only.
	/*
		the for loop has 3 components -> for [init statement]; [condition]; [post statement]
	*/
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Work around for while loop is aviable though.
	// the init statement and the post statement is optional, and so we can do something like this.
	sum := 1
	fmt.Println("Sum current value :", sum)
	for sum < 100 {
		sum++
	}
	fmt.Println("sum updated value :", sum)

	/*
		infinite loop ->
		for { // some block of code. }
	*/
}
