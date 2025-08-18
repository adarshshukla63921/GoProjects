package main 

import "fmt"

func main() {
	// methods to declare an array in go.
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	// variableName := [size]type{... 'values seperated by ,' ...}
	fmt.Println(arr)

	// another way to declare an array in go.
	var arr2[5] int 
	for i:=0;i<len(arr2);i++{
		arr2[i] = i + 1
	}
	fmt.Println(arr2)
}