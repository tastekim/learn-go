package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{
		owner: owner, balance: 0,
	}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Onwer() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Onwer(), "'s account.\nHas: ", a.Balance())
}
