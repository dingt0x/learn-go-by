package main

import (
    "errors"
    "fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    // fmt.Printf("address of balance in Deposit is %p\n", &w.balance)
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    // fmt.Printf("address of balance in Balance is %p\n", &w.balance)
    return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
    if w.balance < amount {
        return ErrInsufficientFunds
    }
    w.balance -= amount
    return nil
}
