package main

import "fmt"

func main() {
    var name string
    fmt.Scan(&name)

    switch name {
    case "Adarsh":
        fmt.Println("Legal Name")
    case "Seenu":
        fmt.Println("Pet Name")
    case "Creator":
        fmt.Println("Designation")
    case "Go":
        fmt.Println("Programming Language")
    default:
        fmt.Println("Unknown")
    }
}
