package main

import "fmt"

// interface defines a 'thing' you can do
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// creating entities that can do the 'thing'
type CreditCard struct {
	cardNumber string
}

type Paypal struct {
	email string
}

// connecting 'entity' to the 'thing'
func (c CreditCard) ProcessPayment(amount float64) string {
	return "Processed credit card payment of $" + fmt.Sprintf("%.2f", amount)
}
func (p Paypal) ProcessPayment(amount float64) string {
	return "Processed PayPal payment of $" + fmt.Sprintf("%.2f", amount)
}

// using the interface's 'thing'
func Checkout(processor PaymentProcessor, amount float64) string{
	result := processor.ProcessPayment(amount)
	return result
}

func main(){
	// creating instances of entites
	creditCard := CreditCard{cardNumber: "WWWW-XXXX-YYYY-ZZZZ"}
	paypalEmail := Paypal{email: "adarsh@xyz.com"}

	// using entities to do the 'thing'
	x :=Checkout(creditCard, 100.0)
	y := Checkout(paypalEmail, 200.0)

	fmt.Println(x)
	fmt.Println(y)
}

// you can later add more entities that implement the same interface without changing existing code