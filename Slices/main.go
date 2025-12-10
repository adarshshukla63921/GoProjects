package main

import "fmt"

// slices are dynamic arrays.
// most used constructs
// very useful methods are available
func main() {
	// declaration of slices.

	var nums []int; // uninitialized slice is nil
	fmt.Println(nums); // size = 0; 

	// make ([]type , length, capacity)
	newNums := make([]int, 5, 10); // length = 5, capacity = 10
	fmt.Println(newNums);

	// adding elements

	nums = append(nums, 10,20,30,30,60);
	fmt.Println(nums);
	fmt.Println(cap(nums), " ", cap(newNums))// to get capacity

	// copying elements from one slice to another
	copy(newNums , nums)
	fmt.Println("newNums is now : ",newNums);

	//slice operators
	fmt.Println("Slice operators ->")
	newSlice := []int {1,2,3,4,5,6,7,9,10};
	fmt.Println(newSlice[0:6]);
}