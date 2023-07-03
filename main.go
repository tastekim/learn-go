package main

import (
	"fmt"
	"github.com/tastekim/learngo2/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	definition, err := dictionary.Search("firs")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	//account := accounts.NewAccount("tastekim")
	//account.Deposit(10)
	//err := account.WithDraw(20)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(account.String())
}
