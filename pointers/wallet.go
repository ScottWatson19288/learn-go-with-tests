package main

import (
	"errors"
	"fmt"
)

// Bitcoin is a type of int
type Bitcoin int

// Stringer interface for strings
type Stringer interface {
	String() string
}

// Wallet is were we keep our Bitcoin
type Wallet struct {
	balance Bitcoin
}

// Deposit takes int some amount of Bitcoin and adds it to our Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in deposit is %v \n", &w.balance)
	w.balance += amount
}

// ErrInsufficientFunds holds our error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw takes out some amout of Bitcoin from our wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance tells us how much Bitcoin is in our Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
