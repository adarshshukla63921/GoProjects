package util

import "fmt"
import "math"
func PrintMessage(name string){
	fmt.Printf("Hello, %s! Welcome to Go Packages!\n", name)
}

func CheckPrime(num int) bool{
	if num <= 1 {
		return false
	}
	for i:=2;i<=int(math.Sqrt((float64(num))));i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}