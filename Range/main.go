package main

import "fmt"

func main() {
	// range is used to iterate over elements in arrays, slices, maps etc.

	arr := []int{10, 20, 30, 40, 50}

	// for idx, value = range collection
	for index, value := range arr {
		println(value, "at the index ", index)
	}
	fmt.Println()


	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4

	for key, value := range m {
		println("Key :", key, ", Value :", value)
	}
}