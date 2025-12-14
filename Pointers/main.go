package main

import "fmt"

func main() {
	var someNumber int = 42

	// here we are declaring a pointer to a variable that stores integer values 
	var pointerToNumber *int;
	// assigning the address of 'someNumber' to the pointer
	pointerToNumber = &someNumber; 
	

	fmt.Println(someNumber); // value of 'someNumber'
	fmt.Println(pointerToNumber); // address that stores 'someNumber'

	fmt.Println(*pointerToNumber); // accessing value of 'someNumber' using its pointer 'pointerToNumber'

	// As you can see above '*' has two uses
	// 1) To declare a variable as a pointer.
	// 2) To access the value stored at the address the pointer is pointing to.
}