package main

import (
	"fmt"
	"errors"
     )

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet)Deposit(amount Bitcoin) error{
	w.balance += amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
        return errors.New("cannot withdraw, insufficient funds")
    }
	w.balance -= amount
	return nil
}

func main(){
	wallet := Wallet{}
	err := wallet.Deposit(250)
	wallet.Withdraw(50)
	balance := wallet.Balance()

	if err == nil {
       fmt.Printf("You have Successfuly Deposited the amount %d", balance)
	}
}