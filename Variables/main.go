package main

import (
	"fmt"
	"reflect"
)

func main () {
	// var name type

	// types of primitive variables in go. [also called simple values in go.]

	var a int = 20;  // int8, int16, int32, int64 are also available
	var a1 uint = 30; // uint8, uint16, uint32, uint64 are also available
	var b string = "go"
	var c float32 = 10.5; // float64 is also available
	var d bool = true;
	var complex complex64 = 1+2i; // complex128 is also available

	// all declared values must be used.
	fmt.Println(a, a1, b, c, d, complex)

	// go lang can also infer the type of a variable.

	var x = 100;
	var y = "golang"
	var z = true
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z))

	// short hand syntax for variable declaration

	xstr := "short hand"
	fmt.Println(xstr)
}