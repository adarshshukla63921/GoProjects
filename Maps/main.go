package main

import "fmt"
func charactersInString(str string) {
	charMap := make(map[string]int)

	// counting the occurance of each character in a string
	for i:=0 ;i<len(str); i++{

		charMap[string(str[i])]++;
	}


	// traversal over a map.
	for key, value := range charMap {
		fmt.Printf("Character: %s, Count: %d\n", key, value)
	}
}
func main() {
	// Maps In Go
	m := make(map[string]int)
	// assigning values to a map.
	m["one"] = 1
	m["two"] = 2

	// if values does not exist, it will return the zero value of that data-type.
	fmt.Println("Value of a key that does not exist : ",m["three"])
	fmt.Println()
	// deleting a key-value pair
	delete(m, "two")

	// printing the map after deletion.
	fmt.Println("Map after deletion : ",m)
	fmt.Println()

	// updating values
	m["one"] = 11
	fmt.Println("Map after updating the value : ",m)
	fmt.Println()

	// safe look up in map. 
	value , ok := m["four"]
	if ok {
		fmt.Println("Safe Look Up OK-Value : ", value)
	} else {
		fmt.Println("Safe Look Up OK-Value : ", value)
	}
	fmt.Println()

	// accessing a value by using key.
	fmt.Println("Calling the value of a key that exists in the map : ",m["one"])
	fmt.Println()


	// calling a function.
	charactersInString("ababccbbaca");
}