package main

import "fmt"

func main() {
	fmt.Println("Variables in Go")

	// 1. Basic Declaration with 'var'
	// Syntax: var <name> <type>
	var age int              // default value for int = 0
	var isStudent bool       // default value for bool = false
	var username string      // default value for string = ""
	
	fmt.Println("Default int value:", age)
	fmt.Println("Default bool value:", isStudent)
	fmt.Println("Default string value:", username)

	// Assigning a value after declaration
	age = 25
	fmt.Println("Updated age:", age)

	// 2. Declaration + Initialization in one statement
	var score int = 95
	var city string = "Delhi"
	fmt.Println("Score:", score, "City:", city)

	// 3. Multiple variables declared and initialized at once
	var temperature, country = 28.5, "India" // Type is inferred
	fmt.Println("Temperature:", temperature, "Country:", country)

	// 4. Short variable declaration ':='
	// This is allowed only inside functions (not at package level).
	price := 49.99
	isAvailable := true
	fmt.Println("Price:", price, "Is Available:", isAvailable)

	// 5. Constants with 'const' keyword
	const Pi = 3.14159
	const WelcomeMessage = "Hello, Go!"
	fmt.Println("Pi:", Pi, "Message:", WelcomeMessage)

	// 6. Zero values are automatic
	// Variables are by default initialized with 0, false

	// 7. Type Inference
	// Go automatically assigns the type based on the value.
	autoNumber := 42           // inferred as int
	autoText := "Learning Go" // inferred as string
	fmt.Printf("autoNumber: %d (type: %T)\n", autoNumber, autoNumber)
	fmt.Printf("autoText: %s (type: %T)\n", autoText, autoText)

	// 8. Grouped Declaration
	var (
		firstName = "Adarsh"
		lastName  = "Shukla"
		ageYears  = 22
	)
	fmt.Println("Full Name:", firstName, lastName, "| Age:", ageYears)

	// 9. Redeclaration rules
	// Using := allows redeclaration ONLY if at least one new variable is introduced.
	a, b := 1, 2
	fmt.Println("a:", a, "b:", b)
	a, c := 3, 4 // 'a' is reassigned, 'c' is new
	fmt.Println("a:", a, "c:", c)

	// 10. Variable naming conventions
	// Use camelCase for variables
	// Start with lowercase for internal use
	// Start with uppercase if exporting (public)
}
