package main

import (
	"fmt"
	"github.com/tastekim/learngo2/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
	//account := accounts.NewAccount("tastekim")
	//account.Deposit(10)
	//err := account.WithDraw(20)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(account.String())
}
