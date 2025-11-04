package FunctionsInGo

import(
	"fmt"
)

// functions may take 0 or more arguments.
func NoArgs(){
	fmt.Println("This function has no arguments.")
}
// function with arguments and a return type (int).
// since both a and b are int, we can write them as 'a,b int' or 'a int, b int'.
func AddTwoNumbers(a , b int) int {
	return a+b
}

// function with multiple return values.
func SwapValue(x,y int) (int,int){
	var temp int = x
	x=y
	y=temp
	return x,y
}