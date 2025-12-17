package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// creating a 'Person' struct instance.
	p1 := Person{
		name: "Adarsh",
		age:  23,
	}
	fmt.Println(p1)

	// assigns 0-value to p2 for string and int.
	p2 := Person{}
	fmt.Println(p2)

	// accessing feilds of structs.
	fmt.Println("Name:", p1.name)
}