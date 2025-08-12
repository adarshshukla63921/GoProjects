package main

import (
	"GoFunctions/FunctionPackage"
	"fmt"
)

func main(){
	fmt.Println("This is to show functions in go.")
	FunctionPacakage.NoInput()
	var str1, str2 string = "string1" , "string2"
	fmt.Println(FunctionPacakage.Add(10,32))
	fmt.Println(FunctionPacakage.Swapper(str1,str2))
	fmt.Println(FunctionPacakage.Spilter(9))
}