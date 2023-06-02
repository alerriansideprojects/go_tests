package pointers_and_errors

import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insuffiecient funds")
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin { return w.balance }
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
