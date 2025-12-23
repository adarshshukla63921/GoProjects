package main
// Goroutines allow functions to run concurrently with other functions.
// These are done using lightweight threads managed by the Go.


import (
	"fmt"
	"time"
)

func downloadFile(name string){
	
	// time.Sleep() -> suspends the execution of a function for a given duration.
	fmt.Println("Downloading file :", name)
	time.Sleep(2 * time.Second)
	fmt.Println("Successfully Downloaded :", name)
}

func main() {
	// go keyword is used to spawn a goroutine.
	go downloadFile("video.mp4")
	go downloadFile("document.pdf")
	go downloadFile("image.jpg")

	// Main function must wait all for all goroutines to finish before exiting.
	// Otherwise, the program will terminate immediately and goroutines may not complete their execution.
	time.Sleep(4 * time.Second)
	fmt.Println("Exiting Main.")
}