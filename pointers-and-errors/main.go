package main

import "fmt"

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

func (w *Wallet)Deposit(amount Bitcoin) bool{
	w.balance += amount
	return true
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin){
	w.balance -= amount
}

func main(){
	wallet := Wallet{}
	ok := wallet.Deposit(250)
	wallet.Withdraw(50)
	balance := wallet.Balance()

	if ok {
       fmt.Printf("You have Successfuly Deposited the amount %d", balance)
	}
}