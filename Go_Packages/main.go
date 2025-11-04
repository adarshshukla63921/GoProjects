package main

import (
	"Go_Packages/util" // creating your own custom packages.
	"fmt"
)

func main(){
	fmt.Println("Hello, Go Packages!")
	util.PrintMessage("Adarsh")
	var num int 
	fmt.Scan(&num)
	fmt.Println("Is",num,"prime ?",util.CheckPrime(num))
}