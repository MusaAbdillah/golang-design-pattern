package main

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet {}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	} 

	fmt.Println("Wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil 
}
