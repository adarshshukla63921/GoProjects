package main

import "fmt"

func main() {
	var name string;
	fmt.Scan(&name)

	if(name == "Adarsh"){
		fmt.Println("Correct Name")
	}else { // else must follow right after the closing brace of if block. 
		fmt.Println("Incorrect Name")
	}
}