package main

import (
	"fmt"
	"github.com/tastekim/learngo2/accounts"
)

func main() {
	account := accounts.NewAccount("tastekim")
	account.Deposit(10)
	err := account.WithDraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.String())
}
