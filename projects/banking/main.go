package main

import (
	banking "bankAccount"
	"fmt"
)

func main() {
	account := banking.NewAccount("Alex")
	account.Deposit(300)
	account.ChangeOwner("babo")
	fmt.Println(account)

}
