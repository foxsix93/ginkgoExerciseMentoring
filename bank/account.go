package bank

import "errors"

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit harus positif")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal harus positif")
	}
	if amount > a.Balance {
		return errors.New("saldo kurang")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}
