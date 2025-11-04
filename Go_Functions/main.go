package main

import (
	"Go_Functions/FunctionsInGo"
	"fmt"
)
func main(){
	fmt.Println("Hello, Functions!")
	FunctionsInGo.NoArgs()
	fmt.Println("the sum is",FunctionsInGo.AddTwoNumbers(10,20));
	fmt.Println("Enter two numbers to swap:")
	var x int
	fmt.Scan(&x)
	var y int
	fmt.Scan(&y)
	x,y = FunctionsInGo.SwapValue(x,y)
	fmt.Println("After swapping, x is",x,"and y is",y)
}