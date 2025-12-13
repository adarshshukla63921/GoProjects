package main

import "fmt"
func charactersInString(str string) {
	charMap := make(map[string]int)
	for i:=0 ;i<len(str); i++{
		// counting occurance of each character in a string
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

	// accessing a value by using key.
	fmt.Println(m["one"])
	charactersInString("adarsh");
}