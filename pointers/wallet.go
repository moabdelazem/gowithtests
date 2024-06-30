package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type BitCoin int

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	WalletBalance BitCoin
}

func (w Wallet) Balance() BitCoin {
	return w.WalletBalance
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.WalletBalance += amount
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.WalletBalance {
		return ErrInsufficientFunds
	}
	w.WalletBalance -= amount
	return nil
}
