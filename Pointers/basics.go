package main

// using variables to swap values.
// call by address -> to modify value at address.
func swapIntegers(valueA *int, valueB *int) {
	temp := *valueA
	*valueA = *valueB
	*valueB = temp
}

// the below function will not swap the values of c and d in main function.
// it will simply copy c and d into something and somethingelse respectively.
// and will swap those values itself. This is 'call by value'. 
func swapInt(something int, somethingelse int){
	temp := something
	something = somethingelse
	somethingelse = temp
	println("Value of something : ",something)
	println("Value of somethingelse : ",somethingelse)
	println()
}

func main() {

	// declaring a pointer. 
	// var nameOfPointer *Type
	var ptr *int; 
	var value int = 10;

	// assigning address of value to pointer
	ptr = &value;
	println("Value of value:", value)
	// getting the address using the &operator
	println("Address of value:", ptr)
	// dereferencing = *pointerName {pointer must hold an address to access the value at that address.}
	println("Current value via pointer:", *ptr)
	// modifying value via pointes.
	*ptr = 20;
	println("New value of value:", value)
	println()


	// using functions with pointers.
	println("Swaping values using pointers.")
	a := 445;
	b := 554;
	swapIntegers(&a , &b)
	println("Value of a after swap:", a)
	println("Value of b after swap:", b)
	println()

	// trying to swap values without pointers.
	println("Trying to swap values without pointers.")
	c := 100;
	d := 200;

	swapInt(c , d)
	println("Value of c after swap:", c)
	println("Value of d after swap:", d)

}