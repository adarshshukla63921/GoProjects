package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait groups in go are used to sync goroutines.

func downloadFile(str string, wg *sync.WaitGroup){
	
	defer wg.Done() // garantee that Done is called when the function exits

	fmt.Println("Download for file ", str, "is starting")
	fmt.Println()
	time.Sleep(2 * time.Second)
	fmt.Println("Download sucessfully completed for file ", str)
	fmt.Println()
}


func main(){
	var wg sync.WaitGroup

	wg.Add(3)

	go downloadFile("video.mp4",&wg)
	
	go downloadFile("image.jpg",&wg)
	
	go downloadFile("document.pdf",&wg)

	wg.Wait()
	fmt.Println("All workers completed")
}

// wg.Add() -> Used to increment / decrement the counter of the wait group

// wg.Done() -> Decrements the counter of the wait group by one

// wg.Wait() -> Blocks until the counter of the wait group is zero

// sync.WaitGroup is used to create wait groups in go

// time.Sleep() is used to simulate a delay in the program