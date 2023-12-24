package main

import "fmt"

// Payer is interface for doing transactions
type Payer interface {
	Transaction() string
}

// AlphaPayment tool to manipulate transactions with specific interface
type AlphaPayment struct {
	id string
}

// Pay is method equal to Transaction()
func (a *AlphaPayment) Pay() string {
	return a.id
}

// AlphaAdapter is adapter struct to Payer interface
type AlphaAdapter struct {
	ap *AlphaPayment
}

// Transaction method implements Payer interface
func (a *AlphaAdapter) Transaction() string {
	return a.ap.Pay()
}

func main() {
	alpha := &AlphaPayment{id: "smth"}
	alphaAdapter := &AlphaAdapter{ap: alpha}
	fmt.Println(alpha.Pay())
	fmt.Println(alphaAdapter.Transaction())
}
