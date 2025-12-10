package main

import "fmt"
func main() {
	// Array is a numbered sequnence of elements of a specific length

	// var arr_name [size]type

	var arr [5]int; // all elements are initialized 0 for int type array.
	fmt.Println(len(arr)); // to get length of array

	// short hand 
	nums := [3]int{10,4,5};
	fmt.Println(nums);

	// 2d arrays 
	matrix := [2][3]int {{1,2,10},{3,5,7}};
	fmt.Println(matrix);
}