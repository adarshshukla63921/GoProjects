package main

import (
	"fmt"
	"example.com/Functions/utils"
)


func main() {
	// Accessing an exported variable from another package.
	fmt.Println(utils.ExportedVariable)
	// using functions declared in another custom made package.
	fmt.Println(utils.ExportedFunction())
	num := 5
	fmt.Println("Incrementing 5 : ",utils.Increment(&num))
}