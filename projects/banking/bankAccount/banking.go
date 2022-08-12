package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("You are so poor lol")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (acc *Account) Deposit(amount int) {
	acc.balance += amount
}

func (acc *Account) Widthdraw(amount int) error {
	if acc.balance < amount {
		return errNoMoney
	}
	acc.balance -= amount
	return nil
}

func (acc *Account) Balance() int {
	return acc.balance
}

func (acc *Account) Owner() string {
	return acc.owner
}

func (acc *Account) ChangeOwner(newOwner string) {
	acc.owner = newOwner
}

func (acc *Account) String() string {
	return fmt.Sprint("Name: ", acc.Owner(), "\nBalance: ", acc.Balance())
}
