package main

import (
	"fmt"
	banking "simple/projects/banking/bankAccount"
)

func main() {
	account := banking.NewAccount("Alex")
	account.Deposit(300)
	account.ChangeOwner("babo")
	fmt.Println(account)

}
