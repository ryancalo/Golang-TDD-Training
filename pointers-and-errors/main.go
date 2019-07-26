package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet)Deposit(amount int) bool{
	w.balance += amount
	return true
}

func (w *Wallet) Balance() int {
	return w.balance
}

func main(){
	wallet := Wallet{}
	ok := wallet.Deposit(100)
	balance := wallet.Balance()
	if ok {
       fmt.Printf("You have Successfuly Deposited the amount %d", balance)
	}
}