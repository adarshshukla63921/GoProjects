package main

// Generics lets you write 1 function / type that works with different data types without losing type safety.

import "fmt"

// A generic type called "Box" that can hold a value of any type T
type Box[T any] struct {
	value T
}

// Functions that "Box" can use to set & get value of any datatype T
func (b *Box[T]) Set(val T) {
	b.value = val
}
func (b *Box[T]) Get() T {
	return b.value
}


// Main function to demonstrate the usage of the generic Box type
func main() {
	// Box of type : int
	var intBox Box[int]
	x := 10
	intBox.Set(x)
	fmt.Println(intBox.Get())

	// Box of type : string
	var stringBox Box[string]
	str := "this is a string"
	stringBox.Set(str)
	fmt.Println(stringBox.Get())

	// Box of type : bool
	var boolBox Box[bool]
	boolBox.Set(true)
	fmt.Println(boolBox.Get())
	
}