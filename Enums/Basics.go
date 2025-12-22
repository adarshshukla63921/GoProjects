package main

import "fmt"

type TrafficLight int

const (
	Red TrafficLight = iota  // 0
	Yellow  			    // 1
	Green                  // 2
)

func canDrive(t TrafficLight)  string {
	switch t {
	case Red:
		return "Stop"
	case Yellow:
		return "Slow Down"
	case Green:
		return  "Go Ahead"
	default:
		return "Unknown"
	}
}
func main() {
	fmt.Println(canDrive(Red))
	fmt.Println(canDrive(Yellow))
	fmt.Println(canDrive(Green))
	fmt.Println(canDrive(TrafficLight(3))) // for any value <= 3 it will print "Unknown"
}