package main

// Channels in Golang allow for communication between goroutines.
// Help avoid race conditions

import "fmt"

func Calculator(num1 int , num2 int, ch chan int){
	var result int
	result = num1+num2
	// sending result **to** the channel
	ch <- result
}


func main(){
	// Creating a channel
	ch := make(chan int)

	// Spawning go routines
	// Order of execution is not fixed
	go Calculator(10,20,ch)
	go Calculator(30,50,ch)
	go Calculator(100,200,ch)


	// Receiving data from channel
	// Channels block until data arrives.
	// Taking data **from** the channel
	result1 := <- ch
	result2 := <- ch
	result3 := <- ch

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
}
