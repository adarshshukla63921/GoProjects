package FunctionPacakage

import "fmt"

// does not take any input, and prints some result.
func NoInput(){
	fmt.Println("This function doesn't take any i/p or give an op.")
}

// takes input and returns a value
func Add(x, y int) int {
	return x+y 
}

// takes input and returns more than 1 value.
func Swapper(str1 , str2 string) (string,string){
	return str2, str1  
}

// take input and returns 'named' values
func Spilter(num int) (x,y int){
	x=num*4/9
	y=num-x
	return  // naked return -> returns the named return values (i.e x and y)
} 
