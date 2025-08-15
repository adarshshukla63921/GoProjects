package main

import "fmt"
func main(){
	fmt.Println("This file will show how to use pointers in go.")

	// declaring pointers in go.
	var ptrX, ptrY *int
	var x, y int = 10,20

	ptrX=&x
	ptrY=&y

	// printing address.
	fmt.Println(ptrX)
	fmt.Println(ptrY)

	//accessing values through the address.
	fmt.Println(*ptrX)
	fmt.Println(*ptrY)

	//modifiying values through pointers.
	*ptrX=20
	*ptrY=80
}