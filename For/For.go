package main

import "fmt"

func main() {
	// there is only for loop in go.

	// while loop
	i := 1

	for i <= 5 {
		fmt.Println(i);
		i++;
	}
	fmt.Println()
	// traditional for loop 

	for j := 10; j <= 50; j+=10 {
		fmt.Println(j);
	}
}