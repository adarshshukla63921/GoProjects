package main

import "fmt"

// essentially the baseValue acts a private variable for the returned function.
func credit() func(amount int) int {
	baseValue := 1000
	return func(amount int) int {
		baseValue += amount
		return baseValue
	}
}
func main() {
	c := credit()
	// the value of count is rememebered even after counter() has finished executing.
	fmt.Println(c(500)); // 1
	fmt.Println(c(500)); // 2
	fmt.Println(c(1000)); // 3
}