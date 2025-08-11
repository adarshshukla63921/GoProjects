package main

import (
	"PackagesInGo/helper"
	"fmt"
)

func main(){
	fmt.Println("this is a func.") // from the fmt package.
	helper.SayHello("Adarsh") // from our own custom pacakage, which is "helper"
}