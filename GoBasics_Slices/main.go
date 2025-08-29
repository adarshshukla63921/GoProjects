package main 

import "fmt"

func main(){
	fmt.Println("Using Slices in Go.")

	arr := [10]int{10,20,30,40,50,60,70,80,90,100}

	// delcaration of a slice.
	var someSlice []int = arr[0:5]
	fmt.Println(someSlice)
	// slice acts like a refernce to the array. 
	someSlice[0] = 100
	fmt.Println(someSlice) // modifying the element of a slice will modify its underlying array also.
	fmt.Println(arr)
}